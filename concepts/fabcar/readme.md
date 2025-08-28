## 🟢 Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

---

## 🟢 Arrays

```go
package main

import "fmt"

func main() {
    var draValue [3]string

    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

---

## 🟢 Functions

```go
func SumDiff(n1 int, n2 int) (int, int) {
    sum := n1 + n2
    diff := n1 - n2
    return sum, diff
}
```

---

## 🟢 Structs (Car & SmartContract)

```go
import "fmt"

type SmartContract struct{}

type Car struct {
    Make   string
    Model  string
    Colour string
    Owner  string
}
```

---

## 🟢 Chaincode Functions (Hyperledger Fabric Example)

```go
// Function selector
function, args := APIstub.GetFunctionAndParameters()

if function == "queryCar" {
    // query car logic
} else if function == "initLedger" {
    // init ledger logic
} else if function == "createCar" {
    // create car logic
} else if function == "queryAllCars" {
    // query all cars logic
} else if function == "changeCarOwner" {
    // change car owner logic
} else {
    return shim.Error("Invalid Smart Contract function name.")
}
```

---

## 🟢 Init Ledger Example

```go
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
    cars := []Car{
        Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
        Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
        Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
        Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
        Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
    }

    for i := 0; i < len(cars); i++ {
        fmt.Println("Added", cars[i])
    }

    return sc.Response{}
}
```

---

## 🟢 Create Car Function

```go
func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 5 {
        return shim.Error("Incorrect arguments. Expecting 5")
    }

    var car = Car{
        Make:  args[1],
        Model: args[2],
        Colour: args[3],
        Owner: args[4],
    }

    // Store car in ledger (omitted)
    return sc.Response{}
}
```

---

## 🟢 Query All Cars

```go
var buffer bytes.Buffer

for resultsIterator.HasNext() {
    // build JSON response
}

fmt.Printf("- queryAllCars:\n%s\n", buffer.String())
```

---

## 🟢 Change Car Owner

```go
func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    if len(args) != 2 {
        return shim.Error("Expecting 2 arguments")
    }

    car := Car{}
    // update logic here
    return sc.Response{}
}
```

---

# ✅ Summary

* **Hello World → Arrays → Functions → Structs → Chaincode (init, create, query, update)**


