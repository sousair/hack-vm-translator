# Hack VM Translator

### Introduction
This project is an implementation of the Hack VM Translator, as suggested in Chapter 7 of the book [The Elements of Computing Systems](https://nand2tetris.org). The translator is designed to convert a stack-based VM code written in the Hack VM language into Hack Assembly files that can be further converted into binary thats executed on the Hack computer architecture.

## Usage (Linux)
To translate your Hack VM code, follow these steps:

1. Build the translator:
   ```bash
   go mod tidy
   go build -v -o vm-translator
   ```

2. Run the translator on your input Hack VM code file (e.g., File.vm):
   ```bash
   ./vm-translator file/path/File.vm
   ```

## Additional Resources
For more details on the specifications of the Hack VM Machine, you can refer to the official book chapter **lecture** [here](https://drive.google.com/file/d/1BPmhMLu_4QTcte0I5bK4QBHI8SACnQSt/view), which covers the core concepts and design of the Hack VM language.
