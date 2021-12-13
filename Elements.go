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
	Derive(wrt string) Expression
	Simplify() Expression
}

type Siner struct {
	A Expression
}

func (s Siner) Derive(wrt string) Expression {

	return Multiplier{
		A: Coser{s.A},
		B: s.A.Derive(wrt),
	}
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

func (c Coser) Derive(wrt string) Expression {
	//-cos(x)*dx
	return Multiplier{
		A: Multiplier{
			A: Constant{
				Value: -1,
			},
			B: Siner{
				A: c.A,
			},
		},
		B: c.A.Derive(wrt),
	}
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

func (a Adder) Derive(wrt string) Expression {
	return Adder{
		A: a.A.Derive(wrt),
		B: a.B.Derive(wrt),
	}.Simplify()
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

func (s Subtractor) Derive(wrt string) Expression {
	return Subtractor{
		A: s.A.Derive(wrt),
		B: s.B.Derive(wrt),
	}.Simplify()
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

func (m Multiplier) Derive(wrt string) Expression {
	return Adder{
		A: Multiplier{
			A: m.A,
			B: m.B.Derive(wrt),
		},
		B: Multiplier{
			A: m.A.Derive(wrt),
			B: m.A,
		},
	}.Simplify()
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

func (d Divider) Derive(wrt string) Expression {
	return Divider{
		A: Subtractor{
			A: Multiplier{
				A: d.B,
				B: d.A.Derive(wrt),
			},
			B: Multiplier{
				A: d.A,
				B: d.B.Derive(wrt),
			},
		},
		B: Powerer{
			Base:     d.B,
			Exponent: Constant{2},
		},
	}.Simplify()
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

type NaturalLogger struct {
	A Expression
}

func (n NaturalLogger) Evaluate(vars map[string]float64) float64 {
	return math.Log(n.A.Evaluate(vars))
}
func (n NaturalLogger) Derive(wrt string) Expression {
	return Multiplier{
		A: Divider{
			A: Constant{1},
			B: n.A,
		},
		B: n.A.Derive(wrt),
	}.Simplify()
}
func (n NaturalLogger) Compile(mm *MemoryManager) int {
	aResult := n.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{LNBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}
func (n NaturalLogger) Latex() string {
	return fmt.Sprintf("\\ln{%s}", n.A.Latex())
}

func (n NaturalLogger) String() string {
	return fmt.Sprintf("ln(%s)", n.A.String())
}

type Powerer struct {
	Base, Exponent Expression
}

//https://www.youtube.com/watch?v=SUxcFxM65Ho
func (p Powerer) Derive(wrt string) Expression {
	return Multiplier{
		A: Powerer{
			Base:     p.Base,
			Exponent: p.Exponent,
		},
		B: Adder{
			A: Divider{
				A: Multiplier{
					A: p.Exponent,
					B: p.Base.Derive(wrt),
				},
				B: p.Base,
			},
			B: Multiplier{
				A: p.Exponent.Derive(wrt),
				B: NaturalLogger{
					A: p.Base,
				},
			},
		},
	}.Simplify()
}

func (p Powerer) Evaluate(vars map[string]float64) float64 {
	return math.Pow(p.Base.Evaluate(vars), p.Exponent.Evaluate(vars))
}
func (p Powerer) String() string {
	return "(" + p.Base.String() + " ^ " + p.Exponent.String() + ")"
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

func (c Constant) Derive(wrt string) Expression {
	return Constant{0}
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
func (v Variable) Derive(wrt string) Expression {
	if v.Symbol == wrt {
		return Constant{1}
	} else {
		return Constant{0}
	}
}
