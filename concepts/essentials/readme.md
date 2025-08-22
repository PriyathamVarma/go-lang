# 🌟 Go Essentials – Learning Path

A step-by-step roadmap to master Go basics 🚀 and then apply it to **Hyperledger Fabric smart contracts**.

---

## 1️⃣ Hello World in Go
- Start with the classic `Hello, World` program.  
- Learn how to run (`go run`) and build (`go build`) Go code.  
- Understand the `main` package and `main()` function.

---

## 2️⃣ Packages & Imports
- 📦 **Packages** → modular units of code (`main`, `fmt`, `math`, etc.).  
- 🔗 **Imports** → bring in external or standard libraries.  
- Example:
  
  ```go
  import "fmt"
  ````

---

## 3️⃣ Data Types & Variables

* 🔢 **Primitive types**: `int`, `float64`, `string`, `bool`.
* 🧩 **Zero values** (default initialization).
* ✍️ Variable declaration:

  ```go
  var x int = 10
  y := "Hello"
  ```

---

## 4️⃣ Arrays

* Fixed-size collection of elements.
* Example:

  ```go
  var nums [3]int = [3]int{1, 2, 3}
  fmt.Println(nums[0]) // Access first element
  ```

---

## 5️⃣ Loops & Conditions

* ✅ **If/else** for decision-making.
* 🔁 **For loop** is the only loop in Go (acts like while, foreach, etc.).

  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }
  ```

---

## 6️⃣ Functions (with multiple returns)

* Reusable blocks of code.
* Can return **multiple values** (unique feature ✨).

  ```go
  func divide(a, b int) (int, int) {
      return a / b, a % b
  }
  ```

---

## 7️⃣ Structs & JSON Tags

* Structs = custom data types with fields.
* JSON tags help in **marshaling/unmarshaling** data for APIs.

  ```go
  type Car struct {
      Make  string `json:"make"`
      Model string `json:"model"`
      Year  int    `json:"year"`
  }
  ```

---

## 8️⃣ Hyperledger Fabric Imports

* Special packages to interact with blockchain network.
* Common ones:

  ```go
  import "github.com/hyperledger/fabric-contract-api-go/contractapi"
  ```

---

## 📖 Go Language Essentials

| Topic                 | Link                                                                 |
|------------------------|----------------------------------------------------------------------|
| 🟢 Hello World         | [hello-world.md](https://github.com/PriyathamVarma/go-lang/blob/main/concepts/essentials/hello-world.md) |
| 🔢 Data Types          | [data-types.md](https://github.com/PriyathamVarma/go-lang/blob/main/concepts/essentials/data-types.md) |
| ✍️ Alternative Variables | [alternative-variables.md](https://github.com/PriyathamVarma/go-lang/blob/main/concepts/essentials/alternative-variables.md) |
| Formatting | [formatting.md](https://github.com/PriyathamVarma/go-lang/edit/main/concepts/essentials/formatting.md) |


