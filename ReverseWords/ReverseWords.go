/* See http://code.google.com/codejam/contest/dashboard?c=351101#s=p1
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader)(string) {
	words := in.Words()

	credit := in.NextNum()
	items := in.NextNum()
	prices := in.NextNums( items)

	p0, p1 := pairSum(credit, prices)
	return fmt.Sprintf("%d %d", p0+1, p1+1)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
