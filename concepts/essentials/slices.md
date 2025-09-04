# 🟢 Slices in Go

---

## 🔴 1. Why Slices?

* Arrays in Go have **fixed size** (cannot grow/shrink).
* Example:

  ```go
  var carValues = [4]int{1, 2, 3, 4}
  ```

  Here `carValues` always has 4 elements.

👉 Problem: What if you want a **dynamic list** (like Python list or JS array)?
✅ **Solution: Slices** → built on top of arrays, but more flexible.

---

## 🟢 2. Creating a Slice from an Array

```go
var carValues = [4]int{1, 2, 3, 4}

// Slice syntax: array[start:end]
requiredValues := carValues[1:3]

fmt.Println(requiredValues) // [2 3]
```

### Key Points:

* `carValues[1:3]` → includes index 1 and 2, excludes index 3.
* Think of it like: **start inclusive, end exclusive**.
* So you get `[2, 3]`.

---

## 🟢 3. Slice Properties

A slice has 3 things:

1. **Pointer** → points to underlying array.
2. **Length** → number of elements in slice.
3. **Capacity** → how many elements you can grow into from starting index.

```go
fmt.Println(len(requiredValues)) // 2
fmt.Println(cap(requiredValues)) // 3 (from index 1 to end of carValues)
```

---

## 🟢 4. Updating Through Slices

Slicing does **not copy**, it points to the same array underneath.

```go
carValues := [4]int{1, 2, 3, 4}
requiredValues := carValues[1:3]

requiredValues[0] = 99

fmt.Println(carValues)      // [1 99 3 4]
fmt.Println(requiredValues) // [99 3]
```

👉 Both updated, because they share memory.

---

## 🟢 5. Declaring Slices Directly

Instead of arrays, you can create slices directly:

```go
cars := []string{"VW", "BMW", "Merc"} // slice, not array
fmt.Println(cars)
```

Notice: no size `[3]` → it’s automatically a slice.

---

## 🟢 6. Append to Slices

Unlike arrays, slices can grow.

```go
cars := []string{"VW", "BMW"}
cars = append(cars, "Merc")

fmt.Println(cars) // [VW BMW Merc]
```

---

## 🟢 7. Copying Slices

If you want a **true copy** (not shared memory):

```go
original := []int{1, 2, 3}
copySlice := make([]int, len(original))
copy(copySlice, original)

copySlice[0] = 99
fmt.Println(original)   // [1 2 3]
fmt.Println(copySlice)  // [99 2 3]
```

---

# ✅ Recap

* **Array** = fixed size.
* **Slice** = flexible, points to array.
* Slice syntax: `[start:end]` (start inclusive, end exclusive).
* Updating through slices also updates underlying array.
* Use `append()` to grow slices.
* Use `copy()` to clone.

