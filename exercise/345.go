package main

import "fmt"

//Write a function that takes a string as input and reverse only the vowels of a string.
//
//Example 1:
//Given s = "hello", return "holle".
//
//Example 2:
//Given s = "leetcode", return "leotcede".
//
//Note:
//The vowels does not include the letter "y".

var vowels = []rune{'a', 'e', 'i', 'o', 'u',
	'A', 'E', 'I', 'O', 'U'}

func isVowel(c rune) bool {
	for _, v := range vowels {
		if v == c {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	tmp := []rune(s)
	i, j := 0, len(tmp)-1
	for j > i {
		for i < j && !isVowel(tmp[i]) {
			i++
		}
		for i < j && !isVowel(tmp[j]) {
			j--
		}
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}

	return string(tmp)
}

func main() {
	var s1 = "hello"
	fmt.Println("reverseVowels(hello), expect holle, result is", reverseVowels(s1))
	var s2 = "leetcode"
	fmt.Println("reverseVowels(leetcode), expect leotcede, result is", reverseVowels(s2))
}
