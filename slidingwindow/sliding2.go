package slidingwindow

func findSubstring2(s string, words []string) []int {
	res := []int{}

	lens := len(s)
	lenw := len(words)
	oneWordLen := len(words[0])
	wordsCache := map[string]int{}
	lettersCache := map[byte]struct{}{}

	for i := 0; i < lenw; i++ {
		v, ok := wordsCache[words[i]]
		if !ok {
			wordsCache[words[i]] = 1

			word := words[i]
			for w := 0; w < len(word); w++ {
				_, wok := lettersCache[word[w]]
				if !wok {
					lettersCache[word[w]] = struct{}{}
				}
			}

		} else {
			wordsCache[words[i]] = v + 1
		}
	}

	l, r := 0, 0

	wordForSearch := []byte{}
	wordsCopy := newWordsControl2(wordsCache)

	for l < lens-(oneWordLen*lenw) {
		_, ok := lettersCache[s[l]]
		if !ok {
			l++
			r = l
			wordForSearch = []byte{}
			wordsCopy = newWordsControl2(wordsCache)
			continue
		}

		wordForSearch = append(wordForSearch, s[r])
		if len(wordForSearch)%oneWordLen == 0 {

		}

		r++
	}

	return res

}

func copyMap2(oldMap map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range oldMap {
		res[k] = v
	}
	return res
}

type wordsControl2 struct {
	total int
	m     map[string]int
}

func newWordsControl2(sourceMap map[string]int) *wordsControl {

	total := 0

	for _, v := range sourceMap {
		total += v
	}

	return &wordsControl{
		total: total,
		m:     copyMap(sourceMap),
	}
}

func Run2() []int {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}

	return findSubstring(s, words)
}
