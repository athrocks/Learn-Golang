# 06_if_else

This example demonstrates how conditional logic works in Go using `if`, `else if`, and `else`. It also covers declaring variables inside `if` statements and highlights that Go does **not** have a ternary (`? :`) operator.

## 📄 File Structure

- `ifelse.go`: Shows multiple uses of conditional branching in Go.

## 🧠 Concepts Covered

- **Basic `if-else if-else` structure**:
  ```go
  if age >= 18 && age < 22 {
    ...
  } else if age >= 22 {
    ...
  } else {
    ...
  }
  ```

- **Multiple conditions with logical AND (`&&`)**:
  ```go
  if role == "admin" && hasPermission {
    ...
  }
  ```

- **Variable declaration inside `if` block**:
  ```go
  if age := 19; age >= 18 {
    ...
  }
  ```

- **Go does not have a ternary operator**:
  You must use standard `if-else` instead of something like `condition ? a : b`.

## ▶️ How to Run

```bash
go run ifelse.go
```

## ✅ Sample Output

```
You are an Teen.
You have access to the admin panel.
You are an Adult of age 19
```

## 📌 Notes

- Variables declared inside an `if` condition (e.g., `if age := 19`) are scoped to that block only.
- Using `&&` and `||` allows for combining multiple conditions.
- Unlike many other languages, **Go intentionally omits the ternary operator** to keep the syntax simple and readable.

---