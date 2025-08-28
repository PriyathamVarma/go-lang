# 🚗 Structs in Go

---

## 🔴 1. What problems do structs solve?

In Go, you often need to group related data together.

* Without structs, you’d have to manage **separate variables** for each property:

```go
carMake := "BMW"
carModel := "X5"
carYear := 2024
carOwner := "Varma"
```

👉 This gets messy when you have many cars — you’d need **arrays of each variable** and keep them in sync.

---

✅ **Solution → Structs**

* A **struct** lets you **group related fields** into a single type.
* It makes your data **organized, reusable, and type-safe**.

---

## 🟢 2. Defining a Struct Type

```go
type Car struct {
    Make  string
    Model string
    Year  int
    Owner string
}
```

* `Car` is a **custom type**.
* It has 4 **fields** (Make, Model, Year, Owner).
* Fields can be of different types (`string`, `int`, etc.).

---

## 🟢 3. Initializing Structs

### Method 1: Full initialization

```go
myCar := Car{
    Make:  "BMW",
    Model: "X5",
    Year:  2024,
    Owner: "Varma",
}
fmt.Println(myCar)
```

👉 Output:

```
{BMW X5 2024 Varma}
```

---

### Method 2: Positional initialization (order matters!)

```go
myCar := Car{"BMW", "X5", 2024, "Varma"}
```

---

### Method 3: Create empty and assign fields later

```go
var myCar Car
myCar.Make = "Tesla"
myCar.Model = "Model 3"
myCar.Year = 2025
myCar.Owner = "Varma"
```

---

## 🟢 4. Accessing Fields

```go
fmt.Println("Car Owner:", myCar.Owner)
```

---

## 🟢 5. Example with Multiple Cars (Array of Structs)

```go
cars := []Car{
    {"VW", "Polo", 2019, "Priyatham"},
    {"BMW", "X5", 2024, "Alice"},
    {"Mercedes", "C-Class", 2023, "John"},
}

for _, car := range cars {
    fmt.Printf("%s %s (%d) owned by %s\n", car.Make, car.Model, car.Year, car.Owner)
}
```

👉 Output:

```
VW Polo (2019) owned by Priyatham
BMW X5 (2024) owned by Alice
Mercedes C-Class (2023) owned by John
```

---

# ✅ Recap

* **Problem solved**: Organizes related values into one type.
* **Define struct**: `type Car struct { ... }`
* **Initialize**: Using `{}` with fields or assign later.
* **Use case**: Collections of cars, owners, etc.

