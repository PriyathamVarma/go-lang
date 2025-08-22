# 🟢 Functions in Go

## 1. Syntax

```go
func functionName(parameterName type, parameterName type, ...) returnType {
    // function body
    return value
}
```

---

## 2. Example: Simple Add Function

```go
func Add(a int, b int) int {
    return a + b
}
```

* `Add` → function name (capitalized, so **exported** if in a package).
* `(a int, b int)` → parameters with types.
* `int` after parentheses → return type.
* `return a + b` → returns the sum.

---

## 3. Calling Functions

```go
func Something() {
    fmt.Println("This is a simple function")
}
```

* Functions can be defined **anywhere in the file** (before or after `main`).
* You just need to **call** them in `main` or another function.

```go
func main() {
    Something()
    fmt.Println(Add(10, 20))
}
```

---

## 4. Return Values and Types

### 🔹 Single Return Value

```go
func Square(n int) int {
    return n * n
}
```

---

### 🔹 Multiple Return Values

Go supports returning multiple values (very common in Go).

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

Calling it:

```go
result, err := Divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

👉 This is how Go handles **errors** — by returning `(value, error)` instead of exceptions.

---

### 🔹 Named Return Values

You can give names to return variables:

```go
func SumAndDiff(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return // implicit return sum, diff
}
```

Calling:

```go
s, d := SumAndDiff(10, 5)
fmt.Println("Sum:", s, "Diff:", d)
```

---

## 5. Variadic Functions

Functions can take **variable number of arguments** using `...`.

```go
func Total(nums ...int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

fmt.Println(Total(1, 2, 3, 4, 5)) // 15
```

---

# ✅ Key Points Recap

* `func` keyword defines a function.
* Parameters require explicit types.
* Functions can return multiple values.
* Errors are usually handled as a second return value.
* Variadic functions accept any number of arguments.

