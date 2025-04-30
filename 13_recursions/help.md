**📁 13_recursions**

This example demonstrates **recursion** in Go through the classic **factorial** function.

---

### 🧠 Key Concept: Recursion

Recursion is when a function calls itself to solve smaller instances of the same problem.

---

### 📌 Example: Factorial Function

```go
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
```

- **Base Case**: `fact(0)` returns `1`.
- **Recursive Case**: `fact(n)` returns `n * fact(n-1)`.

---

### 🧪 Output

```go
println(fact(5)) // 120
println(fact(0)) // 1
println(fact(1)) // 1
```

---

### ▶️ Run the Code

```bash
go run recursions.go
```