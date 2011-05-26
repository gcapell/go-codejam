/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p2
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

const PATTERN = "welcome to code jam"

func solver(in *ProblemReader.ProblemReader) string {
	line := in.Line()
	scores := make([]int, len(PATTERN))
	for j := range line {
		c := line[j]
		for k := len(PATTERN)-1; k>=0; k-- {
			if c == PATTERN[k] {
				if k == 0 {
					scores[k] = (scores[k] + 1) % 10000
				} else {
					scores[k]  = (scores[k] + scores[k-1]) % 10000
				}
			}
		}
	}
	return fmt.Sprintf("%04d", scores[len(scores)-1])
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
