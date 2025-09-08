# Chaincode

* **Chaincode (CC)** is the smart-contract/business logic that lets clients interact with the **shared ledger**.
* When you **invoke a transaction**, you call a chaincode function that performs reads/writes against the world state.
* Chaincode runs inside **containers** (Docker) managed by Fabric — this provides sandboxing/isolation and reproducible runtimes.
* Fabric uses **signed base images** for chaincode runtimes (Go, Node, Java) to enforce secure, auditable execution environments.
* Chaincode should be **deterministic**, lightweight, and avoid heavy external I/O (offload heavy work off-chain).

---

# Sample Go Chaincode — `car_cc` (Car asset example)

chaincode/car_cc.go` Exxample with `CreateCar`, `TransferCar`, `QueryCar`, and `GetCarHistory`. Uses Fabric Contract API (`github.com/hyperledger/fabric-contract-api-go/contractapi`).

```go
// chaincode/car_cc.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Car describes basic details of a car asset
type Car struct {
	CarID string `json:"carID"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Owner string `json:"owner"`
	VIN   string `json:"vin,omitempty"`
	Year  int    `json:"year,omitempty"`
}

// SmartContract provides functions for managing a Car
type SmartContract struct {
	contractapi.Contract
}

// CreateCar creates a new car in the world state
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carID, make, model, owner string) error {
	exists, err := s.CarExists(ctx, carID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("car %s already exists", carID)
	}

	car := Car{
		CarID: carID,
		Make:  make,
		Model: model,
		Owner: owner,
	}

	carJSON, err := json.Marshal(car)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(carID, carJSON)
}

// TransferCar changes owner of a car
func (s *SmartContract) TransferCar(ctx contractapi.TransactionContextInterface, carID, newOwner string) error {
	carJSON, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return fmt.Errorf("failed to get car: %v", err)
	}
	if carJSON == nil {
		return fmt.Errorf("car %s does not exist", carID)
	}

	var car Car
	if err := json.Unmarshal(carJSON, &car); err != nil {
		return err
	}

	car.Owner = newOwner
	updated, err := json.Marshal(car)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(carID, updated)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carID string) (*Car, error) {
	carJSON, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if carJSON == nil {
		return nil, fmt.Errorf("car %s does not exist", carID)
	}

	var car Car
	if err := json.Unmarshal(carJSON, &car); err != nil {
		return nil, err
	}
	return &car, nil
}

// GetCarHistory returns the history of a car key as a JSON array
func (s *SmartContract) GetCarHistory(ctx contractapi.TransactionContextInterface, carID string) ([]map[string]interface{}, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(carID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []map[string]interface{}
	for resultsIterator.HasNext() {
		mod, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		record := map[string]interface{}{
			"TxId":      mod.TxId,
			"Timestamp": mod.Timestamp, // proto timestamp; may need formatting client-side
			"IsDelete":  mod.IsDelete,
		}
		var value interface{}
		if len(mod.Value) > 0 {
			if err := json.Unmarshal(mod.Value, &value); err == nil {
				record["Value"] = value
			} else {
				record["Value"] = string(mod.Value)
			}
		}
		records = append(records, record)
	}
	return records, nil
}

// CarExists helper
func (s *SmartContract) CarExists(ctx contractapi.TransactionContextInterface, carID string) (bool, error) {
	data, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating car chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting car chaincode: %s", err.Error())
	}
}
```

### Build & package

* Build image for chaincode (if using external builder) or use Fabric lifecycle packaging:

  * `GO111MODULE=on go mod init car_cc` and add required modules.
  * Use Fabric samples `peer lifecycle chaincode` commands to package/install/approve/commit.

---

# Minimal Node.js client examples (invoke/query)

Install:

```bash
npm install fabric-network
```

### Submit CreateCar (Node.js)

```js
// client/createCar.js
const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

async function main() {
  const ccpPath = path.resolve(__dirname, 'connection.json'); // your SDK connection profile
  const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

  const wallet = await Wallets.newFileSystemWallet('./wallet');
  const gateway = new Gateway();
  await gateway.connect(ccp, {
    wallet,
    identity: 'DealerUser',
    discovery: { enabled: true, asLocalhost: true }
  });

  const network = await gateway.getNetwork('carchannel');
  const contract = network.getContract('car_cc');

  // submit transaction (create)
  await contract.submitTransaction('CreateCar', 'CAR100', 'Toyota', 'Corolla', 'DealerA');
  console.log('Transaction submitted: CreateCar CAR100');

  await gateway.disconnect();
}

main().catch(console.error);
```

### Query Car (Node.js)

```js
// client/queryCar.js
const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

async function main() {
  const ccp = JSON.parse(fs.readFileSync(path.resolve(__dirname, 'connection.json'), 'utf8'));
  const wallet = await Wallets.newFileSystemWallet('./wallet');
  const gateway = new Gateway();
  await gateway.connect(ccp, { wallet, identity: 'DealerUser', discovery: { enabled:true, asLocalhost:true } });

  const network = await gateway.getNetwork('carchannel');
  const contract = network.getContract('car_cc');

  const result = await contract.evaluateTransaction('QueryCar', 'CAR100');
  console.log('QueryCar result:', result.toString());

  await gateway.disconnect();
}

main().catch(console.error);
```

---

# Notes about Docker & Chaincode containers

* When chaincode is instantiated, a chaincode container is started on the peer host (using the runtime image for Go/Node/Java). That container runs only the chaincode process and communicates with the peer via gRPC.
* Sandboxing + signed base images provide reproducible runtimes and reduce risk of malicious chaincode.
* In production, Fabric uses external builders & runtime images (and you can run chaincode as external service for more control).

