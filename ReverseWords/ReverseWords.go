/* See http://code.google.com/codejam/contest/dashboard?c=351101#s=p1
 */

package main

import (
	"strings"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader)(string) {
	words := in.Words()

	reversed := make([]string, len(words))
	for pos, word := range(words) {
		reversed[len(words)-pos -1] = word
	}
	return strings.Join(reversed, " ")
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
