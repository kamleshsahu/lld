package main

import (
	"container/list"
	"strings"
)

type Tokenizer struct {
	tokens []rune
	pos    int
}

func NewTokenizer(s string) *Tokenizer {
	return &Tokenizer{tokens: []rune(strings.ReplaceAll(s, " ", "")), pos: 0}
}

func (t *Tokenizer) hasNext() bool {
	return t.pos < len(t.tokens)
}

func (t *Tokenizer) current() rune {
	return t.tokens[t.pos]
}

func (t *Tokenizer) next() {
	if t.hasNext() {
		t.pos++
	}
}

type Solution struct {
	tokens *Tokenizer
}

func NewSolution() *Solution {
	return &Solution{}
}

func (s *Solution) calculateExpression() int {
	num := 0
	sign := '+'
	stk := list.New()

	operatorMap := map[rune]func(*list.List, int){
		'+': func(stk *list.List, v int) {
			stk.PushBack(v)
		},
		'-': func(stk *list.List, v int) {
			stk.PushBack(-v)
		},
		'*': func(stk *list.List, v int) {
			top := stk.Remove(stk.Back()).(int)
			stk.PushBack(top * v)
		},
		'/': func(stk *list.List, v int) {
			top := stk.Remove(stk.Back()).(int)
			stk.PushBack(top / v)
		},
	}

	for s.tokens.hasNext() {
		c := s.tokens.current()
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if _, ok := operatorMap[c]; ok {
			operatorMap[rune(sign)](stk, num)
			sign = c
			num = 0
		} else if c == '(' {
			s.tokens.next()
			num = s.calculateExpression()
		} else if c == ')' {
			operatorMap[rune(sign)](stk, num)
			return sumStack(stk)
		}
		s.tokens.next()
	}

	operatorMap[rune(sign)](stk, num)
	return sumStack(stk)
}

func (s *Solution) Calculate(expression string) int {
	s.tokens = NewTokenizer(expression)
	return s.calculateExpression()
}

func sumStack(stk *list.List) int {
	sum := 0
	for e := stk.Front(); e != nil; e = e.Next() {
		sum += e.Value.(int)
	}
	return sum
}

func main() {
	expression := "3+(2*2)-4/2"
	sol := NewSolution()
	result := sol.Calculate(expression)
	println("Result:", result)
}
