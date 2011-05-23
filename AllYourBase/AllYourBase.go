/* See http://code.google.com/codejam/contest/dashboard?c=189252#s=p0
 */

package main

import (
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	word := in.Line()
	return string(word)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
