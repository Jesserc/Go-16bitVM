package main

type Memory []uint16

func CreateMemory(length int) Memory {
	return make(Memory, length)
}
