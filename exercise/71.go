package main

import (
	"fmt"
	"strings"
)

//Given an absolute path for a file (Unix-style), simplify it.
//
//For example,
//path = "/home/", => "/home"
//path = "/a/./b/../../c/", => "/c"

type stack struct {
	data  []string
	count int
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
	s.count++
}

func (s *stack) isEmpty() bool {
	return s.count == 0
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

func simplifyPath(path string) string {
	var strs = strings.Split(path, "/")
	var s stack
	for _, v := range strs {
		switch v {
		case "", ".":
		case "..":
			fmt.Println(v)
			s.pop()
		default:
			fmt.Println(v)
			s.push(v)
		}
	}

	if s.isEmpty() {
		return "/"
	}

	var res string
	for !s.isEmpty() {
		tmp := s.pop()
		res = "/" + tmp + res
	}
	return res
}

func main() {
	var str = "/abc/..."
	fmt.Println(simplifyPath(str))
}
