package ransomnote

func canConstruct(ransomNote string, magazine string) bool {

	rm := map[rune]int{}
	ok := true
	iv := 0

	for _, v := range ransomNote {
		_, ok = rm[v]
		if ok {
			rm[v]++
		} else {
			rm[v] = 1
		}
	}

	for _, v := range magazine {
		iv, ok = rm[v]
		if ok {
			if iv > 0 {
				rm[v]--
			}
		}
	}

	result := true

	for _, v := range rm {
		if v > 0 {
			result = false
			break
		}
	}

	return result

}
