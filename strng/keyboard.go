package strng

import (
	"fmt"
	"strings"
)

func RunfindWords() {
	testcase := []string{"Hello", "Alaska", "Dad", "Peace"}
	findWords(testcase)
}

func findWords(words []string) []string {
	row1 := "qwertyuiop"
	row2 := "asdfghjkl"
	row3 := "zxcvbnm"

	row1 = fmt.Sprintf("%s%s", row1, strings.ToUpper(row1))
	row2 = fmt.Sprintf("%s%s", row2, strings.ToUpper(row2))
	row3 = fmt.Sprintf("%s%s", row3, strings.ToUpper(row3))

	rows := []string{row1, row2, row3}

	rowsmap := make([]map[byte]bool, 3)

	for i, row := range rows {
		rowsmap[i] = map[byte]bool{}
		for j := 0; j < len(row); j++ {
			rowsmap[i][row[j]] = true
		}
	}

	res := []string{}

	for _, word := range words {
		wordlength := len(word)

		fl := word[0]
		rowchoosen := -1

		for i, rowmap := range rowsmap {
			_, ok := rowmap[fl]
			if ok {
				rowchoosen = i
				break
			}
		}

		counter := 0

		for wi := 0; wi < len(word); wi++ {
			letter := word[wi]
			_, ok := rowsmap[rowchoosen][letter]
			if ok {
				counter++
			} else {
				break
			}
		}

		if counter == wordlength {
			res = append(res, word)
		}
	}

	return res
}
