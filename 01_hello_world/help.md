# 01_hello_world

This is a simple Go program that prints **"Hello World!"** to the console. It demonstrates the basic structure of a Go application, including the `main` package, the `main()` function, and the use of the `fmt` package for output.

## 📄 File Structure

- `main.go`: Contains the source code for the Hello World program.

## 🧠 Concepts Covered

- `package main`: Defines the entry point of a Go executable program.
- `import "fmt"`: Imports the format package for I/O operations.
- `func main()`: The main function where execution starts.
- `fmt.Println(...)`: Prints text to the standard output with a newline.

## ▶️ How to Run

### Using `go build` (compiles the binary):

```bash
go build main.go
./main
```

### Using `go run` (compiles and runs in one step):

```bash
go run main.go
```

## ✅ Output

```
Hello World!
```

## 📌 Notes

- `go run` is quicker for testing small programs.
- `go build` creates an executable file which can be run multiple times without recompilation.

---