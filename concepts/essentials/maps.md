# 🟢 Maps in Go

---

## 1. Declaring a Map

```go
cars := map[string]string{
    "car1": "VW",
    "car2": "BMW",
    "car3": "Merc",
}
fmt.Println(cars)
```

👉 Output:

```
map[car1:VW car2:BMW car3:Merc]
```

---

## 2. Accessing & Mutating a Map

```go
fmt.Println(cars["car1"]) // VW

cars["car1"] = "Audi"     // update value
fmt.Println(cars["car1"]) // Audi
```

---

## 3. Adding & Deleting

```go
cars["car4"] = "Tesla" // add new entry
delete(cars, "car2")   // remove entry
```

---

## 4. Iterating a Map

```go
for key, value := range cars {
    fmt.Println(key, ":", value)
}
```

---

## 5. Checking Existence

```go
val, ok := cars["car10"]
if ok {
    fmt.Println("Found:", val)
} else {
    fmt.Println("Not found")
}
```

---

# 🟢 Maps vs Structs

| Feature     | Structs                                    | Maps                         |
| ----------- | ------------------------------------------ | ---------------------------- |
| Keys        | Field names (fixed, known at compile time) | Any type (string, int, etc.) |
| Size        | Fixed fields                               | Dynamic (can grow/shrink)    |
| Performance | Faster, strongly typed                     | Slower (hash lookup)         |
| Use Case    | Well-defined objects (Car, User)           | Dynamic collections, lookup  |

👉 Rule of thumb:

* **Structs** = when fields are known & fixed.
* **Maps** = when keys are dynamic or unknown in advance.

---

# 🟢 `make()` Function

## Problem with `append`

```go
slic := []int{}
slic = append(slic, 1)
fmt.Println("address of slice", &slic[0])

slic = append(slic, 2)
fmt.Println("address of slice", &slic[0])
```

👉 Here the **address may change** because when capacity is exceeded, Go allocates a **new array** behind the scenes → old one is copied over.

---

## Solution: Pre-allocate with `make`

```go
slic := make([]int, 0, 5) // len=0, cap=5

slic = append(slic, 1)
fmt.Println("address of slice", &slic[0])

slic = append(slic, 2)
fmt.Println("address of slice", &slic[0])

fmt.Println("Actual slice", slic)
```

👉 Since we gave capacity 5, appends don’t need to reallocate until we exceed 5 items → memory address stays stable.

---

# 🟢 `make()` with Maps

Just like slices, you can use `make()` with maps:

```go
cars := make(map[string]string)

cars["car1"] = "VW"
cars["car2"] = "BMW"

fmt.Println(cars)
```

👉 This initializes a map ready for use.
Without `make`, a map is `nil` and writing to it will panic.

---

# ✅ Recap

* **Maps** are key-value stores, dynamic & flexible.
* **Structs** are fixed-field, faster, and strongly typed.
* Use **`make`** to pre-allocate slices/maps for performance.
* For **slices** → `make([]T, len, cap)`
* For **maps** → `make(map[K]V, hintSize)`


