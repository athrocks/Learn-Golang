# 03_variables

This program demonstrates how to **declare and use variables** in Go. It covers explicit type declaration, type inference, and the importance of using declared variables.

## 📄 File Structure

- `main.go`: Contains examples of variable declarations and usage.

## 🧠 Concepts Covered

- `var a int = 10`: Declaring a variable with an explicit type.
- `var name string = "golang"`: Declaring a string variable.
- `var name2 = "golang also"`: Using **type inference** — Go automatically determines the type based on the value.
- **Unused variables cause compile-time errors** — every declared variable must be used.

## ▶️ How to Run

```bash
go run main.go
```

## ✅ Output

```
10
golang
golang also
```

## 📌 Notes

- Go enforces **strict variable usage**: if you declare a variable and don’t use it, the program won’t compile.
- Type inference makes variable declarations concise but readable.
- Variable names should be meaningful to improve code readability.

---