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

type QnA struct {
	q string
	a float64
}

func TestEvaluateN(t *testing.T) {
	tests := []QnA{
		{"2+2", 4},
		{"2*3", 6},
		{"3-2", 1},
		{"3*2+cos(0)", 7},
		{"3+cos(0)", 4},
		{"3+sin(0)", 3},

		{"3/2", 1.5},
		{"3^2", 9},
		{"(9-x^2)^.5", 5},
	}
	for i := range tests {
		e, err := ParseExpression(tests[i].q)
		if err != nil {
			t.Error(err)
		}
		if e.Evaluate(map[string]float64{"x": -2}) != tests[i].a {
			t.Errorf("%s should = %g but evaluated to %g. Parsed to %s", tests[i].q, tests[i].a, e.Evaluate(map[string]float64{"x": -2}), e.Simplify().String())
		}
	}
}

type IntegrateQnA struct {
	exp        string
	wrt        string
	from, to   float64
	expected   float64
	numBuckets int
}

func TestIntegrateN(t *testing.T) {
	epsilon := 0.001
	tests := []IntegrateQnA{
		{"2", "x", 0, 4, 8, 1},
		{"x", "x", 0, 4, 8, 1},
		{"x^2", "x", 0, 4, 64.0 / 3.0, 10000},
		{"sin(x)", "x", 0, 4, 1.6536, 100},
		//https://tutorial.math.lamar.edu/classes/calcii/surfacearea.aspx
		{"2*3.14159*((9-x^2)^.5)*(3/((9-x^2)^.5)", "x", -2, 2, 24 * math.Pi, 100},
	}
	//*
	for i := range tests {
		e, err := ParseExpression(tests[i].exp)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("Finished parsing", e.String())

		test := tests[i]
		res := IntegrateV(e, map[string]float64{}, test.wrt, test.from, test.to, test.numBuckets)
		if math.Abs(res-test.expected) > epsilon || math.IsNaN(res) {
			fmt.Println(res-test.expected, epsilon)
			t.Errorf("%s from %g to %g wrt %s should = %g but was %g. Parsed to %s", test.exp, test.from, test.to, test.wrt, test.expected, res, e.String())
		}
	}
}

