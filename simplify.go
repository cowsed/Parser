package parser

import "fmt"

//Do something similar with add and subtract to what is done with multiplication
//get list of a addends and subtractants and simplify that list down by combining like terms, then turn it back into tree
func (a Adder) Simplify() Expression {
	aIs0 := false
	bIs0 := false
	aIsConst := false
	bIsConst := false

	aVal := 0.0
	bVal := 0.0
	A := a.A.Simplify()
	B := a.B.Simplify()
	switch v := A.(type) {
	case Constant:
		aIsConst = true
		aVal = v.Value
		if v.Value == 0 {
			aIs0 = true
		}
	}
	switch v := B.(type) {
	case Constant:
		bIsConst = true
		bVal = v.Value
		if v.Value == 0 {
			bIs0 = true
		}
	}
	if aIsConst && bIsConst {
		return Constant{aVal + bVal}
	}
	//Identity 0+0=0, 0+x=x
	if aIs0 && bIs0 {
		return Constant{0}
	} else if aIs0 {
		return B
	} else if bIs0 {
		return A
	}

	return a
}
func (s Subtractor) Simplify() Expression {
	aIs0 := false
	bIs0 := false
	A := s.A.Simplify()
	B := s.B.Simplify()
	switch v := A.(type) {
	case Constant:
		if v.Value == 0 {
			aIs0 = true
		}
	}
	switch v := B.(type) {
	case Constant:
		if v.Value == 0 {
			bIs0 = true
		}
	}
	if aIs0 && bIs0 {
		return Constant{0}
	} else if aIs0 {
		return Multiplier{
			A: Constant{-1},
			B: B,
		}
	} else if bIs0 {
		return A
	}
	return s
}

func (m Multiplier) Simplify() Expression {
	aIs0 := false
	bIs0 := false
	aIs1 := false
	bIs1 := false
	aIsConst := false
	bIsConst := false
	aVal := 0.0
	bVal := 0.0
	//aSymbol := ""
	//bSymbol := ""

	fmt.Println("Simplifying multiplication{", m.String())
	A := m.A.Simplify()
	B := m.B.Simplify()
	fmt.Println("Simplifying multiplication2}", Multiplier{A, B})

	fmt.Printf("A: %v\n", A)
	fmt.Printf("B: %v\n", B)
	//Get data to check for identity rules (1*x=x, 0*x=0)
	switch v := A.(type) {
	case Constant:
		aIsConst = true
		aVal = v.Value
		aIs0 = v.Value == 0
		aIs1 = v.Value == 1
	case Variable:
		//aSymbol = v.Symbol
	}
	switch v := B.(type) {
	case Constant:
		bIsConst = true
		bVal = v.Value
		bIs0 = v.Value == 0
		bIs1 = v.Value == 1
	case Variable:
		//bSymbol = v.Symbol
	}

	//Check Identity rules
	if aIs0 || bIs0 {
		fmt.Println("ZEROS", m.String())
		return Constant{0}
	} else if aIs1 {
		return B
	} else if bIs1 {
		return A
	}
	//Other possibillities
	if aIsConst && bIsConst {
		return Constant{aVal * bVal}
	} else if aIsConst {
		return Multiplier{
			Constant{aVal},
			B,
		}

	} else if bIsConst {
		return Multiplier{
			A: A,
			B: Constant{bVal},
		}
	}

	////Simplify x*x to x^2
	//if aSymbol == bSymbol && aSymbol != "" {
	//	return Powerer{
	//		Base:     Variable{aSymbol},
	//		Exponent: Constant{2},
	//	}
	//}
	//while If Mul*Div or Mul*mul or Mul*div
	//list of numerators, list of denominators
	//Variable and degree, if its just Variable{x} turn it into x^1 so later can subtract

	//CurrentElement := m
	//Numerator, Denominator := ListMuls(CurrentElement, true)
	//fmt.Printf("Numerator: %v\n", Numerator)
	//fmt.Printf("Denominator: %v\n", Denominator)
	////Turn them into powers
	//simplified := SimplifyFraction(Numerator, Denominator)
	//
	//return simplified
	//for {
	//	switch type of Current element a,b
	//	if is mul or div, follow that branch
	//	if not, add to numerator or denominator, break out of that branch
	//
	//}

	fmt.Println("Never Simplified")
	return Multiplier{
		A: A,
		B: B,
	}
}

//Fix Division
//if !Num2Num, in divisions, things that say they go in num go to denom and vice versa
//Call next level with things reversed
//tmrw implement outer case divider such that it puts a in num and b in denom and also sorts it out with its children
func ListMuls(e Expression, NumToNum bool) ([]Expression, []Expression) {
	num := []Expression{}
	denom := []Expression{}
	switch v := e.(type) {
	case Multiplier:
		fmt.Println("Multiplication count muls")
		n2, d2 := ListMuls(v.A, NumToNum)
		n3, d3 := ListMuls(v.B, NumToNum)
		if NumToNum {
			num = append(num, n2...)
			denom = append(denom, d2...)
			num = append(num, n3...)
			denom = append(denom, d3...)

		} else {
			num = append(num, d2...)
			denom = append(denom, n2...)
			num = append(num, d3...)
			denom = append(denom, n3...)

		}

	case Divider:
		fmt.Println("DIVIDERERER")
		n2, d2 := ListMuls(v.A, NumToNum)

		n3, d3 := ListMuls(v.B, NumToNum)
		if NumToNum {
			num = append(num, n2...)
			denom = append(denom, d2...)
			num = append(num, d3...)
			denom = append(denom, n3...)
		} else {
			num = append(num, d2...)
			denom = append(denom, n2...)
			num = append(num, n3...)
			denom = append(denom, d3...)

		}
		fmt.Println("num============================================", num)
		fmt.Println("denom============================================", denom)

	default:
		num = append(num, v)
	}

	return num, denom
}

