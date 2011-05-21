/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p0
 */

package main

import (
	"fmt"
	"strings"
	"log"
	"codejam/ProblemReader"
)

type Trie map[byte]Trie

func newTrie() Trie {
	return Trie(make(map[byte]Trie))
}

func (trie Trie) add(word []byte) {
	if len(word) == 0 {
		return
	}
	c, remainder := word[0], word[1:]
	child, ok := trie[c]
	if !ok {
		child = newTrie()
		trie[c] = child
	}
	child.add(remainder)
}

func splitParen(pattern []byte) (alternates []byte, remainder []byte) {
	s := string(pattern)
	pos := strings.Index(s, ")")
	if pos == -1 {
		log.Fatalln("expected ) in ", s)
	}
	alternates = pattern[:pos]
	remainder = pattern[pos+1:]

	return alternates, remainder
}

/* Count matches of 'pattern' in 'trie' */
func (trie Trie) count(pattern []byte) int {
	if len(pattern) == 0 {
		return 1
	}
	c, remainder := pattern[0], pattern[1:]
	sum := 0
	switch c {
	case '(':
		alternates, remainder := splitParen(remainder)
		for _, c := range alternates {
			sum += trie.count1(c, remainder)
		}
	default:
		return trie.count1(c, remainder)
	}
	return sum
}

func (trie Trie) count1(c byte, remainder []byte) int {
	child, ok := trie[c]
	if !ok {
		return 0
	}
	return child.count(remainder)
}

func main() {
	in := &ProblemReader.In
	nums := in.Nums(3)
	_, nWords, nCases := nums[0], nums[1], nums[2]

	trie := newTrie()

	for j := 0; j < nWords; j++ {
		word := in.Line()
		trie.add([]byte(word))
	}
	for j := 0; j < nCases; j++ {
		pattern := in.Line()
		matches := trie.count([]byte(pattern))
		fmt.Printf("Case #%d: %d\n", j+1, matches)
	}
}
