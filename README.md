# Monkey Interpreter - in GoLang

> An interpreter for the Monkey programming language written in Go.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)

---

# About

This project is a tree-walking interpreter for the Monkey programming language, built entirely in Go.

What originally started as curiosity quickly became one of the most rewarding learning experiences I’ve had as a developer. Building an interpreter always felt like one of those impossibly smart people project ,the one thing that was out of my leauge.
So naturally, I decided to never try to even go this route until I was reccomended a god sent book by one of my favourite youtubers ( he worked at netflix btw / also uses neovim btw).

And Drum Rollll ..... The incredible book is:

**Writing an Interpreter in Go** by Thorsten Ball

This book has been one of the most fun and technical learning resources I’ve ever used.

---

## Lexer / Tokenizer

- Single-character token support
- Multi-character token support
- Identifier parsing
- Integer parsing
- Keyword recognition
- Illegal token detection

---

## Parser

- Pratt parser implementation
- Abstract Syntax Tree (AST) generation
- Operator precedence parsing
- Prefix expressions
- Infix expressions
- Grouped expressions
- Conditional parsing
- Function literal parsing
- Call expression parsing

---

## REPL

- Interactive command-line REPL
- Live parsing & evaluation
- Error reporting
- Fast testing workflow

---

## Evaluator / Interpreter

- Integer evaluation
- Boolean evaluation
- Prefix operators
- Infix operators
- Conditional execution
- Return statements
- Variable bindings
- Function objects
- Closures
- Lexical scoping
- Environment handling

---

# Not Yet Implemented

- Type system
- Built-in functions
- Arrays & hash maps
- Standard library
- Better parser recovery

---

# Project Structure

```bash
.
├── ast/          # AST node definitions
├── evaluator/    # Runtime evaluator
├── lexer/        # Tokenizer implementation
├── object/       # Runtime object system
├── parser/       # Pratt parser
├── repl/         # Interactive REPL
├── token/        # Token definitions
└── main.go
```

---

# Getting Started

## Clone the Repository

```bash
git clone https://github.com/Tack-12/go-interpreter.git
cd go-interpreter
```

---

## Run the Interpreter

```bash
go run main.go
```

---

# Example

```monkey
let add = fn(a, b) {
    a + b;
};

add(5, 10);
```

Output:

```bash
15
```

---

# Progress

- [x] Lexer
- [x] Parser
- [x] AST generation
- [x] REPL
- [x] Evaluator
- [x] Functions & closures
- [x] Lexical environments
- [x] Type system
- [x] Built-in functions
- [x] Extending the full Programming Language Experience

---

# Example

```monkey
let multiply = fn(a, b) {
    a * b;
};

multiply(6, 7);
```

```bash
42
```
