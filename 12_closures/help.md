**📁 12_closures**

This example demonstrates **closures** in Go — functions that capture and remember variables from their surrounding scope.

---

### 🧠 Key Concept: Closure

A **closure** is a function value that references variables from outside its body. The function may access and modify these variables even after the outer function has finished executing.

---

### 📌 Example: Counter with Closure

```go
func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}
}
```

- `count` is captured by the returned anonymous function.
- Every call to `increment()` updates and returns the new value of `count`.

---

### 🧪 Output

```go
increment := counter()
fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
```

Each call to `increment()` maintains state via closure — `count` persists across calls.

---

### ▶️ Run the Code

```bash
go run closures.go
```