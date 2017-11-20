package main

import (
	"fmt"
	"strconv"
)

//Given a string containing only digits, restore it by returning all possible valid IP address combinations.
//
//For example:
//Given "25525511135",
//
//return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)

func restoreIpAddresses(s string) []string {
	if len(s) < 3 || len(s) > 12 {
		return nil
	}

	var res []string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			for k := j + 1; k < len(s)-1; k++ {
				p1 := s[0 : i+1]
				p2 := s[i+1 : j+1]
				p3 := s[j+1 : k+1]
				p4 := s[k+1 : len(s)]
				if isValid(p1) && isValid(p2) && isValid(p3) && isValid(p4) {
					res = append(res, p1+"."+p2+"."+p3+"."+p4)
				}
			}
		}
	}
	return res
}

func isValid(s string) bool {
	n, _ := strconv.Atoi(s)
	if len(s) != len(strconv.Itoa(n)) || n > 255 {
		return false
	}
	return true
}

func main() {
	var s = "25525511135"
	res := restoreIpAddresses(s)
	for _, v := range res {
		fmt.Println(v)
	}
}
