package groupsize

func GroupThePeople(groupSizes []int) [][]int {
	var m map[int][][]int = map[int][][]int{}
	for pearson, groupsize := range groupSizes {
		groups, ok := m[groupsize]
		if ok {
			currentGroup := groups[len(groups)-1]
			if len(currentGroup) < groupsize {
				currentGroup = append(currentGroup, pearson)
				groups[len(groups)-1] = currentGroup
				m[groupsize] = groups
			} else {
				m[groupsize] = append(m[groupsize], make([]int, 0))
				m[groupsize][len(groups)] = append(m[groupsize][len(groups)], pearson)
			}
		} else {
			m[groupsize] = make([][]int, 0, 1)
			m[groupsize] = append(m[groupsize], make([]int, 0))
			m[groupsize][0] = append(m[groupsize][0], pearson)
		}
	}

	var retval [][]int = make([][]int, 0)

	for _, v := range m {
		for _, v2 := range v {
			retval = append(retval, v2)
		}
	}

	return retval
}
