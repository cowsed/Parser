package parser

import (
	"fmt"
	"math"
)

//Expression is the simplest possible part of an mathematical expression
type Expression interface {
	Evaluate(vars map[string]float64) float64
	String() string
	Latex() string
	Compile(mm *MemoryManager) int
	Derive(wrt string) Expression
	Simplify() Expression
}

//Siner takes the sine of its value
type Siner struct {
	A Expression
}

//Derive takes the derivative of the sin(A) with respect to wrt
func (s Siner) Derive(wrt string) Expression {

	return Multiplier{
		A: Coser{s.A},
		B: s.A.Derive(wrt),
	}
}

//Evaluate evaluates sin of A
func (s Siner) Evaluate(vars map[string]float64) float64 {
	return math.Sin(s.A.Evaluate(vars))
}

//String returns a string representation of sin(a)
func (s Siner) String() string {
	return "sin(" + s.A.String() + ")"
}

//Latex returns the latex representation of sin(a)
func (s Siner) Latex() string {
	return "sin(" + s.A.Latex() + ")"
}

//Compile compiles sin into bytecode
func (s Siner) Compile(mm *MemoryManager) int {
	aResult := s.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{SinBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Coser takes the cosine of its value
type Coser struct {
	A Expression
}

//Derive takes the derivative of the cos(A) with respect to wrt
func (c Coser) Derive(wrt string) Expression {
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

//Evaluate evaluates cosine(a)
func (c Coser) Evaluate(vars map[string]float64) float64 {
	return math.Cos(c.A.Evaluate(vars))
}

//String returns a string representation of cos(a)
func (c Coser) String() string {
	return "cos(" + c.A.String() + ")"
}

//Latex returns a latex representation of cos(a)
func (c Coser) Latex() string {
	return "cos(" + c.A.Latex() + ")"
}

//Compile compiles cos(a) to bytecode
func (c Coser) Compile(mm *MemoryManager) int {
	aResult := c.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{CosBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Adder adds its two parameters
type Adder struct {
	A, B Expression
}

//Derive takes the derivative of A + B.
func (a Adder) Derive(wrt string) Expression {
	return Adder{
		A: a.A.Derive(wrt),
		B: a.B.Derive(wrt),
	}.Simplify()
}

//Evaluate evaluates a+b
func (a Adder) Evaluate(vars map[string]float64) float64 {
	return a.A.Evaluate(vars) + a.B.Evaluate(vars)
}

//Compile compiles A+B to bytecode
func (a Adder) Compile(mm *MemoryManager) int {
	//Add instructions to memory manager as well as index to result
	aResult := a.A.Compile(mm)
	bResult := a.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{AddBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//String returns a string representation of A+B
func (a Adder) String() string {
	return "(" + a.A.String() + " + " + a.B.String() + ")"
}

//Latex returns a latex representation of A+B
func (a Adder) Latex() string {
	return a.A.Latex() + " + " + a.B.Latex()
}

//Subtractor subtracts its two parameters
type Subtractor struct {
	A, B Expression
}

//Derive takes the derivative of A - B.
func (s Subtractor) Derive(wrt string) Expression {
	return Subtractor{
		A: s.A.Derive(wrt),
		B: s.B.Derive(wrt),
	}.Simplify()
}

//Evaluate evaluates A-B
func (s Subtractor) Evaluate(vars map[string]float64) float64 {
	return s.A.Evaluate(vars) - s.B.Evaluate(vars)
}

//String returns a string representation of A-B
func (s Subtractor) String() string {
	return "(" + s.A.String() + " - " + s.B.String() + ")"
}

//Latex returns a latex representation of A-B
func (s Subtractor) Latex() string {
	return s.A.Latex() + " - " + s.B.Latex()
}

//Compile compiles A-B to bytecode
func (s Subtractor) Compile(mm *MemoryManager) int {
	aResult := s.A.Compile(mm)
	bResult := s.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{SubBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Multiplier multiplies its two parameters
type Multiplier struct {
	A, B Expression
}

//Derive takes the derivative of A * B.
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

//Evaluate evaluates A*B
func (m Multiplier) Evaluate(vars map[string]float64) float64 {
	return m.A.Evaluate(vars) * m.B.Evaluate(vars)
}

//String returns a string representation of A*B
func (m Multiplier) String() string {
	return "(" + m.A.String() + " * " + m.B.String() + ")"
}

//Latex returns a latex representation of A*B
func (m Multiplier) Latex() string {
	return m.A.Latex() + " \\times " + m.B.Latex()
}

//Compile compiles A*B to bytecode
func (m Multiplier) Compile(mm *MemoryManager) int {
	aResult := m.A.Compile(mm)
	bResult := m.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{MulBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Divider divides its parameters
type Divider struct {
	A, B Expression
}

//Derive takes the derivative of A / B.
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

//Evaluate evaluates A/B
func (d Divider) Evaluate(vars map[string]float64) float64 {
	return d.A.Evaluate(vars) / d.B.Evaluate(vars)
}

//String returns a string representation of A/B
func (d Divider) String() string {
	return "(" + d.A.String() + " / " + d.B.String() + ")"
}

//Latex returns a latex representation of A/B
func (d Divider) Latex() string {
	return "\\frac{" + d.A.Latex() + "}{" + d.B.Latex() + "}"
}

//Compile compiles A/B to bytecode
func (d Divider) Compile(mm *MemoryManager) int {
	aResult := d.A.Compile(mm)
	bResult := d.B.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{DivBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//NaturalLogger takes the natural log of A
type NaturalLogger struct {
	A Expression
}

//Evaluate evaluates ln(A)
func (n NaturalLogger) Evaluate(vars map[string]float64) float64 {
	return math.Log(n.A.Evaluate(vars))
}

//Derive takes the derivative of ln(A).
func (n NaturalLogger) Derive(wrt string) Expression {
	return Multiplier{
		A: Divider{
			A: Constant{1},
			B: n.A,
		},
		B: n.A.Derive(wrt),
	}.Simplify()
}

//Compile compiles ln(A) to bytecode
func (n NaturalLogger) Compile(mm *MemoryManager) int {
	aResult := n.A.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{LNBytecode, Bytecode(aResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Latex returns the latex representation of ln(a)
func (n NaturalLogger) Latex() string {
	return fmt.Sprintf("\\ln{%s}", n.A.Latex())
}

//String returns the string representation of ln(a)
func (n NaturalLogger) String() string {
	return fmt.Sprintf("ln(%s)", n.A.String())
}

//Powerer raises base to exponent
type Powerer struct {
	Base, Exponent Expression
}

//Derive takes the derivative of A^B
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

//Evaluate evaluates base^power
func (p Powerer) Evaluate(vars map[string]float64) float64 {
	return math.Pow(p.Base.Evaluate(vars), p.Exponent.Evaluate(vars))
}

//String returns a string representation of base^exponent
func (p Powerer) String() string {
	return "(" + p.Base.String() + " ^ " + p.Exponent.String() + ")"
}

//Latex returns a latex representation of base^exponent
func (p Powerer) Latex() string {
	return p.Base.Latex() + "^{" + p.Exponent.Latex() + "}"
}

//Compile compiles it to bytecode
func (p Powerer) Compile(mm *MemoryManager) int {
	aResult := p.Base.Compile(mm)
	bResult := p.Exponent.Compile(mm)
	myResultIndex := mm.GetResultSpace()
	mm.AddBytecode([]Bytecode{PowBytecode, Bytecode(aResult), Bytecode(bResult), Bytecode(myResultIndex)})
	return myResultIndex
}

//Constant holds a numeric constant
type Constant struct {
	Value float64
}

//Derive takes the derivative of a constant (0)
func (c Constant) Derive(wrt string) Expression {
	return Constant{0}
}

//String returns the string representation of the number
func (c Constant) String() string {
	return fmt.Sprintf("%g", c.Value)
}

//Latex returns the latex representation of the number
func (c Constant) Latex() string {
	return fmt.Sprintf("%g", c.Value)
}

//Evaluate returns the value of the constant
func (c Constant) Evaluate(vars map[string]float64) float64 {
	return c.Value
}

//Compile creates a place to store the constant
func (c Constant) Compile(mm *MemoryManager) int {
	i := mm.AddConstant(c.Value)
	return i
}

//Variable holds a variable in an equation
type Variable struct {
	Symbol string
}

//String returns the string representation of the variable
func (v Variable) String() string {
	return v.Symbol
}

//Latex returns the latex representation of the variable
func (v Variable) Latex() string {
	return v.Symbol
}

//Evaluate takes the variable and looks up its numeric value at the time of evaluation
func (v Variable) Evaluate(vars map[string]float64) float64 {
	return vars[v.Symbol]
}

//Compile creates a space for the variable to stay and records where they go to fill in later
func (v Variable) Compile(mm *MemoryManager) int {
	i := mm.AddVariable(v.Symbol)
	return i
}

//Derive takes the derivative of the variable
func (v Variable) Derive(wrt string) Expression {
	if v.Symbol == wrt {
		return Constant{1}
	}
	//Not with respect to this variable. Variable is basically constant
	return Constant{0}

}
