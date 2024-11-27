package arrays

func TestMatrix() {
	m := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m)
}

func setZeroes(matrix [][]int) {
	i := 0
	j := len(matrix)

	i2 := 0

	m := map[int]bool{}
	m2 := map[int]bool{}

	for i <= j {
		j--
		i2 = 0

		for i2 < len(matrix[i]) {

			vi2 := matrix[i][i2]
			vj2 := matrix[j][i2]

			if vi2 == 0 {
				m[i] = true
				m2[i2] = true
			}

			if vj2 == 0 {
				m[j] = true
				m2[i2] = true
			}

			i2++
		}
		i++
	}

	i = 0
	for i < len(matrix) {

		_, zerIOk := m[i]

		if zerIOk {
			matrix[i] = make([]int, len(matrix[i]), len(matrix[i]))
		} else {
			i2 = 0

			for i2 < len(matrix[i]) {

				_, zerIOk2 := m2[i2]

				if zerIOk2 {
					matrix[i][i2] = 0
				}

				i2++
			}
		}
		i++
	}

	// fmt.Println(0)
}
