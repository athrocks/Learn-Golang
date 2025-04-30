**📁 11_functions**

This example demonstrates how to define and use **functions** in Go, including named functions, multiple return values, variadic parameters, and functions as first-class citizens.

---

### 🧠 Key Concepts

#### 1. **Basic Function with Parameters and Return Value**

```go
func add(a int, b int) int {
	return a + b
}
```

Consecutive parameters of the same type can be written more concisely:

```go
func plusPlus(a, b, c int) int {
	return a + b + c
}
```

---

#### 2. **Multiple Return Values**

```go
func getLang() (string, string) {
	return "Go", "Js"
}
```

Use `_` to ignore one of the return values:

```go
_, lang := getLang()
```

---

#### 3. **Variadic Functions**

Functions can accept any number of arguments using `...`:

```go
func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
```

---

#### 4. **Functions as Parameters**

Functions are first-class citizens in Go — you can pass them as arguments:

```go
func processIt(fn func(a int) int) {
	fmt.Println(fn(1))
}
```

---

#### 5. **Returning a Function**

```go
func processIt2() func(a int) int {
	return func(a int) int {
		return 400
	}
}
```

Call the returned function like so:

```go
res := processIt2()
fmt.Println(res(0)) // 400
```

---

### ▶️ Run the Code

```bash
go run functions.go
```

---