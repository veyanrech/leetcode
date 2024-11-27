package strng

import "fmt"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	runes := []rune(s)
	zigzag := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		zigzag[i] = []rune{}
	}

	fmt.Println(zigzag)

	down := false

	// col := 0
	row := 0

	for si := 0; si < len(runes); {
		zigzag[row] = append(zigzag[row], runes[si])

		if down {
			row--
		} else {
			row++
		}

		if row > numRows-1 {
			down = true
			// col++
			row -= 2
		}

		if row < 0 {
			down = false
			// col++
			row += 2
		}

		si++
	}

	final := []rune{}
	for i := 0; i < numRows; i++ {
		final = append(final, zigzag[i]...)
	}

	return string(final)
}

func RunZigzag() {
	s := "PAYPALISHIRING"
	convert(s, 3)
}
