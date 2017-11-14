package main

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

type stack struct {
	data  []rune
	count int
}

func (s *stack) push(r rune) {
	s.data = append(s.data, r)
	s.count++
}

func (s *stack) pop() rune {
	if s.count == 0 {
		return ' '
	}
	s.count--
	res := s.data[s.count]
	s.data = s.data[:s.count]
	return res
}

func (s *stack) isEmpty() bool {
	return s.count == 0
}

func isValid(s string) bool {
	runes := []rune(s)

	var ss stack
	for _, v := range runes {
		switch v {
		case '}':
			if !isParentheses(&ss, '{') {
				return false
			}
		case ')':
			if !isParentheses(&ss, '(') {
				return false
			}
		case ']':
			if !isParentheses(&ss, '[') {
				return false
			}
		case '{', '(', '[':
			ss.push(v)
		}
	}
	if !ss.isEmpty() {
		return false
	}
	return true
}

func isParentheses(s *stack, r rune) bool {
	tmp := s.pop()
	if tmp != r {
		return false
	}
	return true
}
