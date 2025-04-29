**📁 07_switch**

This directory demonstrates the usage of `switch` statements in Go. The examples cover:

- Basic switch-case
- Switch with multiple conditions
- Type switch using interfaces

---

### 🧠 Key Concepts

#### 1. **Basic `switch` Statement**
Go’s `switch` statement automatically breaks after each case, so you **don’t need a `break` keyword**:

```go
switch i {
case 1:
	fmt.Println("i is 1")
...
default:
	fmt.Println("other")
}
```

#### 2. **Multiple Conditions in One Case**
You can combine multiple conditions using commas:

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("It's the weekend")
default:
	fmt.Println("It's a weekday")
}
```

#### 3. **Type Switch**
Used to determine the **dynamic type** of an `interface{}` value:

```go
switch t := i.(type) {
case string:
	fmt.Println("I am a string:", t)
...
default:
	fmt.Println("I don't know what I am:", t)
}
```

Example outputs:
- `"hello"` → string
- `42` → int
- `true` → bool
- `3.14` → unknown (falls into `default`)

---

### 📌 Run the Code

```bash
go run switch.go
```

---