package strng

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	finalWords := []string{}
	res := [][]string{}
	rowLengths := []int{}

	row := make([]string, 0)

	rowLength := 0

	for i := 0; i < len(words); {
		word := words[i]

		if rowLength > 0 {
			if rowLength+1+len([]rune(word)) <= maxWidth {
				row = append(row, word)
				rowLength += len(word) + 1
				i++
			} else {
				res = append(res, row)
				row = make([]string, 0)
				rowLengths = append(rowLengths, rowLength)
				rowLength = 0
			}
		} else {
			if rowLength+len([]rune(word)) <= maxWidth {
				row = append(row, word)
				rowLength += len(word)
				i++
			} else {
				res = append(res, row)
				row = make([]string, 0)
				rowLengths = append(rowLengths, rowLength)
				rowLength = 0
			}
		}
	}

	res = append(res, row)

	for r := 0; r < len(res)-1; r++ {
		row := res[r]
		rowLength := countlen(row)

		numberOfSpaces := maxWidth - rowLength
		numberOfWords := len(row)

		spacePerPlace := 0
		isNumberHalf := 0
		if numberOfWords == 1 {
			spacePerPlace = numberOfSpaces
		} else {
			spacePerPlace = numberOfSpaces / (numberOfWords - 1)
			isNumberHalf = numberOfSpaces % (numberOfWords - 1)
		}

		for i := 0; i < len(row); i++ {
			spaces := []rune{}

			for j := 0; j < spacePerPlace && numberOfSpaces > 0; j++ {
				spaces = append(spaces, ' ')
				numberOfSpaces--
			}
			if isNumberHalf > 0 {
				spaces = append(spaces, ' ')
				isNumberHalf--
				numberOfSpaces--
			}

			row[i] = string(append([]rune(row[i]), spaces...))
		}

		finalWords = append(finalWords, strings.Join(row, ""))
	}

	lastRow := res[len(res)-1]
	lastrowstring := strings.Join(lastRow, " ")
	lastrowstring = lastrowstring + strings.Repeat(" ", maxWidth-len(lastrowstring))
	finalWords = append(finalWords, lastrowstring)

	fmt.Println(finalWords)
	return finalWords
}

func countlen(row []string) int {
	length := 0
	for _, row := range row {
		length += len([]rune(row))
	}
	return length
}

func RunFullJustify() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	fullJustify(words, 16)
}
