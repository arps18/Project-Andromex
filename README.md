# Project Andromex

Welcome to **Project Andromex**, a simple Go-based compiler project designed for learning both Go and compiler fundamentals at the same time!

## What is this?

Project Andromex is a minimal compiler written in Go. It helps you understand how programming languages are tokenized, parsed, and evaluated, while also giving you hands-on experience with Go’s syntax and tooling.

## Features

- **Lexer**: Converts source code into tokens.
- **Token definitions**: Clearly defined token types for easy extension.
- **REPL**: Interactive prompt for experimenting with expressions.
- **Easy to extend**: Add your own language features as you learn.

### Upcoming Updates

- **Parser**: Build an Abstract Syntax Tree (AST) from tokens.
- **Semantic Analysis**: Add symbol tables and type checking.
- **Intermediate Representation (IR)**: Generate a lower-level form for analysis/translation.
- **Optimization**: Constant folding, dead code elimination, peephole optimizations.
- **Code Generation**: Output to a virtual machine or target assembly.
- **Error Handling**: More helpful syntax and runtime error messages.
- **Control Flow**: Support for `if`, `while`, `for`, etc.
- **Functions**: User-defined functions with parameters and return values.
- **Data Structures**: Arrays, structs, and beyond.
- **Object-Oriented Features**: Classes, methods, and inheritance.

## Getting Started

1. **Clone the repository**

   ```bash
   git clone https://github.com/arps18/Project-Andromex.git
   cd Project-Andromex
   ```

2. **Run the REPL**

   ```bash
   go run main.go
   ```

3. **Try it out!**
   Type expressions like `2 + 3 * 4` and see how the lexer tokenizes your input.

## Why use this project?

- Learn Go by building something practical.
- Understand compiler basics: lexing, parsing, and evaluation.
- Experiment and extend as you grow.

## Contributing

Pull requests and suggestions are welcome!  
Feel free to fork and make this project your own.

## License

This project is licensed under the GNU v3 License.

---

Happy hacking and learning!
