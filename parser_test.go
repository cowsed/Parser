package parser

import (
	"fmt"
	"math"
	"testing"
)

func TestCompile(t *testing.T) {
	expr := "2+2-5"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{})
	if res != float64(2+2-5) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(2+2-5), res)
	}
}
func TestCompile2(t *testing.T) {
	expr := "2+2*a"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{"a": 3})
	if res != float64(2+2*3) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(2+2*3), res)
	}
}
func TestCompile3(t *testing.T) {
	expr := "a+2*a"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{"a": 3})
	if res != float64(3+2*3) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(3+2*3), res)
	}
}

func TestCompile5(t *testing.T) {
	expr := "sin(x*y/b)*y+cos(a*x-y)"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println("Evaluated E: ", e.Evaluate(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1}))
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1})
	if res != float64(math.Sin(.1*.1/0.7345345)*.1+math.Cos(0.65343*.1-.1)) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(math.Sin(.1*.1/0.7345345)*.1+math.Cos(0.65343*.1-.1)), res)
	}
}

func TestCompile6(t *testing.T) {
	expr := "sin(y/b)" //Something about this expression is cringe
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println(e.String())
	fmt.Println("Evaluated E: ", e.Evaluate(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1}))
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1})
	if res != float64(math.Sin(.1/0.7345345)) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(math.Sin(.1/0.7345345)), res)
	}
}

func TestCompile7(t *testing.T) {
	expr := "cos(a*x-y)"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println("Evaluated E: ", e.Evaluate(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1}))
	expComp := CompileExpression(e)
	res := expComp(map[string]float64{"a": 0.65343, "b": 0.7345345, "x": 0.1, "y": 0.1})
	if res != float64(math.Cos(0.65343*.1-.1)) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(math.Cos(0.65343*.1-.1)), res)
	}
}

func TestParse(t *testing.T) {
	expr := "2+2-3"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println(e.String())
	res := e.Evaluate(map[string]float64{})
	if res != float64(2+2-3) {
		t.Errorf("Had %s, expected result: %g. Actual result: %g", expr, float64(2+2-3), res)
	}
}

func TestParse2(t *testing.T) {
	expr := "2+2-3*4/2"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println(e.String())
	res := e.Evaluate(map[string]float64{})
	if res != float64(2+2-3*4/2) {
		t.Errorf("Had %s parsed to %s, expected result: %g. Actual result: %g", expr, e.String(), float64(2+2-3*4/2), res)
	}
}
func TestParse3(t *testing.T) {
	expr := "3^4"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println(e.String())
	res := e.Evaluate(map[string]float64{})
	if res != float64(math.Pow(3, 4)) {
		t.Errorf("Had %s parsed to %s, expected result: %g. Actual result: %g", expr, e.String(), math.Pow(3, 4), res)
	}
}

func TestParse4(t *testing.T) {
	expr := "a+x^2"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("Unknown parse error")
	}
	fmt.Println(e.String())
	res := e.Evaluate(map[string]float64{"x": 3, "a": 2})
	if res != float64(2+(3*3)) {
		t.Errorf("Had %s parsed to %s, expected result: %g. Actual result: %g", expr, e.String(), float64(2+(3*3)), res)
	}
}
func TestParse5(t *testing.T) {
	expr := "a*3+1/2+a+x^(a+2)+1"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("parse error")
	}
	fmt.Println("Latex", e.Latex())
	res := e.Evaluate(map[string]float64{"x": 3, "a": 2})
	if res != float64(2*3+1.0/2.0+2+math.Pow(3, 2+2)+1) {
		t.Errorf("Had %s parsed to %s, expected result: %g. Actual result: %g", expr, e.String(), float64(2*3+1.0/2.0+2+math.Pow(3, 2+2)+1), res)
	}
}

func TestParse6(t *testing.T) {
	expr := "2+3*sin(x)"
	e, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("error")
		t.Errorf(err.Error())
	}
	if e == nil {
		t.Errorf("parse error")
	}
	fmt.Println("Latex", e.Latex())
	fmt.Println("String", e.String())

	res := e.Evaluate(map[string]float64{"x": 0})
	if res != float64(2+3*math.Sin(0)) {
		t.Errorf("Had %s parsed to %s, expected result: %g. Actual result: %g", expr, e.String(), float64(2+3*math.Sin(0)), res)
	}
}

