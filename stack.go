package parser

import "fmt"

//TokenStack is FIFO stack of tokens for parsing
type TokenStack struct {
	elements []Token
	top      int
}

//NewTokenStack returns an empty TokenStack
func NewTokenStack() TokenStack {
	return TokenStack{
		elements: []Token{},
		top:      0,
	}

}

//Len Returns the length of the stack
func (s *TokenStack) Len() int {
	return len(s.elements)
}

//Push pushes a token to the stack
func (s *TokenStack) Push(n Token) {
	s.elements = append(s.elements, n)
	s.top++
}

//Pop takes a token from the stack
func (s *TokenStack) Pop() (Token, error) {
	if s.top <= 0 {
		return Token{}, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	s.elements = s.elements[0 : s.top-1]
	s.top--
	return n, nil
}

//Peek looks at the top of the stack without removing it
func (s *TokenStack) Peek() (Token, error) {
	if s.top <= 0 {
		return Token{}, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	return n, nil
}

//ExpressionStack is a FIFO stack of expression elements
type ExpressionStack struct {
	elements []Expression
	top      int
}

//NewExpressionStack returns an empty expression stack
func NewExpressionStack() ExpressionStack {
	return ExpressionStack{
		elements: []Expression{},
		top:      0,
	}

}

//Len Returns the length of the stack
func (s *ExpressionStack) Len() int {
	return len(s.elements)
}

//Push pushes a token to the stack
func (s *ExpressionStack) Push(n Expression) {
	s.elements = append(s.elements, n)
	s.top++
}

//Pop takes a token from the stack
func (s *ExpressionStack) Pop() (Expression, error) {
	if s.top <= 0 {
		return nil, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	s.elements = s.elements[0 : s.top-1]
	s.top--
	return n, nil
}

//Peek looks at the top of the stack without removing it
func (s *ExpressionStack) Peek() (Expression, error) {
	if s.top <= 0 {
		return nil, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	return n, nil
}
