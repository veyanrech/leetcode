package brackets

import "fmt"

func RunBRTest() {
	s := "(){}[](((((((({]}))))))))"
	fmt.Println("result:", br(s))
}

func br(s string) bool {
	stack := []string{}
	for _, v := range s {
		sv := string(v)
		switch sv {
		case "(", "[", "{":
			stack = append(stack, sv)
		case "}", "]", ")":
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				if (last == "{" && sv == "}") || (last == "(" && sv == ")") || (last == "[" && sv == "]") {
					stack = append([]string{}, stack[:len(stack)-1]...)
				} else {
					stack = append(stack, ".")
				}
			} else {
				stack = append(stack, ".")
			}
		}
	}

	return len(stack) == 0
}
