**Car supply-chain** example 

1. A **high-level flowchart** showing orgs, peers, chaincode, CAs, orderer.
2. A **sequence diagram** showing the exact endorsement → ordering → validation sequence for a `createCar` and `transferCar` txn (including invalid path).
3. A simple **UML (class) diagram** for on-chain data model + chaincode interface.

After each diagram a compact explanation and the concrete mapping to Manufacturer / Assembler / Dealer / Buyer.

---

## 1) High-level Flowchart (Architecture + Components)

```mermaid
flowchart TB
  subgraph ClientApp["Client Applications"]
    DealerApp["Dealer App (UI)"]
    AssemblerApp["Assembler App (UI)"]
    BuyerApp["Buyer App (UI)"]
  end

  subgraph Network["Fabric Network (Channel: carchannel)"]
    subgraph Org_Mfg["Org: Manufacturer"]
      M_Peer["Peer Mfg (endorser, ledger)"]
      M_CA["CA (Mfg)"]
      M_CC["Chaincode (car_cc)"]
      M_Peer -->|runs| M_CC
    end

    subgraph Org_Assem["Org: Assembler"]
      A_Peer["Peer Assem (endorser, ledger)"]
      A_CA["CA (Assem)"]
      A_CC["Chaincode (car_cc)"]
      A_Peer -->|runs| A_CC
    end

    subgraph Org_Dealer["Org: Dealer"]
      D_Peer["Peer Dealer (endorser, ledger)"]
      D_CA["CA (Dealer)"]
      D_CC["Chaincode (car_cc)"]
      D_Peer -->|runs| D_CC
    end

    subgraph Org_Buyer["Org: Buyer"]
      B_Peer["Peer Buyer (committer/ledger)"]
      B_CA["CA (Buyer)"]
    end

    Orderer["Ordering Service (Raft)"]
  end

  %% client -> SDK -> peers
  DealerApp -->|SDK: submitProposal/submitTransaction| D_Peer
  AssemblerApp --> A_Peer
  BuyerApp --> B_Peer

  %% SDK also talks to endorsers in other orgs (simulated)
  D_Peer --- M_Peer
  D_Peer --- A_Peer
  D_Peer --- B_Peer

  %% Orderer connections
  M_Peer <--> Orderer
  A_Peer <--> Orderer
  D_Peer <--> Orderer
  B_Peer <--> Orderer

  %% CAs are off-to-side for identity mgmt
  M_CA -.-> M_Peer
  A_CA -.-> A_Peer
  D_CA -.-> D_Peer
  B_CA -.-> B_Peer
```

**Explanation / Mapping**

* `car_cc` is the chaincode (smart contract) deployed/installed on peers of Manufacturer, Assembler, Dealer. Buyer may have a committer peer (no need to endorse).
* Each Org has a CA that issues identity certificates (MSP). Clients (apps) use certs to sign proposals.
* Ordering Service (Raft) orders transactions into blocks and sends them to all peers.

---

## 2) Sequence Diagram (CreateCar / TransferCar + invalid path)

```mermaid
sequenceDiagram
    autonumber
    participant DealerApp as Dealer App (Client)
    participant SDK as Fabric SDK
    participant EndM as Endorser (Manufacturer Peer)
    participant EndA as Endorser (Assembler Peer)
    participant Orderer as Ordering Service
    participant PeerM as Peer Manufacturer (Commit)
    participant PeerD as Peer Dealer (Commit)
    participant World as World State DBs

    Note over DealerApp,EndA: CREATE CAR (CAR100)
    DealerApp->>SDK: createProposal("createCar","CAR100",...)
    SDK->>EndM: proposalRequest
    SDK->>EndA: proposalRequest
    EndM-->>SDK: proposalResponse (RWSet, Sig)
    EndA-->>SDK: proposalResponse (RWSet, Sig)
    SDK-->>Orderer: submitTransaction (RWSet + endorsements)
    Orderer-->>PeerM: deliverBlock
    Orderer-->>PeerD: deliverBlock
    PeerM->>PeerM: validate (endorse policy, MVCC)
    PeerD->>PeerD: validate (endorse policy, MVCC)
    PeerM-->>World: apply WriteSet (CAR100 created)
    PeerD-->>World: apply WriteSet (CAR100 created)

    Note over DealerApp,PeerD: TRANSFER CAR (race condition example)
    DealerApp->>SDK: createProposal("transferCar","CAR100","Buyer1")
    SDK->>EndM: proposalRequest (read version v1)
    SDK->>EndA: proposalRequest (read version v1)
    EndM-->>SDK: proposalResponse (RWSet v1->v2)
    EndA-->>SDK: proposalResponse (RWSet v1->v2)
    SDK-->>Orderer: submitTransaction (TxA)
    %% Meanwhile, another client submits a competing transaction
    alt competing tx earlier applied
      Note right of Orderer: TxB (competing) ordered before TxA
      Orderer-->>PeerM: deliverBlock (contains TxB then TxA)
      PeerM->>PeerM: validate TxB (valid) -> apply WriteSet (owner=BuyerX)
      PeerM->>PeerM: validate TxA (MVCC fail) -> mark INVALID
      PeerM-->>World: apply TxB writes
    else no conflict
      Orderer-->>PeerM: deliverBlock (TxA ordered)
      PeerM->>PeerM: validate TxA (valid) -> apply WriteSet (owner=Buyer1)
      PeerM-->>World: apply TxA writes
    end
```

