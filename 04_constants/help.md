# 04_constants

This program introduces **constants** in Go. It demonstrates how to declare constants, both individually and in groups, and highlights Go's restriction that constants cannot be reassigned once declared.

## 📄 File Structure

- `constants.go`: Contains examples of constant declaration and usage.

## 🧠 Concepts Covered

- `const name = "golang"`: Declaring a constant with an inferred type.
- `const age = 30`: Declaring a package-level constant.
- Constants **cannot be reassigned** (e.g., `name = "js"` will cause a compile-time error).
- Grouping constants using `const (...)` syntax:
  ```go
  const (
    port = 8080
    host = "localhost"
  )
  ```

## ▶️ How to Run

```bash
go run constants.go
```

## ✅ Output

```
Name: golang
Age: 30
Port: 8080 Host: localhost
```

## 📌 Notes

- Constants are immutable values known at compile time.
- They help avoid magic numbers and improve readability and maintainability.
- Can be declared both inside and outside functions.

---