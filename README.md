# crafting-interpreters

Implementation of the Jlox and Clox compilers from the book [Crafting Interpreters](https://craftinginterpreters.com/), with the former being wriiten in Go instead of Java, and the latter in the original C code found in the book.

## Glox

Glox is a tree-walk intrepreter for the Lox programming language - a type of interpreter that executes a program by traversing its abstract syntax tree (AST) and performing the appropriate operations for each node. The AST is a tree-like data structure that represents the syntactic structure of the program.

## Clox

Clox is a bytecode compiler for the Lox programming language - a program that translates source code written in a high-level programming language into bytecode, which is a low-level, platform-independent representation of the program. The bytecode can then be executed by a virtual machine or interpreter.
