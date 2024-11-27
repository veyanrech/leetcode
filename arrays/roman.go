package arrays

import "fmt"

func RunromanToInt() {
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0

	prev := ""

	for i := len(s) - 1; i >= 0; i-- {
		st := string(s[i])
		if i <= len(s)-2 {
			prev = string(s[i+1])
			switch st {
			case "I":
				if prev == "V" || prev == "X" {
					sum -= 1
				} else {
					sum += symbols[st]
				}
			case "X":
				if prev == "L" || prev == "C" {
					sum -= 10
				} else {
					sum += symbols[st]
				}
			case "C":
				if prev == "D" || prev == "M" {
					sum -= 100
				} else {
					sum += symbols[st]
				}
			default:
				sum += symbols[st]
			}
		} else {
			sum += symbols[st]
		}
		fmt.Println(sum)
	}

	return sum
}
