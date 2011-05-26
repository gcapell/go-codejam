/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p2
 */

package main

import (
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	return in.Line()
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
