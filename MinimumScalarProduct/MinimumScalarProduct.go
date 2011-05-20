/* See http://code.google.com/codejam/contest/dashboard?c=32016#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader)(string) {
	n := in.Num()
	v1 := in.Nums(n)
	v2 := in.Nums(n)

	return fmt.Sprint("v1: ", v1, " v2: ", v2)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
