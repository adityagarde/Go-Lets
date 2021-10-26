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

