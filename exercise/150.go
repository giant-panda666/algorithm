package main

import (
	"strconv"
)

//Evaluate the value of an arithmetic expression in Reverse Polish Notation.
//
//Valid operators are +, -, *, /. Each operand may be an integer or another expression.
//
//Some examples:
//  ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
//  ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6

type stack struct {
	data  []string
	count int
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
	s.count++
}

func (s *stack) pop() string {
	if s.count == 0 {
		return ""
	}
	s.count--
	res := s.data[s.count]
	s.data = s.data[:s.count]
	return res
}

func evalRPN(tokens []string) int {
	var s stack
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			str1 := s.pop()
			str2 := s.pop()
			num1, _ := strconv.Atoi(str1)
			num2, _ := strconv.Atoi(str2)
			if v == "+" {
				s.push(strconv.Itoa(num2 + num1))
			} else if v == "-" {
				s.push(strconv.Itoa(num1 - num2))
			} else if v == "*" {
				s.push(strconv.Itoa(num2 * num1))
			} else if v == "/" {
				s.push(strconv.Itoa(num1 / num2))
			}
		default:
			s.push(v)
		}
	}

	str := s.pop()
	num, _ := strconv.Atoi(str)
	return num
}
