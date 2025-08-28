# 🟢 Pointers in Go

## 1. What is a Pointer?

* Every variable in Go is stored somewhere in **memory**.

* Example:

  ```go
  a := 2
  ```

  Here:

  * Value: `2`
  * Address: e.g., `0x231451100` (random memory location)

* A **pointer** stores the **address** of a value instead of the value itself:

  ```go
  age := 10
  ageIdentifier := &age   // & = "address of"
  fmt.Println(ageIdentifier) // prints something like 0xc0000180a8
  ```

---

## 2. Why Use Pointers?

### ✅ Advantage 1: Avoid Unnecessary Value Copies

* By default, when you pass variables to functions, Go **copies** their values.
* For big structs/arrays, this wastes memory and CPU.

👉 With pointers, you just pass the **address**, not the whole data.

---

### ✅ Advantage 2: Directly Mutate Values

* If you want a function to change a variable’s value outside its scope, you need a pointer.

---

# 🔴 Example 1: Without Pointers (Value Copy)

```go
package main

import "fmt"

func main() {
    age := 10
    fmt.Println("Before function call:", age)

    aging(age) // passing by value (copy)

    fmt.Println("After function call:", age) // still 10
}

func aging(age int) {
    age = age + 5
    fmt.Println("Inside function:", age) // 15
}
```

👉 Output:

```
Before function call: 10
Inside function: 15
After function call: 10
```

* The original `age` didn’t change, because Go passed a **copy**.

---

# 🟢 Example 2: With Pointers (Mutating Original)

```go
package main

import "fmt"

func main() {
    age := 10
    fmt.Println("Before function call:", age)

    aging(&age) // pass address (pointer)

    fmt.Println("After function call:", age) // now updated
}

func aging(age *int) {
    *age = *age + 5 // * = dereference → get/set value at that address
    fmt.Println("Inside function:", *age)
}
```

👉 Output:

```
Before function call: 10
Inside function: 15
After function call: 15
```

---

## 🔑 Key Symbols

* `&` → **address-of** operator
  Gets the memory address of a variable.

* `*` → **dereference** operator
  Gets/sets the value stored at the memory address.

---

## ✅ Recap

* Variables store values, **pointers store addresses**.
* **Without pointers** → Go passes a **copy** to functions.
* **With pointers** → Functions can directly modify original values.
* Pointers avoid unnecessary memory copying when working with large data.
