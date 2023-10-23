package main

import "fmt"

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
	memory[increment()] = 0x1234      // value for the above instruction
	memory[increment()] = R1          // value for the above instruction and so on...
	memory[increment()] = MOV_LIT_REG
	memory[increment()] = 0xABCD
	memory[increment()] = R2
	memory[increment()] = ADD_REG_REG
	memory[increment()] = R1
	memory[increment()] = R2
	memory[increment()] = MOV_REG_MEM
	memory[increment()] = ACC
	memory[increment()] = 0x0014 // 20
	memory[increment()] = JUMP_NOT_EQ
	memory[increment()] = 0xbe01
	memory[increment()] = 0x10 // 42

	cpu := NewCPU(memory)

	fmt.Println() // line space
	fmt.Printf("Full cpu state before executions: %#+04v\n\n", cpu)

	cpu.debug()
	ip1, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip1)
	cpu.viewMemoryAt(0x14)

	fmt.Println() // line space
	fmt.Println("Step 1")

	cpu.step()
	cpu.debug()
	ip2, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip2)
	cpu.viewMemoryAt(0x14)

	fmt.Println() // line space
	fmt.Println("Step 2")

	cpu.step()
	cpu.debug()
	ip3, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip3)
	cpu.viewMemoryAt(0x14)

	fmt.Println() // line space
	fmt.Println("Step 3")

	cpu.step()
	cpu.debug()
	ip4, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip4)
	cpu.viewMemoryAt(0x14)

	fmt.Println() // line space
	fmt.Println("Step 4")

	cpu.step()
	cpu.debug()
	ip5, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip5)
	cpu.viewMemoryAt(0x14)

	fmt.Println() // line space
	fmt.Println("Step 5")

	cpu.step()
	cpu.debug()
	ip6, _ := cpu.getRegister("ip")
	cpu.viewMemoryAt(ip6)
	cpu.viewMemoryAt(0x14)
	fmt.Println() // line space

	fmt.Printf("Full cpu state after executions: %#+04v\n\n", cpu)

}
