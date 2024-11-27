package slidingwindow

func findSubstring(s string, words []string) []int {
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

	i, j := 0, 0

	wordForSearch := []byte{}
	wordsCopy := newWordsControl(wordsCache)
	for j < lens {
		_, ok := lettersCache[s[j]]
		if !ok {
			j++
			i = j
			wordForSearch = []byte{}
			wordsCopy = newWordsControl(wordsCache)
			continue
		}
		wordForSearch = append(wordForSearch, s[j])
		if len(wordForSearch)%oneWordLen == 0 {

			_, wordok := wordsCache[s[j-oneWordLen+1:j+1]]
			if wordok {
				if wordsCopy.m[s[j-oneWordLen+1:j+1]] > 0 && wordsCopy.total > 0 {
					wordsCopy.m[s[j-oneWordLen+1:j+1]] -= 1
					wordsCopy.total -= 1

					if wordsCopy.total == 0 {
						res = append(res, i)
						wordForSearch = append([]byte{}, wordForSearch[oneWordLen:]...)
						i = i + oneWordLen

						//reset
						wordsCopy = newWordsControl(wordsCache)

						for k := 0; k+oneWordLen <= len(wordForSearch); k += oneWordLen {
							_s := wordForSearch[k : k+oneWordLen]
							wordsCopy.m[string(_s)] -= 1
							wordsCopy.total -= 1
						}
					}

				} else {
					//doubles
					i = j - (oneWordLen*lenw - lenw) + 1
					j++
					wordForSearch = append([]byte{}, s[i:j]...)
					wordsCopy = newWordsControl(wordsCache)
					for k := 0; k+oneWordLen <= len(wordForSearch); k += oneWordLen {
						_s := wordForSearch[k : k+oneWordLen]
						wordsCopy.m[string(_s)] -= 1
						wordsCopy.total -= 1

						if wordsCopy.m[string(_s)] < 0 {
							wordsCopy.total = 0
						}
					}
					continue
				}
			}

		}
		j++
	}

	return res

}

func copyMap(oldMap map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range oldMap {
		res[k] = v
	}
	return res
}

type wordsControl struct {
	total int
	m     map[string]int
}

func newWordsControl(sourceMap map[string]int) *wordsControl {

	total := 0

	for _, v := range sourceMap {
		total += v
	}

	return &wordsControl{
		total: total,
		m:     copyMap(sourceMap),
	}
}

func Run() []int {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}

	return findSubstring(s, words)
}