**Explanation**

* **Proposal stage**: SDK sends proposal to endorsers (Manufacturer, Assembler). Endorsers simulate and return signed RWSet.
* **Submit stage**: SDK packages endorsements and sends to **Orderer**.
* **Ordering**: Orderer sequences transactions into blocks.
* **Validation & commit**: Peers validate endorsement policy and MVCC. If a competing transaction (TxB) got ordered earlier and changed the read-version, TxA becomes invalid — recorded on chain but not applied to world state.

---

## 3) UML Class Diagram (Asset model + Chaincode interface)

```mermaid
classDiagram
    class Car {
      +string carID
      +string make
      +string model
      +string owner
      +string vin
      +int year
      +string status
      +getHistory() : []TransactionRecord
    }

    class CarChaincode {
      +Init(ctx)
      +CreateCar(ctx, carID, make, model, owner) : error
      +TransferCar(ctx, carID, newOwner) : error
      +QueryCar(ctx, carID) : Car
      +GetCarHistory(ctx, carID) : []TransactionRecord
    }

    class TransactionRecord {
      +string txID
      +string timestamp
      +string action
      +string initiator
      +Car before
      +Car after
    }

    CarChaincode --> Car
    Car --> TransactionRecord : produces
```

**Explanation**

* `Car` is the asset stored in world state (JSON). Fields include `carID`, `make`, `model`, `owner`, `vin`, etc.
* `CarChaincode` is the smart contract API: `CreateCar`, `TransferCar`, `QueryCar`, `GetCarHistory`.
* `TransactionRecord` represents ledger entries you can build from block history (use `GetHistoryForKey` in chaincode).

---

## Extra: Concrete SDK + CLI examples mapped to this car flow

### Create Car (Node.js SDK)

```js
const network = await gateway.getNetwork('carchannel');
const contract = network.getContract('car_cc');

// submitTransaction does proposal+submit => waits for commit
await contract.submitTransaction('CreateCar', 'CAR100', 'Toyota', 'Corolla', 'DealerA');
```

### Transfer Car (Node.js SDK)

```js
await contract.submitTransaction('TransferCar', 'CAR100', 'Buyer1');
```

### CLI Equivalent (invoke)

```bash
# create
peer chaincode invoke -C carchannel -n car_cc \
  -c '{"Args":["CreateCar","CAR100","Toyota","Corolla","DealerA"]}' \
  --orderer orderer.example.com:7050

# transfer
peer chaincode invoke -C carchannel -n car_cc \
  -c '{"Args":["TransferCar","CAR100","Buyer1"]}' \
  --orderer orderer.example.com:7050
```

---

## How this maps to your original image content

* **Solution User / Solution Provider** → Dealer/Assembler apps (Client Applications) + SDK.
* **Non-validating node (NV)** → optional gateway nodes or client-side peers that forward requests (could be API servers).
* **Validating nodes** → endorsing peers (Manufacturer, Assembler, Dealer peers).
* **MSP / CA** → each org’s CA issuing ECerts, TCerts (if used), TLS certs.
* **Chaincode state and logic** → `car_cc` deployed on endorsing peers.
* **Ordering service** → Raft orderers in the network (shared).
* **Private/Confidential transactions** → can be implemented via private data collections if some fields (e.g., price, VIN) should be visible only to subset of orgs.


