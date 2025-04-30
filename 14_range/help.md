**📁 14_range**

This example explores the use of the `range` keyword in Go — a powerful way to iterate over slices, maps, and strings.

---

### 🧠 Key Concept: `range` in Go

The `range` keyword is used in `for` loops to iterate over:
- **Slices/Arrays** — returns index and value
- **Maps** — returns key and value
- **Strings** — returns index and Unicode code point (rune)

---

### 📌 Examples

#### 🔢 Iterating Over a Slice

```go
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
	sum += num
}
fmt.Println("sum:", sum)
```

#### 🔍 Finding Index by Value

```go
for i, num := range nums {
	if num == 3 {
		fmt.Println("index:", i)
	}
}
```

#### 🗺️ Iterating Over a Map

```go
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
	fmt.Printf("%s -> %s\n", k, v)
}
```

#### 🔑 Iterating Over Keys Only

```go
for k := range kvs {
	fmt.Println("key:", k)
}
```

#### 🔠 Iterating Over a String

```go
for i, c := range "go" {
	fmt.Println(i, c)
}
```
- Prints the index and Unicode code point of each character.

---

### ▶️ Run the Code

```bash
go run 01_range.go
```