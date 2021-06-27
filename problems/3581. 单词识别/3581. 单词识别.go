package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	wordCount := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.ToLower(s)
	sr := []rune(s)
	for i, j := 0, 0; i < len(sr); i++ {
		if !unicode.IsLetter(sr[i]) {
			continue
		}
		j = i + 1
		for j < len(sr) {
			if unicode.IsLetter(sr[j]) {
				j++
			} else {
				break
			}
		}
		word := string(sr[i:j])
		wordCount[word]++
		i = j

	}
	words := make([]string, 0, len(wordCount))
	for word := range wordCount {
		words = append(words, word)
	}
	sort.Strings(words)
	for _, word := range words {
		fmt.Printf("%s:%d\n", word, wordCount[word])
	}
}

// abc def ghi
