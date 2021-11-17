package parser

import (
	"fmt"
	"math"
)

type Expression interface {
	Evaluate(vars map[string]float64) float64
	String() string
	Latex() string
	Compile(mm *MemoryManager) int
}

type Siner struct {
	A Expression
}

func (s Siner) Evaluate(vars map[string]float64) float64 {
	return math.Sin(s.A.Evaluate(vars))
}
func (s Siner) String() string {
	return "sin(" + s.A.String() + ")"
}
func (s Siner) Latex() string {
	return "sin(" + s.A.Latex() + ")"
}
func (s Siner) Compile(mm *MemoryManager) int {
	aResult := s.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{SinBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Coser struct {
	A Expression
}

func (c Coser) Evaluate(vars map[string]float64) float64 {
	return math.Cos(c.A.Evaluate(vars))
}
func (c Coser) String() string {
	return "cos(" + c.A.String() + ")"
}
func (c Coser) Latex() string {
	return "cos(" + c.A.Latex() + ")"
}
func (c Coser) Compile(mm *MemoryManager) int {
	aResult := c.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{CosBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Adder struct {
	A, B Expression
}

func (a Adder) Evaluate(vars map[string]float64) float64 {
	return a.A.Evaluate(vars) + a.B.Evaluate(vars)
}
func (a Adder) Compile(mm *MemoryManager) int { //Add instructions to memory manager as well as index to result
	aResult := a.A.Compile(mm)
	bResult := a.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{AddBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}
func (a Adder) String() string {
	return "(" + a.A.String() + " + " + a.B.String() + ")"
}
func (a Adder) Latex() string {
	return a.A.Latex() + " + " + a.B.Latex()
}

type Subtractor struct {
	A, B Expression
}

func (s Subtractor) Evaluate(vars map[string]float64) float64 {
	return s.A.Evaluate(vars) - s.B.Evaluate(vars)
}
func (s Subtractor) String() string {
	return "(" + s.A.String() + " - " + s.B.String() + ")"
}
func (s Subtractor) Latex() string {
	return s.A.Latex() + " - " + s.B.Latex()
}
func (s Subtractor) Compile(mm *MemoryManager) int {
	aResult := s.A.Compile(mm)
	bResult := s.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{SubBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Multiplier struct {
	A, B Expression
}

func (m Multiplier) Evaluate(vars map[string]float64) float64 {
	return m.A.Evaluate(vars) * m.B.Evaluate(vars)
}
func (m Multiplier) String() string {
	return "(" + m.A.String() + " * " + m.B.String() + ")"
}
func (m Multiplier) Latex() string {
	return m.A.Latex() + " \\times " + m.B.Latex()
}

func (m Multiplier) Compile(mm *MemoryManager) int {
	aResult := m.A.Compile(mm)
	bResult := m.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{MulBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Divider struct {
	A, B Expression
}

func (d Divider) Evaluate(vars map[string]float64) float64 {
	return d.A.Evaluate(vars) / d.B.Evaluate(vars)
}
func (s Divider) String() string {
	return "(" + s.A.String() + " / " + s.B.String() + ")"
}
func (s Divider) Latex() string {
	return "\\frac{" + s.A.Latex() + "}{" + s.B.Latex() + "}"
}
func (d Divider) Compile(mm *MemoryManager) int {
	aResult := d.A.Compile(mm)
	bResult := d.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{DivBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Powerer struct {
	Base, Exponent Expression
}

func (p Powerer) Evaluate(vars map[string]float64) float64 {
	return math.Pow(p.Base.Evaluate(vars), p.Exponent.Evaluate(vars))
}
func (p Powerer) String() string {
	return p.Base.String() + " ^ " + p.Exponent.String()
}
func (p Powerer) Latex() string {
	return p.Base.Latex() + "^{" + p.Exponent.Latex() + "}"
}
func (p Powerer) Compile(mm *MemoryManager) int {
	aResult := p.Base.Compile(mm)
	bResult := p.Exponent.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{PowBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

type Constant struct {
	Value float64
}

func (c Constant) String() string {
	return fmt.Sprintf("%g", c.Value)
}
func (c Constant) Latex() string {
	return fmt.Sprintf("%g", c.Value)
}
func (c Constant) Evaluate(vars map[string]float64) float64 {
	return c.Value
}
func (c Constant) Compile(mm *MemoryManager) int {
	i := mm.AddConstant(c.Value)
	return i
}

type Variable struct {
	Symbol string
}

func (v Variable) String() string {
	return v.Symbol
}
func (v Variable) Latex() string {
	return v.Symbol
}

func (v Variable) Evaluate(vars map[string]float64) float64 {
	return vars[v.Symbol]
}

func (v Variable) Compile(mm *MemoryManager) int {
	i := mm.AddVariable(v.Symbol)
	return i
}
