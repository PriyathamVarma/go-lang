
# 📦 Project Structure

```
calculator/                 <- Module root (go.mod here)
│
├── go.mod
│
├── calculator.go           <- package main (entry point)
│
└── add/                    <- folder = package add
    └── add.go              <- package add (Add function)
```

---

# 🟢 Code

## `calculator.go` (main package)

```go
// package main means this is the executable entry point
package main

import (
    "fmt"
    "priyatham/calculator/add" // importing the add package
)

func main() {
    fmt.Println("Calculator")

    result := add.Add(10, 20) // calling Add from add package
    fmt.Println("10 + 20 =", result)
}
```

---

## `add/add.go` (add package)

```go
package add

// Add adds two integers and returns the sum
func Add(a int, b int) int {
    return a + b
}
```

---

# 🔍 Explanation

### 1. Why `package main` in `calculator.go`?

* Because Go requires an **entry point** (executable program).
* Only the `main` package with `main()` function is run when you execute `go run` or `go build`.

---

### 2. Why `package add` in `add.go`?

* Because each folder is a **separate package**.
* The folder `add/` defines `package add`.
* Functions in `add.go` can be used by other packages if they start with a **capital letter** (e.g. `Add`).

---

### 3. Why `"priyatham/calculator/add"` in import?

* This matches your **module name**.
* If you did:

  ```bash
  go mod init priyatham/calculator
  ```

  then the **import path** for `add` is:

  ```
  priyatham/calculator/add
  ```
* Go uses `go.mod` to resolve this import.

---

### 4. Data Types in Play

* `int` is one of Go’s **basic data types**.
* Your `Add` function takes two `int` arguments and returns an `int`.
* Go is strongly typed, so you can’t add strings here, only integers.

---

## ✅ Key Takeaways

* **Folders = Packages** in Go.
* **Capitalized functions/variables = exported** (public).
* `main` package = entry point.
* Other packages = reusable libraries.
* Data types like `int` must be explicitly declared in function signatures.

---

