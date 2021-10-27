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

go env GOPATH
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
var username string = "Aditya Garde";
var smallValue uint8 = 255;
// 2. Implicit Type detection
var someString = "Aditya Garde";
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

#### **Scope of Variables -** 

- In Go, all identifiers are lexically (statically) scoped, i.e. scope of a variable can be determined at compile time.
- **Local Variables** are variables declared inside function or block and **Global Variables** are variables declared outside some function or block.
