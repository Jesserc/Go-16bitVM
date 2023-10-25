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

func main() {

	i := 0
	increment := func() int {
		old := i
		i++
		return old
	}
	memory := CreateMemory(40) // or we use 256 to make memory space large
	// The order goes like this:
	// [instruction, value (two values where necessary)]
	memory[increment()] = MOV_LIT_REG // instruction
	memory[increment()] = 0x0020      // value for the above instruction
	memory[increment()] = R1          // value for the above instruction and so on...
	memory[increment()] = MOV_LIT_REG
	memory[increment()] = 0x0020
	memory[increment()] = R2
	memory[increment()] = ADD_REG_REG
	memory[increment()] = R1
	memory[increment()] = R2
	memory[increment()] = MOV_REG_MEM
	memory[increment()] = ACC
	memory[increment()] = 0x0014 // 20
	memory[increment()] = JUMP_NOT_EQ
	memory[increment()] = 0xffff
	memory[increment()] = 0x0000

	cpu := NewCPU(memory)

	fmt.Println() // line space
	fmt.Printf("Full cpu state before executions: %#+04v\n\n", cpu)

	cpu.debug()
	ip1, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip1)
	cpu.viewMemoryAt(0x14)

	// This will prompt users to press the ENTER key in the terminal to loop through execution process
	for i := 0; i < 5; i++ {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		fmt.Println() // line space
		fmt.Println("Step", i+1)

		cpu.step()
		cpu.debug()
		ip, _ := cpu.getRegister("ip")
		cpu.viewMemoryAt(ip)
		cpu.viewMemoryAt(0x14)
	}

	fmt.Printf("Full cpu state after executions: %#+04v\n\n", cpu)
}
