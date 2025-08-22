

# 🟢 Variables in Go

## 1. Explicit Type Declaration

```go
var something int = 10
```

* Here you tell Go the type (`int`).
* Use this when you want **strict typing**.

---

## 2. Type Inference

```go
var something = 10
```

* Go automatically infers `something` is an `int`.
* Cleaner and easier for most cases.

---

## 3. Short Declaration (`:=`)

```go
something := 10
```

* Shorthand for `var something = 10`.
* **Most commonly used** inside functions.
* ❌ Not allowed at package level (only inside functions).

---

## 4. Multiple Declarations

```go
var a, b int = 10, 20
var x, y = 10, "20"
```

* You can declare multiple variables at once.
* Types can be same or mixed.

---

# 🟢 Constants in Go

```go
const forever = 100
```

* `const` values **cannot be changed** after declaration.

* If you try:

  ```go
  const forever = 100
  forever = 200 // ❌ compile error
  ```

  Error:

  ```
  cannot assign to forever (declared const)
  ```

* Useful for values that should never change:

  * Pi, max buffer size, app version, etc.

---

# 📘 Calculator Example with Variables + Constants

```go
package main

import (
    "fmt"
    "priyatham/calculator/add"
)

func main() {
    fmt.Print("Calculator \n")

    // Different variable declarations
    var a int = 10         // explicit type
    b := 20                // short declaration
    var c = 1.1            // type inferred as float64

    // Constants
    const forever = 100
    // forever = 200 // ❌ this will cause compile error

    // Using variables in add package
    result := add.Add(a, int(c))
    fmt.Println("10 + 1.1 =", result)

    subtract := add.Sub(10, b)
    fmt.Println("10 - 20 =", subtract)

    division := add.Div(10, 20)
    fmt.Println("10 / 20 =", division)

    multiply := add.Mul(10, 20)
    fmt.Println("10 * 20 =", multiply)

    fmt.Println("Constant forever =", forever)
}
```

---

# 🟢 Importance of Variables

* **Store values** → keep track of numbers, strings, etc.
* **Reusability** → instead of writing numbers directly everywhere, assign to a variable.
* **Flexibility** → values can change at runtime (except constants).
* **Type safety** → Go enforces type rules, preventing silly bugs.
* **Readability** → variable names describe what the value is for.

---

# ✅ Key Points Recap

* Use `:=` for quick variable declaration inside functions.
* Use `var` when:

  * You want explicit type.
  * You’re at package level.
* Use `const` when value should **never change**.
* Trying to reassign to a `const` gives **compile-time error**.