func TestParseToPostfix(t *testing.T) {
	expr := "3*4+2"
	tokens, err := tokenize(expr)
	if err != nil {
		t.Errorf(err.Error())
	}
	pf, err := MakePostFix(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(pf)
	wanted := []Token{
		{
			Type:  NumberType,
			Value: "3",
		},
		{
			Type:  NumberType,
			Value: "4",
		},
		{
			Type:  OperatorType,
			Value: "*",
		},
		{
			Type:  NumberType,
			Value: "2",
		},
		{
			Type:  OperatorType,
			Value: "+",
		},
	}
	for i := 0; i < len(pf); i++ {
		if pf[i] != wanted[i] {
			t.Errorf("Wanted %v, got %v at index %d", wanted[i], pf[i], i)
		}
	}

}

func TestPostfix(t *testing.T) {
	tokens := []Token{
		{
			Type:  NumberType,
			Value: "2",
		},
		{
			Type:  OperatorType,
			Value: "+",
		},
		{
			Type:  NumberType,
			Value: "2",
		},
	}
	pf, err := MakePostFix(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}
	wanted := []Token{
		{
			Type:  NumberType,
			Value: "2",
		},
		{
			Type:  NumberType,
			Value: "2",
		},
		{
			Type:  OperatorType,
			Value: "+",
		},
	}
	fmt.Println("Infix", tokens)
	fmt.Println("Postfix", pf)
	for i := 0; i < len(pf); i++ {
		if pf[i] != wanted[i] {
			t.Errorf("Wanted %v, got %v at index %d", wanted[i], pf[i], i)
		}
	}
}

func TestAdder(t *testing.T) {
	got := Adder{Constant{1}, Constant{3}}.Evaluate(map[string]float64{})
	if got != 4 {
		t.Errorf("1-3 = %g; want -2", got)
	}
}
func TestSubtractor(t *testing.T) {
	got := Subtractor{Constant{1}, Constant{3}}.Evaluate(map[string]float64{})
	if got != -2 {
		t.Errorf("1+3 = %g; want 3", got)
	}
}
func TestStack(t *testing.T) {
	s := NewTokenStack()
	s.Push(Token{
		Type:  NumberType,
		Value: "2",
	})
	s.Push(Token{
		Type:  NumberType,
		Value: "2",
	})
	s.Push(Token{
		Type:  OperatorType,
		Value: "+",
	})

	a, _ := s.Pop()
	b, _ := s.Pop()
	c, _ := s.Pop()
	fmt.Println(a, b, c, "outputs")
	d, err := s.Pop()
	if err != nil {
		fmt.Println(d)
	} else {
		fmt.Println("popped off the end")
	}
}

func TestTokenize(t *testing.T) {
	got, err := tokenize("aa*x^2-bx+c")
	fmt.Println(got, err)
}

func TestTokenize2(t *testing.T) {
	ts, err := tokenize("2+3*cos(3)")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ts)
}

// ================ Benchmarks ================

var result float64 //https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go compiler optimization sections

func BenchmarkExpressionEvaluate(b *testing.B) {
	//Parse the expression
	expr := "2+2*5+3*5+2*5+3*72"
	e, _ := ParseExpression(expr)
	var r float64
	// evaluate the expression N times
	for n := 0; n < b.N; n++ {
		r = e.Evaluate(map[string]float64{})
	}
	result = r
}

func BenchmarkExpressionCompiled(b *testing.B) {
	//Parse the expression
	expr := "2+2*5+3*5+2*5+3*72"
	e, _ := ParseExpression(expr)
	f := CompileExpression(e)

	var r float64
	// evaluate the expression N times
	for n := 0; n < b.N; n++ {
		r = f(map[string]float64{})
	}
	result = r
}
func BenchmarkExpressionNative(b *testing.B) {
	//Parse the expression
	//expr := "2+2*5+3*5+2*5+3*72"
	e := func(a, b, c, d float64) float64 { return a + a*c + c*c + a*c + b*d }
	var r float64
	// evaluate the expression N times
	for n := 0; n < b.N; n++ {
		r = e(2, 3, 5, 72)
	}
	result = r
}

func BenchmarkExpressionJIT(b *testing.B) {
	////Parse the expression
	//expr := "2+2+3+4"
	//e, _ := ParseExpression(expr)
	//f := JitCompileExpression(e)
	//var r float64
	//// evaluate the expression N times
	//for n := 0; n < b.N; n++ {
	//	r = f(map[string]float64{})
	//}
	//result = r
}

//Segmentation and or nil pointer dereference
//something to do with way executablePrintFunc is made?
//Maybe with way it is copied over
//Maybe the way it is compiled but i doubt that a bit

func TestExpressionJIT(t *testing.T) {
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
	data := []float64{1, 2, 3, -8008, -8008}

	f := MakeMathFunc(mathFunction2)

	res := f(data)
	fmt.Println(res)
	//Parse the expression
	//expr := "2+2+3+4"
	//e, err := ParseExpression(expr)
	//fmt.Println(e.String())
	//if err != nil {
	//	t.Errorf(err.Error())
	//	return
	//}
	//f := JitCompileExpression(e)
	//var r = f(map[string]float64{"a": 1})
	//if r != float64(2+2+3+4) {
	//	t.Errorf("Had 2+2+3+4. Wanted %g, got %g", float64(2+2+3+4), r)
	//}
}
