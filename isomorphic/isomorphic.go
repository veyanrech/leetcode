package isomorphic

func IsIsomorphic(s string, t string) bool {

	p1 := makePattern(s)
	p2 := makePattern(t)

	retvalBool := true

	for i, v := range p1 {
		if v != p2[i] {
			retvalBool = false
			break
		}
	}

	return retvalBool

}

func makePattern(s string) []int {
	m := map[byte]int{}
	var pattern []int = make([]int, len(s), len(s))
	ok := false
	patternItem := 0
	var v byte
	for i := 0; i < len(s); i++ {
		v = s[i]
		_, ok = m[v]
		if !ok {
			m[v] = len(m)
		}
		patternItem = m[v]
		pattern[i] = patternItem
	}
	return pattern
}
