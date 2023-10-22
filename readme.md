# Minimal 16 bits register based virtual machine

# How it works

The VMs CPU has very small dedicated memory cells (registers) that hold values that are being used while the CPU is executing its operations.
We can move values from main memory to registers, registers to memory and registers to registers.

# Overview

## 1. Components:

#### a. CPU Structure:

- `Memory`: A slice of uint16 to emulate the memory of the VM.
- `Register Names`: A slice of strings representing the names of the registers.

```go
// ip = Instruction Pointer (IP) (register in the CPU that holds the memory address of the next instruction to be executed),
// acc = Accumulator (register in the CPU that holds result of arithmetic operations),
// r1 - r8 = random registers
registerNames: []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"},
```

- `Registers`: A slice of uint16 to store the values of the registers.
- `Register Map`: A map to link register names to their indices in the registers slice.

#### b. NewCPU Function:

Initializes a new CPU with a given memory, set of register names, an empty register map, and a register slice of a specified size.

#### c. Instruction set:

Only three instructions are defined:

```go
const (
	MOV_LIT_R1  = 0x10 // move literal r1 (register)
	MOV_LIT_R2  = 0x11 // move literal r2 (register)
	ADD_REG_REG = 0x12 // add register (r1) to register (r2)
)
```

## 2. Function Operations:

#### a. Register Operations:

- `getRegister`: Searches for a register by name and returns its value. Returns an error if the register name is not found.
- `setRegister`: Searches for a register by name and sets its value. Returns an error if the register name is not found.

#### b. Instruction Fetching and Execution (functions):

- `fetch`: Fetches the next instruction from memory, increments the instruction pointer (ip), and returns the instruction.
- `fetch16`: Calls fetch to get the next 16-bit instruction from memory.
- `execute`: Decodes and executes a given instruction.
  The supported instructions are:

  - `MOV_LIT_R1`: Moves a literal value into register r1.
  - `MOV_LIT_R2`: Moves a literal value into register r2.
  - `ADD_REG_REG`: Adds the values of two registers and stores the result in the accumulator (acc) register.

#### c. Step:

- `step`:
  Calls fetch and execute functions in sequence to fetch and execute the next instruction.

### d. `Debug`:

- `debug`:
  Prints the values of all registers in a formatted manner, padding the hexadecimal representation of register values with leading zeros to ensure a consistent width of 4 characters.

## 3. Execution Steps:

A block comment in `cpu.go` illustrates the sequence of state changes when executing a series of `step` instructions.

## 4. Missing/Incomplete Parts:

- #### Error Handling:

  Error handling could be improved, particularly in the fetch, fetch16, execute, and step methods, where errors are currently ignored on purpose.

- #### Instruction Set:

  The instruction set is limited (by design).

- #### Instruction Encoding/Decoding:
  There's no formalized instruction encoding/decoding scheme.
