**📁 09_slices**

This example explores **slices**, one of the most powerful and commonly used data structures in Go. Slices are dynamic, flexible views into arrays and are essential in idiomatic Go code.

---

### 🧠 Key Concepts

#### 1. **What Are Slices?**
- **Slices are dynamic arrays** — unlike arrays, their size can change.
- An uninitialized slice is `nil` and has zero length and capacity.

```go
var nums []int
fmt.Println(nums == nil) // true
```

---

#### 2. **Using `make` to Create Slices**

```go
nums2 := make([]int, 5)         // length = 5, capacity = 5
nums3 := make([]int, 3, 5)      // length = 3, capacity = 5
```

You can append elements to a slice:

```go
nums3 = append(nums3, 1, 2, 3)
```

The capacity grows automatically if needed.

---

#### 3. **Literal Slice Initialization**

```go
nums4 := []int{1, 2, 3, 4, 5}
```

---

#### 4. **Slicing Operations**

```go
s := []string{"a", "b", "c", "d", "e", "f"}
l := s[1:4] // ["b", "c", "d"]
m := s[:3]  // ["a", "b", "c"]
n := s[3:]  // ["d", "e", "f"]
```

---

#### 5. **Copying Slices**

```go
c := make([]string, len(s))
copy(c, s)
```

---

#### 6. **Comparing Slices**

Slices can't be compared with `==` directly. Use the `slices` package:

```go
import "slices"

slices.Equal([]int{1, 2}, []int{1, 2}) // true
```

---

#### 7. **2D Slices**

```go
num7 := [][]int{{1, 2}, {3, 4}}
```

---

### ▶️ Run the Code

```bash
go run slices.go
```