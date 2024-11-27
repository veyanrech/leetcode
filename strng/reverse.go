package strng

import "fmt"

func reverseWords(s string) string {
	slen := len(s)
	newstr := []byte{}

	end := slen - 1
	start := end

	for start >= 0 {

		ss := s[start]
		se := s[end]

		fmt.Println("start: ", start, " end: ", end, " ss: ", string(ss), " se: ", string(se))

		if s[start] == 32 && s[end] == 32 {
			start--
			end--
		} else {
			if start >= 0 && s[start] == 32 && s[end] != 32 {
				if len(newstr) > 0 {
					newstr = append(newstr, byte(32))
				}

				newstr = append(newstr, s[start+1:end+1]...)

				start--
				end = start
			} else if start == 0 && s[start] != 32 && s[end] != 32 {
				if len(newstr) > 0 {
					newstr = append(newstr, byte(32))
				}
				newstr = append(newstr, s[start:end+1]...)
				start--
				end = start
			} else {
				start--
			}
		}
	}

	// newstr = append([]byte{}, newstr[0:len(newstr)]...)

	return string(newstr)
}

func RunReverse() {
	s := " asdasd df f"
	fmt.Println(reverseWords(s))
}
