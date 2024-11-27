package calc

import (
	"fmt"
	"strconv"
)

func RunTestcalculate() {
	// s := "(1+(43+5+2)-3)+(6+8)"
	// s1 := "1-(     -2)"
	// s := "1-11"
	// s := "(3-(2-(5-(9-(4)))))"
	//3-2+5-9+4
	// s := "(7)-(0)+(4)"
	s := "2-4-(8+2-6+(8+4-(1)+8-10))"
	res := calculate(s)

	fmt.Println(res)
}

func calculate(s string) int {
	var result int

	// sarr := make([]string, 0, len(s)+1)
	// sarr = append(sarr, strings.Split(s, "")...)
	prevCifraStack := []int{}
	plus := 1 //0-false 1-true
	sarrv := ""
	for i := 0; i < len(s); i++ {
		sarrv = string(s[i])
		// fmt.Println(sarrv)
		switch sarrv {
		case "(":
			prevCifraStack = append(prevCifraStack, result, plus)
			result = 0
			plus = 1
		case ")":
			result = result * pop(&prevCifraStack) //sign
			result += pop(&prevCifraStack)
		case "+":
			plus = 1
		case "-":
			plus = -1
		case " ":
			continue
		default:
			prevNum := 0
			for i < len(s) {
				convsarrv, err := strconv.Atoi(string(s[i]))
				if err == nil {
					prevNum = prevNum*10 + convsarrv
					i++
				} else {
					break
				}
			}
			result += prevNum * plus
			i--
		}

	}

	return result
}

func pop(a *[]int) int {
	temp := (*a)[len((*a))-1]
	(*a) = (*a)[:len((*a))-1]
	return temp
}

func peek(a *[]int) int {
	temp := (*a)[0]
	(*a) = (*a)[1:]
	return temp
}
