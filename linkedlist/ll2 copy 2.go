package linkedlist

func C2combinationSum4(types []int, target int) int {

	// var count int
	var min int = types[0]
	for _, v := range types {
		if v < min {
			min = v
		}
	}

	maxLength := target / min

	c := combs{
		combs:  make([][]int, 0),
		target: target,
		count:  0,
	}

	c.makeCombination(types, []int{}, maxLength, 0)

	// fmt.Println(len(c.combs))

	// for _, v := range c.combs {
	// 	fmt.Println("t", v)
	// 	sum := summirize(v)
	// 	if sum == target {
	// 		count++
	// 	}
	// }

	return c.count
}

type combs struct {
	combs  [][]int
	target int
	count  int
}

func (c *combs) makeCombination(types []int, comb []int, length int, summ int) {

	if summ == c.target {
		// fmt.Println("COMB return targeet", comb)
		c.count++
		return
	}

	// fmt.Println("COMB", comb)
	if length <= 0 {
		// fmt.Println("COMB return", comb)
		return
	}

	if summ > c.target {
		// fmt.Println("COMB return targeet", comb)
		return
	}

	for _, v := range types {
		var copyComb []int = make([]int, 0)
		newsumm := summ + v
		copyComb = copyy(comb)
		newCombo := append(copyComb, v)
		c.makeCombination(types, newCombo, length-1, newsumm)
	}
}

func copyy(src []int) []int {
	var retval []int = make([]int, 0)
	for _, v := range src {
		retval = append(retval, v)
	}
	return retval
}
