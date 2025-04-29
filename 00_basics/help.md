In Go, you write `package main` at the top of a file when that file is part of a **standalone executable program** — meaning, when it should be compiled into a binary that can be run directly (like a command-line app or web server).

### Here's why:

#### ✅ `package main` is **special**
- Go treats `package main` as the entry point for your application.
- When you run `go run` or `go build`, Go looks for a package named `main`, and **a `main()` function inside it**.

#### ✅ It tells the Go compiler: "This is the start of a runnable program"
- Without `package main`, you can’t build a binary.
- For example, if you write `package util` or `package math`, Go assumes you're writing a **reusable library**, not a program to run.

---

### Example

```go
// This will compile and run
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

But this:

```go
// This will NOT compile as a runnable program
package math

func Add(a, b int) int {
	return a + b
}
```

Would compile **only as a library**, not something you can `go run`.

---

### Summary

| `package` name | Purpose                            | Can `go run` it? |
|----------------|-------------------------------------|------------------|
| `main`         | Defines an executable program       | ✅ Yes           |
| any other      | Used for libraries/reusables        | ❌ No            |

