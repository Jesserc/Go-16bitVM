package main

import (
	"errors"
	"fmt"
)

type CPU struct {
	memory        []uint16
	registerNames []string
	registers     []uint16
	registerMap   map[string]int
}

func NewCPU(memory Memory) CPU {
	cpu := CPU{
		memory:        memory,
		registerNames: []string{"ip", "acc", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "r8"},
		registerMap:   make(map[string]int),
	}

	cpu.registers = CreateMemory(len(cpu.registerNames))

	for i, name := range cpu.registerNames {
		cpu.registerMap[name] = i
	}
	return cpu
}
func (c *CPU) debug() {
	fmt.Printf("======Printing Registers======\n\n")

	for _, name := range c.registerNames {
		registerValue, _ := c.getRegister(name)
		formattedRegisterValue := fmt.Sprintf("0x%04x", registerValue)
		fmt.Printf("%v: %v\n", name, formattedRegisterValue)
	}
}
func (c *CPU) viewMemoryAt(address uint16) {

	nextEightBytes := make([]uint16, 4) // or 8

	for i := 0; i < 4; i++ {
		// hexValue := fmt.Sprintf("0x%#04x", (*c).memory[address+uint16(i)])
		nextEightBytes[i] = (*c).memory[address+uint16(i)]
		// nextEightBytes[i] = fmt.Sprintf("0x%04x", (*c).memory[address+uint16(i)])

	}

	fmt.Printf("nextEightBytes: %#+04v\n\n", nextEightBytes)
}
func (c CPU) getRegister(name string) (uint16, error) {
	for _, regName := range c.registerNames {
		if regName == name {
			return c.registers[c.registerMap[name]], nil
		}
	}
	return 0, errors.New("register name not found")
}

func (c *CPU) setRegister(name string, value uint16) error {
	for _, regName := range c.registerNames {
		if regName == name {
			c.registers[c.registerMap[name]] = value
			return nil
		}
	}
	return errors.New("register name not found")
}

func (c *CPU) fetch() uint16 {
	instructionIndex, _ := c.getRegister("ip")
	instruction := c.memory[instructionIndex]
	c.setRegister("ip", instructionIndex+1)
	return instruction
}

func (c *CPU) fetch16() uint16 {
	return c.fetch()
}

func (c *CPU) execute(instruction uint16) {
	switch instruction {
	// Move literal into register
	case MOV_LIT_REG:
		{
			literal := (*c).fetch16()
			register := (*c).fetch() % uint16(len(c.registerNames))
			// (*c).setRegister("r1", literal)
			(*c).registers[register] = literal

			return
		}
	// Move register to register
	case MOV_REG_REG:
		{
			registerFrom := (*c).fetch() % uint16(len(c.registerNames))
			registerTo := (*c).fetch() % uint16(len(c.registerNames))
			value := (*c).registers[registerFrom]

			(*c).registers[registerTo] = value
			return
		}
		// Move register to memory
	case MOV_REG_MEM:
		{
			registerFrom := (*c).fetch() % uint16(len(c.registerNames))
			address := (*c).fetch16()
			value := (*c).registers[registerFrom]
			(*c).memory[address] = value
			return
		}
		// Move memory to register
	case MOV_MEM_REG:
		{
			address := (*c).fetch16()
			registerTo := (*c).fetch() % uint16(len(c.registerNames))
			value := (*c).memory[address]
			(*c).registers[registerTo] = value
			return
		}
	// Add register to the register (we add values in r1 and r2 and save in acc register)
	case ADD_REG_REG:
		{
			r1 := (*c).fetch()
			r2 := (*c).fetch()

			registerValue1 := c.registers[r1]
			registerValue2 := c.registers[r2]

			(*c).setRegister("acc", registerValue1+registerValue2)
			return
		}
	}
}

func (c *CPU) step() {
	instruction := (*c).fetch()
	(*c).execute(instruction)
}

/*
	EXECUTION STEP {
		we have 3 instructions at the time of writing this:
		MOV_LIT_R1  = 0x10
		MOV_LIT_R2  = 0x11
		ADD_REG_REG = 0x12,
		This will be the progression of the state,
		when we call step() 3 times in `main.go` to execute all 3 implemented instruction this time
		ip=5, mem[5]=2, acc=0 => ip=6, mem[6]=3, acc=0 => ip=7, mem[6]=3, acc=0x1234+0xabcd

			case ADD_REG_REG:
					{
						r1 := (*c).fetch() //2
						r2 := (*c).fetch() //3

						registerValue1 := c.registers[r1] // 0x1234
						registerValue2 := c.registers[r2] // 0xabcd

						(*c).setRegister("acc", registerValue1+registerValue2)
						return
					}
	}


*/
