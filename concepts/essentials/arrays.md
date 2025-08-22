

# 🟢 Arrays in Go

## 1. What is an Array?

* An **array** is a **fixed-size collection** of elements of the same type.
* The size is part of the type → `[3]int` is **different** from `[5]int`.

---

## 2. Declaring Arrays

### Explicit Size

```go
var arr [3]int
arr[0] = 10
arr[1] = 20
arr[2] = 30

fmt.Println(arr) // [10 20 30]
```

### Initialization

```go
arr := [3]int{1, 2, 3}
fmt.Println(arr) // [1 2 3]
```

### Infer Size

```go
arr := [...]int{5, 6, 7, 8}
fmt.Println(arr) // [5 6 7 8]
```

(`...` lets Go count elements for you)

---

## 3. Accessing & Modifying

```go
arr := [3]string{"Go", "is", "cool"}
fmt.Println(arr[0]) // Go
arr[1] = "super"
fmt.Println(arr) // [Go super cool]
```

---

## 4. Iterating Over Arrays

### Using `for`

```go
arr := [3]int{10, 20, 30}
for i := 0; i < len(arr); i++ {
    fmt.Println("Index:", i, "Value:", arr[i])
}
```

### Using `range`

```go
for index, value := range arr {
    fmt.Println(index, value)
}
```

---

## 5. Arrays are Value Types

* Arrays are **copied by value** when passed to functions.
* Changes inside a function won’t affect the original array.

```go
func change(arr [3]int) {
    arr[0] = 99
}

func main() {
    nums := [3]int{1, 2, 3}
    change(nums)
    fmt.Println(nums) // [1 2 3] (unchanged)
}
```

👉 If you want modifications, pass a **pointer** or use a **slice**.

---

## 6. Multidimensional Arrays

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
fmt.Println(matrix[1][2]) // 6
```

---

# 📘 Example: Calculator with Array of Numbers

```go
package main

import "fmt"

func main() {
    numbers := [5]int{10, 20, 30, 40, 50}

    sum := 0
    for _, n := range numbers {
        sum += n
    }

    fmt.Println("Numbers:", numbers)
    fmt.Println("Sum of array elements =", sum)
}
```

👉 Output:

```
Numbers: [10 20 30 40 50]
Sum of array elements = 150
```

---

# ✅ Key Takeaways

* Arrays are **fixed size**.
* Use `[...]` to let Go infer size.
* Arrays are **value types** (copied on assignment/passing).
* For flexibility → use **slices** (next step).

