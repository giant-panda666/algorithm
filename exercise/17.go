package main

import "fmt"

//Given a digit string, return all possible letter combinations that the number could represent.
//
//A mapping of digit to letters (just like on the telephone buttons) is given below.
// 0 " "; 1 ""; 2 "abc"; 3 "def"; 4 "ghi"; 5 "jkl"; 6 "mno"; 7 "pqrs" 8 "tuv"; 9 "wxyz"
//
//
//Input:Digit string "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
func findCombinations(digits string, index int, s string, res *[]string) {
	if len(digits) == index {
		*res = append(*res, s)
		return
	}

	str := lettersMap[int([]rune(digits)[index]-'0')]
	for _, v := range str {
		tmp := s + string(v)
		findCombinations(digits, index+1, tmp, res)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var res []string
	var s string
	findCombinations(digits, 0, s, &res)
	return res
}

var lettersMap = map[int]string{
	0: " ",
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func main() {
	fmt.Println(letterCombinations("2"))
}
