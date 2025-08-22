
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

# 🟢 Variables in Go

```go
var something = 10 // declare + initialize

// RULE in Go: every declared variable must be used,
// otherwise Go will throw a compile-time error.
```

👉 This is different from languages like JavaScript or Python where unused variables are allowed.

---

# 📘 Example with Variables

```go
package main

import (
	"fmt"
	"priyatham/calculator/add"
)

func main() {
	fmt.Print("Calculator \n")

	// declare and initialize
	var _a = 10
	var _b = 20

	// using variables in functions
	result := add.Add(_a, _b)
	fmt.Println("10 + 20 =", result)

	subtract := add.Sub(10, 20)
	fmt.Println("10 - 20 =", subtract)

	division := add.Div(10, 20)
	fmt.Println("10 / 20 =", division)

	multiply := add.Mul(10, 20)
	fmt.Println("10 * 20 =", multiply)
}
```

✅ Key point:

* If you wrote `var unused = 100` and never used it → **compile error**:

  ```
  unused variable 'unused'
  ```

---

# 🟢 Value Types

Go is **strongly typed**, meaning you can’t freely mix types like `int` and `float64` in operations.

---

## Example: Mixing `int` and `float64`

```go
package main

import (
	"fmt"
	"priyatham/calculator/add"
	"math"
)

func main() {
	fmt.Print("Calculator \n")

	var _a = 10      // int
	var _b = 20      // int
	var _c = 1.1     // float64 (default for decimals in Go)

	// must convert float64 to int before passing to Add
	result := add.Add(_a, int(_c))
	fmt.Println("10 + 1.1 =", result) // prints 11

	subtract := add.Sub(10, _b)
	fmt.Println("10 - 20 =", subtract)

	division := add.Div(10, 20)
	fmt.Println("10 / 20 =", division)

	multiply := add.Mul(10, 20)
	fmt.Println("10 * 20 =", multiply)

	// Example using math package (works with float64)
	fmt.Println("Square root of 25 =", math.Sqrt(25)) // 5
}
```

---

# 🔍 Explanation of Value Types

* `int` → whole numbers (`10, 20`).
* `float64` → decimal numbers (`1.1, 2.5`).
* Conversion is **explicit** in Go → no automatic type casting.

  * Example: `int(1.1)` → converts `1.1` into `1`.

---

# ✅ Key Takeaways

* **Variables** must always be used → else compile-time error.
* **Strong typing** → Go doesn’t mix `int` and `float64` automatically.
* **Explicit conversion** needed → `int(_c)` or `float64(_a)`.
* **math package** functions (like `Sqrt`) work with `float64`, so you may need conversions.

---

Let’s dive into **data types in Go**.
Go is a **statically typed language**, which means every variable has a specific type, and the type is checked at compile time.

---

# 🟢 Basic Data Types in Go

## 1. Numbers

### 🔹 Integers

* Signed: `int`, `int8`, `int16`, `int32`, `int64`
* Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`

```go
var a int = 10
var b uint = 20
var small int8 = -128
var big int64 = 9223372036854775807
```

⚡ `int` and `uint` are platform-dependent (32-bit or 64-bit).

---

### 🔹 Floating Point

* `float32` (approx 7 decimal digits)
* `float64` (approx 15 decimal digits, default for decimals)

```go
var pi float32 = 3.14
var e float64 = 2.718281828459
```

---

### 🔹 Complex Numbers

* `complex64` (float32 real + float32 imaginary)
* `complex128` (float64 real + float64 imaginary)

```go
var c1 complex64 = 1 + 2i
var c2 complex128 = complex(3, 4)
fmt.Println(real(c2)) // 3
fmt.Println(imag(c2)) // 4
```

---

## 2. Strings

* Immutable sequences of bytes (UTF-8 encoded text).

```go
var name string = "Priyatham"
fmt.Println(len(name))       // length in bytes
fmt.Println(name[0])         // byte at index 0
fmt.Println(string(name[0])) // convert byte to string
```

⚡ Double quotes (`" "`) → string literal
⚡ Backticks (`` ` ` ``) → raw string (multi-line, no escapes)

---

## 3. Boolean

* Only `true` or `false`.

```go
var flag bool = true
if flag {
    fmt.Println("Yes!")
}
```

---

## 4. Derived / Composite Types

### 🔹 Arrays

* Fixed size collection of elements of the same type.

```go
var arr [3]int = [3]int{1, 2, 3}
fmt.Println(arr[0]) // 1
```

---

### 🔹 Slices

* Dynamic, flexible view into arrays.
* Most used collection in Go.

```go
nums := []int{1, 2, 3, 4}
nums = append(nums, 5) // dynamic resize
```

---

### 🔹 Maps (key-value pairs)

```go
m := map[string]int{
    "apples":  10,
    "oranges": 20,
}
fmt.Println(m["apples"]) // 10
```

---

### 🔹 Structs

* Custom data type grouping fields.

```go
type Car struct {
    Make  string
    Model string
    Year  int
}
myCar := Car{Make: "Tesla", Model: "X", Year: 2025}
```

---

### 🔹 Pointers

* Stores memory address of a value.

```go
x := 10
p := &x
fmt.Println(*p) // 10 (dereference pointer)
```

---

### 🔹 Functions as Types

* Functions can be assigned to variables or passed around.

```go
func add(a, b int) int {
    return a + b
}
var f func(int, int) int = add
fmt.Println(f(2, 3)) // 5
```

---

### 🔹 Interfaces

* Defines behavior, not data.

```go
type Shape interface {
    Area() float64
}
```

---

## 🔑 Quick Table of Go Data Types

| Category  | Types                                                                         |
| --------- | ----------------------------------------------------------------------------- |
| Numeric   | int, int8, int16, int32, int64, uint, float32, float64, complex64, complex128 |
| Text      | string                                                                        |
| Boolean   | bool                                                                          |
| Composite | array, slice, map, struct                                                     |
| Special   | pointer, function, interface                                                  |

---

## ⚡ Key Points

* Go is **strongly typed** → no automatic conversion between `int` and `float64`.
* **Zero values**: if not initialized, variables take default values:

  * Numbers → `0`
  * Strings → `""`
  * Booleans → `false`
  * Pointers, slices, maps → `nil`



















