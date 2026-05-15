# Go Interpreter

An interpreter written in Golang for the programming language Monkey.

## Why an Interpreter?

Why not?

I have always loved the idea of building something fun while learning along the way. The thought of building something this complex used to intimidate me, so I decided to take the plunge and explore the world of designing more advanced systems.

I chose to do it in Go, a language I originally learned for building CLI tools and one that has quickly become one of my favorites alongside TypeScript. The learning curve has been steep, but the book _Writing an Interpreter in Go_ by Thorsten Ball has been an incredible motivator and learning resource throughout the journey.

## Progress So Far

1. Implemented the basic tokenizer/lexer with support for both single-character and multi-character tokens.
2. Built a basic REPL for testing the tokenizer alongside the official lexer tests.
3. Implemented the parser, which is now functional and integrated with the REPL.
