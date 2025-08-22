# 🟢 Input & Output in Go (`fmt` package)

The `fmt` package is used for **printing, scanning, and formatting strings**.

---

## 1. Taking Input with `fmt.Scan`

```go
var something int
fmt.Scan(&something)  // note the & (address of variable)
```

* `fmt.Scan` reads input from the user.
* `&something` → we pass the **memory address** so `Scan` can store the value into the variable.
* Example:

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Println("You entered:", num)
}
```

👉 If you type `42`, output will be:

```
Enter a number: 42
You entered: 42
```

---

## 2. Improving User Input Prompt

Instead of mixing everything in one line, split them:

```go
fmt.Println("Enter a value for something:")
fmt.Scan(&something)
```

👉 This is clearer because `Println` shows a message, then waits for input.

---

## 3. Formatting Output with `fmt.Printf`

`Printf` = **print with format specifiers** (similar to C/Java).

### Common verbs:

* `%v` → default format (works for most values).
* `%d` → integers.
* `%f` → floats.
* `%s` → strings.
* `%t` → booleans.
* `%T` → type of the value.

Example:

```go
something := 42
pi := 3.14159
name := "Priyatham"

fmt.Printf("Value: %v\n", something) // Value: 42
fmt.Printf("Integer: %d\n", something) // Integer: 42
fmt.Printf("Float: %.2f\n", pi)       // Float: 3.14
fmt.Printf("String: %s\n", name)      // String: Priyatham
fmt.Printf("Type of pi: %T\n", pi)    // Type of pi: float64
```

---

## 4. Full Calculator Example with Input + Formatting

```go
package main

import (
    "fmt"
    "priyatham/calculator/add"
)

func main() {
    fmt.Println("Calculator")

    var a, b int
    fmt.Print("Enter first number: ")
    fmt.Scan(&a)

    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    result := add.Add(a, b)
    fmt.Printf("%d + %d = %d\n", a, b, result)
}
```

👉 Sample Run:

```
Calculator
Enter first number: 10
Enter second number: 20
10 + 20 = 30
```

---

# ✅ Key Takeaways

* `fmt.Scan(&x)` → reads user input into variable `x`.
* Always pass the **address** (`&`) so Go can modify the variable.
* Use `fmt.Printf` for clean formatted output.
* Specifiers like `%d`, `%f`, `%s`, `%T` help control output format.


