package parser

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

func JitCompileExpression(e Expression) func(vs map[string]float64) float64 {
	mm := NewMemoryManager()
	e.Compile(&mm)

	f := JitCompile(&mm)
	return f
}

var header = []uint8{0x48, 0x83, 0xec, 0x18, 0x48, 0x89, 0x6c, 0x24, 0x10, 0x48, 0x8d, 0x6c, 0x24, 0x10, 0x48, 0x89,
	0x44, 0x24, 0x20, 0x48, 0x85, 0xdb, 0x76, 0x43}
var footer = []uint8{0x48, 0x8b, 0x6c, 0x24,
	0x10, 0x48, 0x83, 0xc4, 0x18, 0x90, 0xc3, 0xb8, 0x02, 0x00, 0x00, 0x00, 0x48, 0x89, 0xd9, 0xe8}

//Takes a completed memory manager
func JitCompile(mm *MemoryManager) func(vs map[string]float64) float64 {
	//Track of which vars are needed
	vars := make([]string, len(mm.varLocations))
	i := 0
	for k := range mm.varLocations {
		vars[i] = k
		i++
	}

	//Create operating memory (memory in which operations are performed. only parts can be overwritten, others must stay unchanged for the function to wo0rk muiltiple times)
	operating := make([]float64, len(mm.constants))
	for i := range operating {
		operating[i] = mm.constants[i]
	}
	code := []uint8{}
	code = append(code, header...)
	//Translate intermediate representation into x86 Assembly
	/*
		for i := 0; i < len(mm.bc); {
			ins := mm.bc[i]
			switch ins {
			case AddBytecode:
				Ai := mm.bc[i+1]
				Bi := mm.bc[i+2]
				Ri := mm.bc[i+3]

				//float64s are 8 byte long
				AiMem := uint8(Ai) * 8
				BiMem := uint8(Bi) * 8
				RiMem := uint8(Ri) * 8

				tempCode := []uint8{
					//Load A from memory
					0xf2, 0x0f, 0x10, 0x40, AiMem,

					//Add B from memory to A
					0xf2, 0x0f, 0x58, 0x40, BiMem,

					//Save to R in memory
					0xf2, 0x0f, 0x11, 0x40, RiMem,
				}
				code = append(code, tempCode...)
				i += 4
			case SubBytecode:
				Ai := mm.bc[i+1]
				Bi := mm.bc[i+2]
				Ri := mm.bc[i+3]

				//float64s are 8 byte long
				AiMem := uint8(Ai) * 8
				BiMem := uint8(Bi) * 8
				RiMem := uint8(Ri) * 8

				tempCode := []uint8{
					//Load A from memory
					0xf2, 0x0f, 0x10, 0x40, AiMem,

					//Sub B from memory from A
					0xf2, 0x0f, 0x5c, 0x40, BiMem,

					//Save to R in memory
					0xf2, 0x0f, 0x11, 0x40, RiMem,
				}
				code = append(code, tempCode...)
				i += 4
			default:
				fmt.Println("this shouldnt happen. Unrecognized function:", ins)
				return nil
			}

		}
	*/

	code = append(code, footer...)

	AsmFunction := MakeMathFunc(code)
	if AsmFunction == nil {
		panic(fmt.Errorf("should not happen. Nil function"))
	}
	fmt.Println(&AsmFunction)
	return func(vs map[string]float64) float64 {
		//Place variables in operating memory
		for _, k := range vars {
			index := mm.varLocations[k]
			val := vs[k]
			operating[index] = val

		}
		fmt.Println("operating mem", operating)
		fmt.Println("executing")
		res := AsmFunction(operating)
		fmt.Println("result")
		fmt.Println(res)
		return res
	}
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
