package main

import (
	"bufio"
	"fmt"
	"os"
)

const IP = 0
const ACC = 1
const R1 = 2
const R2 = 3
const R3 = 4
const R4 = 5
const R5 = 6
const R6 = 7
const R7 = 8
const R8 = 9
const SP = 10
const FP = 11

const MEMORY_LENGTH = 40

func main() {

	i := 0
	increment := func() int {
		old := i
		i++
		return old
	}
	memory := CreateMemory(MEMORY_LENGTH) // or we use 256 to make memory space large
	// The order goes like this:
	// [instruction, value (two values where necessary)]
	// memory[increment()] = MOV_LIT_REG // instruction
	// memory[increment()] = 0x0020      // value for the above instruction
	// memory[increment()] = R1          // value for the above instruction and so on...
	// memory[increment()] = MOV_LIT_REG
	// memory[increment()] = 0x0020
	// memory[increment()] = R2
	// memory[increment()] = ADD_REG_REG
	// memory[increment()] = R1
	// memory[increment()] = R2
	// memory[increment()] = MOV_REG_MEM
	// memory[increment()] = ACC
	// memory[increment()] = 0x0014 // 20
	// memory[increment()] = JUMP_NOT_EQ
	// memory[increment()] = 0xffff
	// memory[increment()] = 0x0000

	memory[increment()] = MOV_LIT_REG
	memory[increment()] = 0x5151
	memory[increment()] = R1
	memory[increment()] = MOV_LIT_REG
	memory[increment()] = 0x4242
	memory[increment()] = R2
	memory[increment()] = PUSH_REG
	memory[increment()] = R1
	memory[increment()] = PUSH_REG
	memory[increment()] = R2
	// swap the value of R1 and R2
	memory[increment()] = POP
	memory[increment()] = R1
	memory[increment()] = POP
	memory[increment()] = R2

	cpu := NewCPU(memory)

	fmt.Println() // line space
	fmt.Printf("Full cpu state before executions: %#+04v\n\n", cpu)

	cpu.debug()
	ip, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip)
	cpu.viewMemoryAt(MEMORY_LENGTH - 8) // view last 8 indexes in memory

	// This will prompt users to press the ENTER key in the terminal to loop through execution process
	for i := 0; i < 6; i++ {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		fmt.Println() // line space
		fmt.Println("Step", i+1)

		cpu.step()
		cpu.debug()
		ip, _ := cpu.getRegister("ip")
		cpu.viewMemoryAt(ip)
		cpu.viewMemoryAt(MEMORY_LENGTH - 8) // view last 8 indexes in memory

	}

	fmt.Printf("Full cpu state after executions: %#+04v\n\n", cpu)
}
