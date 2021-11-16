package parser

import "fmt"

//FIFO stack
type TokenStack struct {
	elements []Token
	top      int
}

func NewTokenStack() TokenStack {
	return TokenStack{
		elements: []Token{},
		top:      0,
	}

}

func (s *TokenStack) Len() int {
	return len(s.elements)
}
func (s *TokenStack) Push(n Token) {
	s.elements = append(s.elements, n)
	s.top++
}

func (s *TokenStack) Pop() (Token, error) {
	if s.top <= 0 {
		return Token{}, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	s.elements = s.elements[0 : s.top-1]
	s.top--
	return n, nil
}

func (s *TokenStack) Peek() (Token, error) {
	if s.top <= 0 {
		return Token{}, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	return n, nil
}

type ExpressionStack struct {
	elements []Expression
	top      int
}

func NewExpressionStack() ExpressionStack {
	return ExpressionStack{
		elements: []Expression{},
		top:      0,
	}

}

func (s *ExpressionStack) Len() int {
	return len(s.elements)
}
func (s *ExpressionStack) Push(n Expression) {
	s.elements = append(s.elements, n)
	s.top++
}

func (s *ExpressionStack) Pop() (Expression, error) {
	if s.top <= 0 {
		return nil, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	s.elements = s.elements[0 : s.top-1]
	s.top--
	return n, nil
}

func (s *ExpressionStack) Peek() (Expression, error) {
	if s.top <= 0 {
		return nil, fmt.Errorf("0 Length Stack")
	}
	n := s.elements[(s.top - 1)]
	return n, nil
}
