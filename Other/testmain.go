package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func PrintHexBytes(s []byte) {
	fmt.Print("[")
	for i := range s {
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Println("]")
}

func main() {
	mathFunction2 := []uint8{
		//Setup stuff
		0x48, 0x83, 0xec, 0x18, 0x48, 0x89, 0x6c, 0x24, 0x10, 0x48, 0x8d, 0x6c, 0x24, 0x10, 0x48, 0x89,
		0x44, 0x24, 0x20, 0x48, 0x85, 0xdb, 0x76, 0x43,

		//data[3]+data[0]+data[1]
		0xf2, 0x0f, 0x10, 0x00, //MOVESD

		//Returning stuff
		0x48, 0x8b, 0x6c, 0x24, 0x10,
		0x48, 0x83, 0xc4, 0x18, 0x90, 0xc3, 0xb8, 0x02, 0x00, 0x00, 0x00, 0x48, 0x89, 0xd9, 0xe8,
	}
	data := []float64{2, 2, 3, -8008, -8008}

	f := MakeMathFunc(mathFunction2)

	fmt.Println(f(data))
}
func MakeMathFunc(mathFunction []uint8) func([]float64) float64 {
	type floatFunc func([]float64) float64

	fmt.Println("function length", len(mathFunction))
	if len(mathFunction) > 128 {
		panic(fmt.Errorf("function too long for memory alloted"))
	}
	PrintHexBytes(mathFunction)

	executablePrintFunc, err := syscall.Mmap(
		-1,
		0,
		128,
		syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC,
		syscall.MAP_PRIVATE|syscall.MAP_ANONYMOUS)
	if err != nil {
		fmt.Printf("mmap err: %v", err)
	}

	copy(executablePrintFunc, mathFunction) ///When going back this is where it gets switched out for debug function

	PrintHexBytes(executablePrintFunc)

	unsafePrintFunc := (uintptr)(unsafe.Pointer(&executablePrintFunc))
	function := *(*floatFunc)(unsafe.Pointer(&unsafePrintFunc))
	return function
}
