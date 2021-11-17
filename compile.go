package parser

import (
	"fmt"
	"math"
)

type Bytecode int

const (
	AddBytecode = iota //Add two memory locations and save to third memory location
	SubBytecode        //Subtract two memory locations and save to third memory location
	MulBytecode
	DivBytecode
	PowBytecode //Raise first location to power of second location and save it at third
	CosBytecode //Take the cosine of the first locartion and save it to the second
	SinBytecode
)

type MemoryManager struct {
	bc        []Bytecode
	constants []float64
	//Guide for which places to fill with which variables
	//Now there is only one memory bank
	varLocations map[string]int
}

func NewMemoryManager() MemoryManager {
	return MemoryManager{
		constants:    []float64{},
		varLocations: map[string]int{},
	}
}
func (mm *MemoryManager) AddBytecode(bc []Bytecode) {
	mm.bc = append(mm.bc, bc...)
}

func (mm *MemoryManager) AddVariable(name string) int {
	//Add Variable to memory and return the index to it
	if i, ok := mm.varLocations[name]; ok {
		return i
	} else {
		index := mm.GetResultSpace()
		mm.varLocations[name] = index
		return index
	}
}
func (mm *MemoryManager) AddConstant(v float64) int {
	//Add Constant to memory and return the index to it
	i := len(mm.constants)
	mm.constants = append(mm.constants, v)
	return i
}
func (mm *MemoryManager) GetResultSpace() int {
	//Add place in memory to store a result and return the index to it
	i := len(mm.constants)
	mm.constants = append(mm.constants, -8008)
	return i
}

func CompileExpression(e Expression) func(vs map[string]float64) float64 {
	mm := NewMemoryManager()
	lastResIndex := e.Compile(&mm)

	compiledFunc := func(vs map[string]float64) float64 {
		//Copy over code and constants
		code := mm.bc
		consts := mm.constants
		//save vs to mem bank
		for k, v := range mm.varLocations {
			consts[v] = vs[k]
		}
		//Execute code
		for i := 0; i < len(code); {
			ins := code[i]
			switch ins {
			case AddBytecode:
				Ai := code[i+1]
				Bi := code[i+2]

				Ri := code[i+3]
				consts[Ri] = consts[Ai] + consts[Bi]
				i += 4
			case SubBytecode:
				Ai := code[i+1]
				Bi := code[i+2]

				Ri := code[i+3]
				consts[Ri] = consts[Ai] - consts[Bi]
				i += 4
			case MulBytecode:
				Ai := code[i+1]
				Bi := code[i+2]

				Ri := code[i+3]
				consts[Ri] = consts[Ai] * consts[Bi]
				i += 4

			case DivBytecode:
				Ai := code[i+1]
				Bi := code[i+2]

				Ri := code[i+3]
				consts[Ri] = consts[Ai] / consts[Bi]
				i += 4
			case PowBytecode:
				Ai := code[i+1]
				Bi := code[i+2]

				Ri := code[i+3]
				consts[Ri] = math.Pow(consts[Ai], consts[Bi])
				i += 4
			case CosBytecode:
				Ai := code[i+1]

				Ri := code[i+2]
				consts[Ri] = math.Cos(consts[Ai])
				i += 3

			case SinBytecode:
				Ai := code[i+1]
				Ri := code[i+2]
				consts[Ri] = math.Sin(consts[Ai])
				i += 3

			default:
				//This should really never happen but just go to next instruction
				i++
			}

		}
		fmt.Println(consts)
		return consts[lastResIndex]
	}
	return compiledFunc
}
