package colorgame

import (
	"fmt"
)

func WinnerOfGame(colors string) bool {
	if len(colors) < 3 {
		return false
	}

	tripleColors := map[string]int{"A": 0, "B": 0}

	tripleCounterFlag := map[string]int{"A": 0, "B": 0}

	i := 0
	v := colors[i]
	for i < len(colors) {
		v = colors[i]
		// fmt.Println(i, string(v))
		if i > 0 {
			if string(v) == string(colors[i-1]) {
				tripleCounterFlag[string(v)] += 1
			} else {
				tripleCounterFlag = map[string]int{"A": 0, "B": 0}
				tripleCounterFlag[string(v)] += 1
			}
		} else {
			tripleCounterFlag[string(v)] += 1
		}

		// fmt.Println("tripleCounterFlag", tripleCounterFlag, i, string(v))

		if tripleCounterFlag[string(v)] == 3 {
			tripleColors[string(v)] += 1
			tripleCounterFlag[string(v)] = 0
			i--
			// fmt.Println("рефрегш", i, string(v), string(colors[i-1]))
		} else {
			i++
		}
	}

	fmt.Println(tripleColors)

	if tripleColors["B"] > tripleColors["A"] {
		return false
	}

	if tripleColors["A"] == 0 {
		return false
	}

	return true
}
