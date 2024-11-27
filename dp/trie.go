package dp

import "fmt"

func RunTest() {
	xmpl := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak("catsandog", xmpl))
}

func wordBreak(s string, wordDict []string) bool {
	trieExample := createTrie()
	max := 0
	for _, v := range wordDict {
		if len(v) > max {
			max = len(v)
		}
		trieExample.insert(v)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= maxf(i-max-1, 0); j-- {
			ok, _ := trieExample.searchPrefix(s[j:i])
			if ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func createTrie() *trie {
	return &trie{
		root: &trieNode{
			next:      map[string]*trieNode{},
			endOfWord: false,
		},
	}
}

type trie struct {
	root *trieNode
}

type trieNode struct {
	next      map[string]*trieNode
	endOfWord bool
}

func (t *trie) insert(s string) {
	node := t.root
	for _, char := range s {
		_, ok := node.next[string(char)]
		if !ok {
			node.next[string(char)] = &trieNode{
				next:      map[string]*trieNode{},
				endOfWord: false,
			}
		}
		node = node.next[string(char)]
	}
	node.endOfWord = true
}

func (t *trie) searchPrefix(s string) (bool, map[string]*trieNode) {
	node := t.root
	for _, char := range s {
		if node.next[string(char)] == nil {
			return false, nil
		}
		node = node.next[string(char)]
	}
	return node.endOfWord, node.next
}

func maxf(a, b int) int {
	if a > b {
		return a
	}
	return b
}