func TestParseToPostfix(t *testing.T) {
	expr := "3*4+2"
	tokens, err := tokenize(expr)
	if err != nil {
		t.Errorf(err.Error())
	}
	pf, err := makePostFix(tokens)
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
	pf, err := makePostFix(tokens)
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

func TestSimplify1(t *testing.T) {
	e, err := ParseExpression("0*ln(x)")
	if err != nil {
		t.Error(err)
	}
	if e.Simplify().String() != "0" {
		t.Errorf("Wanted %s, got %s", "0", e.Simplify().String())
	}
}
func TestSimplify2(t *testing.T) {
	e, err := ParseExpression("2+0*ln(x)")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("e: %v\n", e)
	if e.Simplify().String() != "2" {
		t.Errorf("Wanted %s, got %s", "2", e.Simplify().String())
	}
}
func TestSimplify3(t *testing.T) {

	e, err := ParseExpression("2*x*x")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("e: %v\n", e)
	fmt.Printf("e.Simplify(): %v\n", e.Simplify())
	s := e.Simplify()
	if e.Simplify().String() != "(2 * (x ^ 2))" {
		t.Errorf("Wanted %s, got %s", "(2 * (x ^ 2))", e.Simplify().String())
	}
	fmt.Println("WOWW", s)
}

func TestSimplify4(t *testing.T) {

	e, err := ParseExpression("2*x*b*b*x")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("e: %v\n", e)
	if e.Simplify().String() != "(2 * ((x ^ 2) * (b ^ 2)))" {
		t.Errorf("Wanted %s, got %s", " (2 * ((x ^ 2) * (b ^ 2)))", e.Simplify().String())
	}
}
func TestSimplify5(t *testing.T) {

	e, err := ParseExpression("(x+1)*(x+1)")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("e: %v\n", e)
	if e.Simplify().Simplify().String() != "((x + 1) ^ 2)" {
		t.Errorf("Wanted %s, got %s", "((x + 1) ^ 2))", e.Simplify().Simplify().String())
	}
}

func TestSimplifyN(t *testing.T) {
	tests := [][2]string{
		{"2+2", "4"},
		{"(x+1) * (x+1)", "((x + 1) ^ 2)"},
		{"x*x*x*x", "(x ^ 4)"},
		{"2/2", "1"},
		{"0/x", "0"},
		{"0/2", "0"},
		{"x*0*2", "0"},
		{"2*(2/1)", "4"},
		{"(2*2)/1", "4"},
		{"(2*2)/(3*3)", fmt.Sprint(4.0 / 9.0)},

		{"3^(x+2+3)", "(3 ^ (x + 5))"},

		{"x/x", "1"},
		{"x/x*x", "x"},
		{"x*x*x/x", "(x ^ 2)"},
		{"x/x*x/x", "1"},
	}
	//{"x/x", "1"},

	for i := range tests {
		e, err := ParseExpression(tests[i][0])
		if err != nil {
			t.Error(err)
		}
		//fmt.Printf("%s Simplifies to %s\n", tests[i][0], tests[i][1])
		if e.Simplify().String() != tests[i][1] {
			t.Errorf("%s should simplify to %s but simplified to %s", tests[i][0], tests[i][1], e.Simplify().String())
		}
	}
}

func TestCountMuls(t *testing.T) {
	e, _ := ParseExpression("x+x/x")
	fmt.Println("Expression,", e)
	n, d := ListMuls(e, true)
	fmt.Println(n, d)
	fmt.Printf("SimplifyFraction(n, d): %v\n", SimplifyFraction(n, d))
}

func TestDerivative1(t *testing.T) {
	e, err := ParseExpression("x^2")
	if err != nil {
		t.Errorf(err.Error())
	}
	derivative := e.Derive("x")
	//derivative = derivative.Simplify()
	res := derivative.Evaluate(map[string]float64{"x": 4})
	fmt.Println(res)
	fmt.Println("Derivative", derivative.String())
	fmt.Printf("Wanted %s=%g. got %s = %g", "2 * x", 8.0, derivative.String(), res)

	if res != 8 {
		t.Errorf("Wanted %s=%g. got %s = %g", "2 * x", 8.0, derivative.String(), res)
	}
}

func TestDerivative2(t *testing.T) {
	e, err := ParseExpression("x^3")
	if err != nil {
		t.Errorf(err.Error())
	}
	derivative := e.Derive("x")
	res := derivative.Evaluate(map[string]float64{"x": 4})
	fmt.Println(res)
	fmt.Println(derivative.String())
	if res != 16*3 {
		t.Errorf("Wanted %s=%g. got %s=%g", "3 * x ^ 2", 48.0, derivative.String(), res)
	}
}

type DerivQnA struct {
	e            string
	testNum, ans float64
}

func TestDerivativeN(t *testing.T) {
	tests := []DerivQnA{{
		e:       "2*x^2",
		testNum: 1,
		ans:     4,
	},
		{
			e:       "9-x",
			testNum: 1,
			ans:     -1,
		},
		{
			e:       "(9-x)^2",
			testNum: 9,
			ans:     0,
		},
	}
	for i := range tests {
		e, err := ParseExpression(tests[i].e)
		if err != nil {
			t.Error(err)
		}
		test := tests[i]
		d := e.Derive("x")
		res := d.Evaluate(map[string]float64{"x": tests[i].testNum})
		if res != tests[i].ans {
			t.Errorf("Expected (%g,%g). got (%g,%g). Derivative is %s", test.testNum, test.ans, test.testNum, res, d.String())
		}
	}
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
