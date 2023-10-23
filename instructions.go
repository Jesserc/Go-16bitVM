package main

const (
	MOV_LIT_REG = 0x10 // move literal to register
	MOV_REG_REG = 0x11 // move register to register
	MOV_REG_MEM = 0x12 // move register to register
	MOV_MEM_REG = 0x13 // move register to register
	// MOV_LIT_MEM = 0x11 // move register to register
	ADD_REG_REG = 0x14 // add register (r1) to register (r2)
	JUMP_NOT_EQ = 0x15
)
