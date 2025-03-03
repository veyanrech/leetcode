package matrix

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	if rows == 0 || cols == 0 {
		return 0
	}

	fresh_count := 0
	rotten := [][2]int{} //0-row, 1-col

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				fresh_count++
			} else if grid[r][c] == 2 {
				rotten = append(rotten, [2]int{r, c})
			}
		}
	}

	if fresh_count == 0 {
		return 0
	}

	if len(rotten) == 0 {
		return -1
	}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	time_count := 0

	for len(rotten) > 0 && fresh_count > 0 {

		current_size := len(rotten)
		for i := 0; i < current_size; i++ {

			rot := rotten[0]
			rotten = rotten[1:]

			rot_row := rot[0]
			rot_col := rot[1]

			for _, di := range directions {
				check_item_row := rot_row + di[0]
				if check_item_row < 0 || check_item_row >= len(grid) {
					continue
				}

				check_item_col := rot_col + di[1]
				if check_item_col < 0 || check_item_col >= len(grid[0]) {
					continue
				}

				item := grid[check_item_row][check_item_col]
				if item == 1 {
					grid[check_item_row][check_item_col] = 2
					rotten = append(rotten, [2]int{check_item_row, check_item_col})
					fresh_count--
				}

			}
		}

		time_count++

	}

	if fresh_count > 0 {
		return -1
	}

	return time_count
}
