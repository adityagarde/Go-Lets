## GO Let's

- Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- The Gopher is the official mascot of the Go Language. 
<img src = https://miro.medium.com/max/1400/1*jnT9eoQMlc96bGxTnIbK9g.jpeg width=300>

### 0. Getting Started
- [Download and Install](https://golang.org/doc/install)

- To check the installation and version use -
```go
$$ go version
```

- Go is a **compiled** language - which means the go code is directly converted into machine code that the processor can execute.
- On the other hand, for interpreted languages like JavaScript, code is read line by line and a different program, i.e. the interpreter, reads and executes the code.
- Because of this compiled languages like Go tend to be faster than Interpreted languages.

<br>

- Running a program & HELP commands - 
```go
go run main.go

go help // go help <command>
```

### 1. Go Fundamentals

- Types
    - Almost everything is a `Type` in Go.
    - Various `Types` in Go
        - String, Bool, Integer (uint8, uint64, int8, int64, uintptr), Floating (float32, float64), Complex
        - Arrays, Slices, Structs, Maps, Pointers
        - Functions, Channels, Mutex.

- Is Go Object Oriented ?
    - Yes and No. Although Go has types and methods and *allows* an object-oriented style of programming, there is no type hierarchy.

#### **Variable and Constant Declaration -**

- General Variable declaration - [Examples](https://github.com/adityagarde/Go-Lets/blob/main/02variables/main.go)
```go
// 1. Explicitly mentioning the type
var username string = "Aditya Garde"
var smallValue uint8 = 255;
// 2. Implicit Type detection
var someString = "Aditya Garde"
// 3. Short Variable declaration (Walrus Operation)
someNumber := 347209.4590;
```
- In Go, we will get error if we don't use a variable for which we have assigned some value.
- Short Variable declaration (using `:=`) is allowed ONLY inside a function and not outside.
- Constants can be declared using `const` keyword.
```go
const LoginToken string = "Sherlock Holmes"
const Pi = 3.14159265358979323846
```
- In Go, all identifiers are lexically (statically) scoped, i.e. scope of a variable can be determined at compile time.
- **Local Variables** are variables declared inside function or block and **Global Variables** are variables declared outside some function or block.

#### **Packages**

- A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

- The first statement in a Go source file must be package name. Executable commands must always use package main.

- A Go repository typically contains only one module, located at the root of the repository. A file named [`go.mod`](https://github.com/adityagarde/Go-Lets/tree/main/03userinput) there declares the module path - the import path prefix for all packages within the module.
```go
$ go mod init example/user/hello
```
- Standard packages can be searched [here](https://pkg.go.dev/).

#### [**Taking User Input**](https://github.com/adityagarde/Go-Lets/blob/main/03userinput/main.go)

```go
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("User Input : ")

	// comma ok OR comma err syntax - nearly equals for try catch
	input, _ := reader.ReadString('\n')
```

#### **Environment Variables**

- To check environment variables -
```go
$ go env
$ go env <NAME> //go env GOPATH
```
- To change the current property -
```go
$ go env -w <NAME>=<VALUE> //go env -w GOBIN=/some/dfferent/bin
```
- To unset a variable previously set by `go env -w` -
```go
$ go env -u GOBIN
```
- Build for any OS (MocOS, Windows, Linux) can be build from any host OS using the environment `GOOS` (and NOT `GOHOSTOS`)
```go
$ go build //default
$ GOOS="linux" go build // "darwin", "windows"
```

### 2. Memory Management

- Memory allocation and de-allocation happens automatically in Golang.
- `new()` vs `make()`
    - `new()` 
        - Memory is allocated, but is **not** initialized.
        - You will get a memory address.
        - Zeroed Storage.
    - `make()` - Generally preferred.
        - Memory is allocated and initialized as well.
        - You will get a memory address.
        - Non - Zeroed Storage.
- Garbage collection happens automatically. Anything which becomes out of scope or nil is available for garbage collection.
    - The `GOGC` variable sets the initial garbage collection target percentage.
    - A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default value is GOGC=100.

#### **Pointers**

-  A pointer holds the memory address of a value (variable).
- Instead of passing the value of the variable (which can be big, or can be changed by some other operation contextually) - we just pass the address location to that variable / value, ensuring only the correct value from the memory location is picked.

- The `&` operator generates a pointer to its operand.
- The `*` operator denotes the pointer's underlying value.
```go
    i := 42
    ptr := &i         // point to i
    fmt.Println("ptr == ", ptr)  // random memory location
    fmt.Println("*ptr == ", *ptr) // read i through the pointer, actual value
    *ptr = *ptr + 20         // set i through the pointer
    fmt.Println("i == ", i)  // see the new value of i
    fmt.Println("*ptr == ", i)
```
output - 
```text
ptr ==  0xc000018030
*ptr ==  42
i ==  62
*ptr ==  62
```

### Some commands for go modules

- `go mod tidy`
- `go mod verify`
- `go list`
- `go list all`
- `go list -m all`
- `go list -m -versions github.com/gorilla/mux`
- `go mod why github.com/gorilla/mux`
- `go mod graph`
- `go mod edit -go 1.16` || `go mod edit -module newModuleName`
- `go mod vendor`
- `go run -mod=vendor main.go`