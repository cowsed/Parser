package parser

import (
	"math"
)

//Bytecode is an instruction type for the interpreter
type Bytecode int

//List of Bytecodes
const (
	AddBytecode = iota //Add two memory locations and save to third memory location
	SubBytecode        //Subtract two memory locations and save to third memory location
	MulBytecode
	DivBytecode
	PowBytecode //Raise first location to power of second location and save it at third
	CosBytecode //Take the cosine of the first locartion and save it to the second
	SinBytecode
	LNBytecode
)

//MemoryManager keeps track of constants, variables and working memory
//Also holds place for instructions
type MemoryManager struct {
	bc        []Bytecode
	constants []float64
	//Guide for which places to fill with which variables
	varLocations map[string]int
}

//NewMemoryManager returns a new default memory manager
func NewMemoryManager() MemoryManager {
	return MemoryManager{
		constants:    []float64{},
		varLocations: map[string]int{},
	}
}

//AddBytecode adds a bytecode to the instructions
func (mm *MemoryManager) AddBytecode(bc []Bytecode) {
	mm.bc = append(mm.bc, bc...)
}

//AddVariable adds a variable and tracks it to be set at execution time
func (mm *MemoryManager) AddVariable(name string) int {
	//Add Variable to memory and return the index to it

	//If its already here
	if i, ok := mm.varLocations[name]; ok {
		return i
	}
	//if its new
	index := mm.GetResultSpace()
	mm.varLocations[name] = index
	return index

}

//AddConstant adds a constant into the memory
func (mm *MemoryManager) AddConstant(v float64) int {
	//Add Constant to memory and return the index to it
	i := len(mm.constants)
	mm.constants = append(mm.constants, v)
	return i
}

//GetResultSpace returns a location for the next variable or constant to be put in
func (mm *MemoryManager) GetResultSpace() int {
	//Add place in memory to store a result and return the index to it
	i := len(mm.constants)
	mm.constants = append(mm.constants, -8008)
	return i
}

//CompileExpression takes an expression and turns it into bytecode
func CompileExpression(e Expression) func(vs map[string]float64) float64 {
	mm := NewMemoryManager()
	lastResIndex := e.Compile(&mm)
	vars := make([]string, len(mm.varLocations))
	i := 0
	for k := range mm.varLocations {
		vars[i] = k
		i++
	}

	compiledFunc := func(vs map[string]float64) float64 {
		//Copy over code and constants
		code := mm.bc
		consts := mm.constants
		//save vs to mem bank
		for _, k := range vars {
			index := mm.varLocations[k]
			val := vs[k]
			consts[index] = val

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
			case LNBytecode:
				Ai := code[i+1]
				Ri := code[i+2]
				consts[Ri] = math.Log(consts[Ai])
				i += 3

			default:
				//This should really never happen but just go to next instruction
				i++
			}

		}
		return consts[lastResIndex]
	}
	return compiledFunc
}
