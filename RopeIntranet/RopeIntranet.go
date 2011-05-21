/* See http://code.google.com/codejam/contest/dashboard?c=619102#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader)(string) {
	n := in.Num()
	for j:=0; j<n; j++ {
		wire := in.Nums(2)
		fmt.Println("wire", j, wire)
	}
	return "dude"
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
