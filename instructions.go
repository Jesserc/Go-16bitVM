package main

const (
	MOV_LIT_REG = 0x10 // move literal to register
	MOV_REG_REG = 0x11 // move register to register
	MOV_REG_MEM = 0x12 // move register to register
	MOV_MEM_REG = 0x13 // move register to register
	ADD_REG_REG = 0x14 // add register (r1) to register (r2)
	JUMP_NOT_EQ = 0x15 // set IP to a particular location if value of ACC isn't equal to a literal to be compared with
	PUSH_LIT    = 0x17 // push a literal value to the stack
	PUSH_REG    = 0x18 // push the value of a register to the stack
	POP         = 0x1a // pop a value from stack to register. stack will still have the value but SP will change, i.e the value can be overridden but is left there intentionally to save computational resource
)
