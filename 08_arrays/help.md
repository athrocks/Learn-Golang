**📁 08_arrays**

This example demonstrates how to use **arrays** in Go, including basic operations, initialization techniques, and 2D arrays.

---

### 🧠 Key Concepts

#### 1. **Array Declaration & Default Values**

```go
var nums [5]int
fmt.Println(nums) // [0 0 0 0 0]
```

All elements are initialized to the **zero value** for the type (`0` for `int`, `false` for `bool`, `""` for `string`).

---

#### 2. **Accessing & Modifying Elements**

```go
nums[0] = 1
fmt.Println(nums[0]) // 1
```

---

#### 3. **Array Initialization**

```go
a := [3]int{11, 22, 33}
```

Use `...` to let the compiler infer the length:

```go
b := [...]int{1, 2, 3, 4, 5}
```

---

#### 4. **Sparse Array with Index Specifier**

```go
c := [...]int{100, 3: 400, 500}
fmt.Println(c) // [100 0 0 400 500]
```

Indexes `1` and `2` are auto-filled with `0`.

---

#### 5. **2D Arrays**

```go
d := [2][2]int{
	{1, 2},
	{3, 4},
}
fmt.Println(d) // [[1 2] [3 4]]
```

---

### ✅ When to Use Arrays

- You **know the size** at compile time.
- You need **constant-time** access to elements.
- You want **fixed-size** containers (consider slices for dynamic needs).

---

### ▶️ Run the Code

```bash
go run arrays.go
```