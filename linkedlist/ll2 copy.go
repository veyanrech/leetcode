package linkedlist

func CcombinationSum4(types []int, target int) int {
	var count int
	var min int = types[0]
	for _, v := range types {
		if v < min {
			min = v
		}
	}

	maxLength := target / min

	for combination := range generateCombination(types, maxLength) {
		sum := summirize(combination)
		if sum == target {
			count++
		}
	}

	return count
}

func generateCombination(ints []int, length int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		addInt(c, []int{}, ints, length)
	}(c)
	return c
}

func addInt(c chan []int, combo []int, ints []int, length int) {
	if length <= 0 {
		return
	}

	var newCombo []int
	for _, ii := range ints {
		newCombo = append(combo, ii)
		c <- newCombo
		addInt(c, newCombo, ints, length-1)
	}
}

func summirize(ints []int) int {
	var total int = 0
	for _, v := range ints {
		total += v
	}
	// fmt.Println("sU", ints, total)
	return total
}
