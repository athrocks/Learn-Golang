**📁 10_maps**

This example explores **maps** in Go, which are unordered key-value pairs similar to dictionaries in Python or objects in JavaScript.

---

### 🧠 Key Concepts

#### 1. **Creating a Map**

```go
m := make(map[string]int)
```

Add key-value pairs:

```go
m["a"] = 11
m["b"] = 22
```

---

#### 2. **Accessing Values**

```go
fmt.Println(m["a"]) // returns 11
fmt.Println(m["h"]) // key doesn't exist → returns 0 (zero value for int)
```

To check if a key exists:

```go
v, ok := m["a"]
if ok {
	fmt.Println("a:", v)
}
```

---

#### 3. **Iterating Over a Map**

```go
for k := range m {
	fmt.Println(k, m[k])
}
```

---

#### 4. **Deleting and Clearing**

```go
delete(m, "a") // removes key "a"
clear(m)       // empties the map
```

---

#### 5. **Map Literals**

```go
mpp := map[string]int{
	"k1": 11,
	"k2": 22,
}
```

---

#### 6. **Comparing Maps**

Maps cannot be compared using `==`. Use the `maps` package (Go 1.21+):

```go
import "maps"
maps.Equal(m1, m2) // true if all keys and values are equal
```

---

### ▶️ Run the Code

```bash
go run maps.go
```