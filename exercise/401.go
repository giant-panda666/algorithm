package main

import "fmt"

//A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).
//
//Each LED represents a zero or one, with the least significant bit on the right.
//
//
//For example, the above binary watch reads "3:25".
//
//Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.
//
//Example:
//
//Input: n = 1
//Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

var hourMap = map[int][]string{
	0: []string{"0"},
	1: []string{"1", "2", "4", "8"},
	2: []string{"3", "5", "6", "9", "10"},
	3: []string{"7", "11"},
}

func readBinaryWatch(num int) []string {
	if num < 0 || num > 10 {
		return nil
	}
	var res []string
	for i := 0; i <= num && i <= 3; i++ {
		hours := hourMap[i]
		for _, v := range hours {
			for j := 0; j < 60; j++ {
				if numberOfBinary(j) == num-i {
					res = append(res, fmt.Sprintf("%s:%02d", v, j))
				}
			}
		}
	}
	return res
}

func numberOfBinary(i int) int {
	var counter int
	for i > 0 {
		if i%2 == 1 {
			counter++
		}
		i /= 2
	}
	return counter
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
