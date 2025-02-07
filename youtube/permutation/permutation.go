// You can edit this code!
// Click here and start typing.
package permutation

import "fmt"

func perm(a []int) [][]int {

	res := [][]int{}

	amap := []map[int]int{}

	for i := 0; i < len(a); i++ {

		digitsMap := map[int]int{}

		num := a[i]

		for num > 0 {
			d := num % 10

			if d == 0 {
				num /= 10
				continue
			} else {
				num /= 10
				digitsMap[d]++
			}

		}

		amap = append(amap, digitsMap)

	}

	for i := 0; i < len(a); i++ {

		if a[i] == -1 {
			continue
		}

		portion := []int{}
		portion = append(portion, a[i])

		for j := i + 1; j < len(a); j++ {

			if a[j] == -1 {
				continue
			}

			sorcemap := amap[i]
			destmap := amap[j]
			possible := false
			for k, v := range sorcemap {
				possible = true
				if destmap[k] != v {
					possible = false
					break
				}
			}
			if possible {
				portion = append(portion, a[j])
				a[j] = -1
			}
		}

		res = append(res, portion)
	}

	return res

}

func RUN() {

	fmt.Println("starting")

	fmt.Println(perm([]int{1231, 1300201, 909090, 102030, 999, 10000023, 9, 91919, 19199, 0, 2, 3, 5}))
}