func (d Divider) Simplify() Expression {
	bIs1 := false
	aIs0 := false
	aVal := 0.0
	bVal := 0.0
	aIsConst := false
	bIsConst := false
	A := d.A.Simplify()
	B := d.B.Simplify()
	fmt.Println("Simplifying devision of ", A, "/", B)
	switch v := A.(type) {
	case Constant:
		aIsConst = true
		aVal = v.Value
		aIs0 = v.Value == 0
	}
	switch v := B.(type) {
	case Constant:
		bIsConst = true
		bVal = v.Value
		bIs1 = v.Value == 0
	}
	//Identities
	if bIs1 {
		return A
	} else if aIs0 {
		return Constant{0}
	} else if aIsConst && bIsConst {
		//const over const simplifies to const
		return Constant{aVal / bVal}
	}
	////Get list of numerator and denominator
	//Numerator, Denominator := ListMuls(d, true)
	////Numerator2, Denominator2 := ListMuls(B, false)
	////Numerator = append(Numerator, Numerator2...)
	////Denominator = append(Denominator, Denominator2...)
	//fmt.Println("Divider num", Numerator)
	//fmt.Println("Divider denom", Denominator)
	//fmt.Printf("For dividing num: %v, denom: %v\n", Numerator, Denominator)
	//
	//simplified := SimplifyFraction(Numerator, Denominator)
	//return simplified
	////Didn't simplify
	return Divider{
		A: A,
		B: B,
	}
}
func SimplifyFraction(Numerator, Denominator []Expression) Expression {
	degreeCounts := map[Expression][]Expression{}
	coefficientsInNum := []float64{}
	coefficientsInDenom := []float64{}

	//simplifiedDenom := []Expression{}
	fmt.Printf("Numerator: %v\n", Numerator)

	for _, e := range Numerator {
		switch v := e.(type) {
		case Powerer:
			//if is var to power, add power to degreecounts,
			//if is something else to power, use its string representation as key and add degree

			degreeCounts[v.Base] = append(degreeCounts[v.Base], v.Exponent)
		case Variable:
			degreeCounts[v] = append(degreeCounts[v], Constant{1})
		case Constant:
			coefficientsInNum = append(coefficientsInNum, v.Value)
		default:
			degreeCounts[v] = append(degreeCounts[v], Constant{1})
			fmt.Println("Counting degree of ", v.String())

		}
	}
	for _, e := range Denominator {
		switch v := e.(type) {
		case Powerer:
			//if is var to power, subtract power to degreecounts,
			degreeCounts[v] = append(degreeCounts[v], Multiplier{A: Constant{-1}, B: v.Exponent})
		case Variable:
			fmt.Println("Variavle in denom")
			degreeCounts[v] = append(degreeCounts[v], Constant{-1})
		case Constant:
			coefficientsInDenom = append(coefficientsInDenom, v.Value)
		default:
			//simplifiedNum = append(simplifiedNum, v)
			degreeCounts[v] = append(degreeCounts[v], Constant{-1})
			fmt.Println("Counting degree of ", v.String())

		}
	}

	//Find the product of all the coeffecients
	CoeffecientProduct := 1.0
	for i := range coefficientsInNum {
		CoeffecientProduct *= coefficientsInNum[i]
	}
	for i := range coefficientsInDenom {
		CoeffecientProduct /= coefficientsInDenom[i]
	}

	fmt.Printf("degreeCounts: %v\n", degreeCounts)
	fmt.Printf("prod: %v\n", CoeffecientProduct)

	parts := []Expression{}
	for k, v := range degreeCounts {
		if len(v) == 1 {
			parts = append(parts, k.Simplify())
			continue
		} else if len(v) == 2 {
			parts = append(parts, Powerer{Base: k, Exponent: Adder{
				A: v[0],
				B: v[1],
			}}.Simplify())
		} else {
			degree := Adder{
				A: v[0],
				B: v[1],
			}
			for i := 2; i < len(v); i++ {
				degree = Adder{
					A: degree,
					B: v[i],
				}
			}

			parts = append(parts, Powerer{
				Base:     k,
				Exponent: degree,
			}.Simplify())
		}
	}
	fmt.Printf("parts: %v\n", parts)
	if len(parts) == 0 {
		return Constant{CoeffecientProduct}
	} else if len(parts) == 1 {
		return Multiplier{
			A: Constant{CoeffecientProduct},
			B: parts[0],
		}.Simplify()
	} else {
		prevProduct := Multiplier{
			A: parts[0],
			B: parts[1],
		}
		for i := 2; i < len(parts); i++ {
			prevProduct = Multiplier{
				A: prevProduct,
				B: parts[i],
			}
		}
		return Multiplier{
			A: Constant{CoeffecientProduct},
			B: prevProduct,
		} //.Simplify()
	}
}

func (p Powerer) Simplify() Expression {
	One := Constant{1}
	Zero := Constant{0}

	if p.Exponent.Simplify() == One {
		return p.Base
	} else if p.Exponent.Simplify() == Zero {
		return Constant{1}
	}
	return Powerer{
		Base:     p.Base.Simplify(),
		Exponent: p.Exponent.Simplify(),
	}
}
func (n NaturalLogger) Simplify() Expression {
	return NaturalLogger{n.A.Simplify()}
}
func (c Constant) Simplify() Expression {
	return c
}
func (v Variable) Simplify() Expression {
	return v
}

func (c Coser) Simplify() Expression {
	return Coser{c.A.Simplify()}
}

func (s Siner) Simplify() Expression {
	return Siner{s.A.Simplify()}
}
