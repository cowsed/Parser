package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type ElementType int

const (
	OperatorType ElementType = iota
	NumberType
	VariableType
	LeftParenType
	RightParenType
	FunctionType
)

type Token struct {
	Type  ElementType
	Value string
}

func ParseExpression(expr string) (Expression, error) {
	tokens, err := tokenize(expr)

	if err != nil {
		return nil, err
	}

	postfix, err := MakePostFix(tokens)
	if err != nil {
		return nil, err
	}

	return ParsePostfix(postfix)
}

func ParsePostfix(tokens []Token) (Expression, error) {
	var PartsStack = NewExpressionStack()
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		switch t.Type {
		case NumberType:
			v, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				return nil, err
			}
			PartsStack.Push(Constant{
				Value: v,
			})
		case OperatorType:
			switch t.Value {
			case "+":
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				B, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Adder{
					A: B,
					B: A,
				})
			case "-":
				B, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Subtractor{
					A: A,
					B: B,
				})
			case "*":
				B, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Multiplier{
					A: A,
					B: B,
				})
			case "/":
				B, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Divider{
					A: A,
					B: B,
				})
			case "^":
				B, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Powerer{
					Base:     A,
					Exponent: B,
				})
			}
		case VariableType:
			symbol := t.Value
			PartsStack.Push(Variable{
				Symbol: symbol,
			})
		case FunctionType:
			switch t.Value {
			case "cos":
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Coser{A})
			case "sin":
				A, err := PartsStack.Pop()
				if err != nil {
					return nil, err
				}
				PartsStack.Push(Siner{A})
			}
		}
	}
	return PartsStack.Pop()
}

func MakePostFix(tokens []Token) ([]Token, error) {
	precedence := map[string]int{
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"^": 4,
	}
	output := []Token{}

	operatorStack := NewTokenStack()
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == NumberType || tokens[i].Type == VariableType {

			output = append(output, tokens[i])
		} else if tokens[i].Type == FunctionType {
			operatorStack.Push(tokens[i])
		} else if tokens[i].Type == OperatorType {

			o1 := tokens[i]
			for {
				// there is an operator o2 other than the left parenthesis at the top
				// of the operator stack, and (o2 has greater precedence than o1
				// or they have the same precedence and o1 is left-associative)
				o2, err := operatorStack.Peek()
				if err != nil { //there is an operator o2
					break
				}
				if o2.Value == "(" { //o2 other than the left parenthesis at the top of the operator stack
					break
				}
				if !(precedence[o2.Value] > precedence[o1.Value] || (precedence[o2.Value] == precedence[o1.Value] && o1.Value != "^")) { // (o2 has greater precedence than o1 or they have the same precedence and o1 is left-associative)
					break
				}
				o2_2, _ := operatorStack.Pop()
				output = append(output, o2_2)
			}
			operatorStack.Push(o1)
		} else if tokens[i].Type == LeftParenType {
			operatorStack.Push(tokens[i])
		} else if tokens[i].Type == RightParenType {
			for {
				o, err := operatorStack.Peek()
				if err != nil {
					return nil, fmt.Errorf("unbalanced parentheses at index %d '%s'. next token: %s", i, tokens[i].Value, o.Value)
				}
				if o.Type != LeftParenType {
					o_2, _ := operatorStack.Pop()
					output = append(output, o_2)
				} else {
					break
				}
			}
			//{assert there is a left parenthesis at the top of the operator stack}
			o_1, err := operatorStack.Peek()
			if err != nil || o_1.Value != "(" {
				return nil, fmt.Errorf("should be a left parenthesis here")
			}
			operatorStack.Pop()
			//if there is a function token at the top of the operator stack, then:
			//pop the function from the operator stack into the output queue
			o, err := operatorStack.Peek()
			if err == nil {
				if o.Type == FunctionType {
					o, _ = operatorStack.Pop()
					output = append(output, o)
				}
			}

		}
	}
	for operatorStack.Len() > 0 {
		o, _ := operatorStack.Pop()
		output = append(output, o)
	}
	return output, nil
}

func tokenize(s string) ([]Token, error) {
	varParts := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	functions := []string{"cos", "sin"}
	numberParts := "1234567890."
	operators := "+-*/^"
	s = strings.ReplaceAll(s, " ", "")
	tokens := []Token{}

	var MidNumber = false
	var MidVar = false
	var currentTokenVal = ""
	for i := 0; i < len(s); i++ {
		//Error checking
		if MidNumber && MidVar {
			return nil, fmt.Errorf("error at index %d, char: %s; a number and a variable together", i, string(s[i]))
		}

		//Character to analyze
		var char string = string(s[i])

		//Is part of a number
		IsNumberPart := strings.ContainsAny(char, numberParts)
		if IsNumberPart {
			if MidVar {
				//Finish variable
				var typeOf = VariableType
				if MatchesAny(currentTokenVal, functions) {
					typeOf = FunctionType
				}
				tokens = append(tokens, Token{
					Type:  typeOf,
					Value: currentTokenVal,
				})
				currentTokenVal = ""
				MidVar = false
			}
			currentTokenVal += char
			MidNumber = true
			continue
		}
		//End Number
		if MidNumber && !IsNumberPart {
			//End of the previous number
			tokens = append(tokens, Token{
				Type:  NumberType,
				Value: currentTokenVal,
			})
			currentTokenVal = ""
			MidNumber = false
		}

		//Is part of a variable
		IsVarPart := strings.ContainsAny(char, varParts)
		if IsVarPart {
			currentTokenVal += char
			MidVar = true
			//Check if var so far matches any functions, if it does MidVar = false token = whatever that funciton is

			continue
		}
		//End Variable
		if MidVar && !IsVarPart {
			//End of variable name
			var typeOf = VariableType
			if MatchesAny(currentTokenVal, functions) {
				typeOf = FunctionType
			}
			tokens = append(tokens, Token{
				Type:  typeOf,
				Value: currentTokenVal,
			})
			currentTokenVal = ""
			MidVar = false
		}

		IsOperator := strings.ContainsAny(char, operators)
		if IsOperator {

			currentTokenVal = char
			tokens = append(tokens, Token{
				Type:  OperatorType,
				Value: currentTokenVal,
			})
			currentTokenVal = ""
			continue
		}

		if char == "(" {
			tokens = append(tokens, Token{
				Type:  LeftParenType,
				Value: "(",
			})
			continue
		}
		if char == ")" {
			tokens = append(tokens, Token{
				Type:  RightParenType,
				Value: ")",
			})
			continue
		}

	}
	//Check if it was the end and partway through a number
	if MidNumber {
		tokens = append(tokens, Token{
			Type:  NumberType,
			Value: currentTokenVal,
		})
		currentTokenVal = ""
	}
	if MidVar {
		//End of variable name
		tokens = append(tokens, Token{
			Type:  VariableType,
			Value: currentTokenVal,
		})
		currentTokenVal = ""
	}
	return tokens, nil
}

func MatchesAny(s string, substrs []string) bool {
	for i := 0; i < len(substrs); i++ {
		if s == substrs[i] {
			return true
		}
	}
	return false
}
