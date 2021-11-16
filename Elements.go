package parser

import (
	"fmt"
	"math"
)

type Expression interface {
	Evaluate(vars map[string]float64) float64
	String() string
	Latex() string
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

type Adder struct {
	A, B Expression
}

func (a Adder) Evaluate(vars map[string]float64) float64 {
	return a.A.Evaluate(vars) + a.B.Evaluate(vars)
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

type Multiplier struct {
	A, B Expression
}

func (m Multiplier) Evaluate(vars map[string]float64) float64 {
	return m.A.Evaluate(vars) * m.B.Evaluate(vars)
}
func (s Multiplier) String() string {
	return "(" + s.A.String() + " * " + s.B.String() + ")"
}
func (s Multiplier) Latex() string {
	return s.A.Latex() + " \\times " + s.B.Latex()
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

type Powerer struct {
	Base, Exponent Expression
}

func (d Powerer) Evaluate(vars map[string]float64) float64 {
	return math.Pow(d.Base.Evaluate(vars), d.Exponent.Evaluate(vars))
}
func (s Powerer) String() string {
	return s.Base.String() + " ^ " + s.Exponent.String()
}
func (s Powerer) Latex() string {
	return s.Base.Latex() + "^{" + s.Exponent.Latex() + "}"
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
