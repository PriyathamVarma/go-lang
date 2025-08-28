# 🟢 Conditionals & Loops in Go

---

## ✅ `if` Statement

```go
package main

import "fmt"

func main() {
    var cars = [3]string{"VW", "BMW", "Merc"}
    var selectCar int

    fmt.Println("Select car [0,1,2]:")
    fmt.Scan(&selectCar)

    // Simple if condition
    if selectCar >= len(cars) {
        fmt.Println("No cars")
    }

    fmt.Println(cars[selectCar])
}
```

### Key Points:

* `if` does **not require parentheses** around condition.
* The `{}` braces are **mandatory**.
* `len(cars)` returns the length (3 in this case).
* ⚠️ Use `>=` instead of `>` here — because arrays are **0-indexed**, index `3` would be invalid.

---

## ✅ `if ... else`

```go
package main

import "fmt"

func main() {
    var cars = [3]string{"VW", "BMW", "Merc"}
    var selectCar int

    fmt.Println("Select car [0,1,2]:")
    fmt.Scan(&selectCar)

    if selectCar >= len(cars) {
        fmt.Println("No cars")
    } else {
        fmt.Println("You selected:", cars[selectCar])
    }
}
```

---

## ✅ `for` Loops

Go has only one loop: `for`. It can be used in multiple ways.

### Classic `for`

```go
for i := 0; i < 2; i++ {
    fmt.Println("Iteration", i)
}
```

### With condition only (like `while`)

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### Infinite loop

```go
for {
    fmt.Println("Runs forever unless break")
    break
}
```

---

### Example with Cars + Loop

```go
package main

import "fmt"

func main() {
    var cars = [3]string{"VW", "BMW", "Merc"}

    for i := 0; i < 2; i++ {
        var selectCar int
        fmt.Println("Select car [0,1,2]:")
        fmt.Scan(&selectCar)

        if selectCar >= len(cars) {
            fmt.Println("No cars")
        } else {
            fmt.Println("You selected:", cars[selectCar])
        }
        fmt.Println("---------------")
    }
}
```

---

## ✅ `switch` Statement

```go
package main

import "fmt"

func main() {
    var car int
    fmt.Println("Enter a day number [1-3]:")
    fmt.Scan(&car)

    switch car {
    case 1:
        fmt.Println("BMW")
    case 2:
        fmt.Println("VW")
    case 3:
        fmt.Println("MERC")
    default:
        fmt.Println("NO Car")
    }
}
```

### Key Points:

* No need for `break`; Go automatically exits after a case.
* You can use `default` like in other languages.
* Multiple matches allowed:

  ```go
  switch car {
  case 1, 2, 3:
      fmt.Println("VW  models")
  case 4, 5:
      fmt.Println("Tata cars")
  default:
      fmt.Println("No cars")
  }
  ```

---

# ✅ Recap

* `if` / `else` → branch logic
* `for` → the **only loop** (acts as for, while, infinite)
* `switch` → clean way to handle multiple conditions

