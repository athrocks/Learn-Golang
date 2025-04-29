# 05_for_loop

This program demonstrates all the ways the `for` loop can be used in Go. It's the only loop construct in Go but can be adapted to behave like `while`, `for-each`, infinite loops, and more.

## 📄 File Structure

- `for.go`: Contains examples of all types of `for` loop usage.

## 🧠 Concepts Covered

- **While-style loop**:
  ```go
  i := 1
  for i <= 3 {
    ...
  }
  ```
- **Classic for loop**:
  ```go
  for i := 1; i <= 3; i++ {
    ...
  }
  ```
- **Range loop over a slice**:
  ```go
  for i, v := range []int{10, 21, 32} {
    ...
  }
  ```
- **Range over an integer (Go 1.22+)**:
  ```go
  for i := range 3 {
    fmt.Println(i)
  }
  ```
  Iterates from `0` to `2`.

- **Infinite loop** (commented out to avoid hanging):
  ```go
  for {
    ...
  }
  ```

- **Loop with `break`**: Exits early when a condition is met.
- **Loop with `continue`**: Skips the current iteration when a condition is met.

## ▶️ How to Run

```bash
go run for.go
```

## ✅ Sample Output

```
1
2
3
1
2
3
0 10
1 21
2 32
0
1
2
1
1
3
```

## 📌 Notes

- `range <int>` is valid **from Go 1.22 onwards** and behaves like a loop from `0` to `n-1`.
- All looping in Go is done with `for`.
- `break` and `continue` give you more control over flow inside loops.

---