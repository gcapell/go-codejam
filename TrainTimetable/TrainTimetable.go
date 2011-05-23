/* See http://code.google.com/codejam/contest/dashboard?c=32013#s=p1
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	turnAround := in.Num()
	fmt.Println("turnAround", turnAround)
	nums := in.Nums(2)
	nA, nB := nums[0], nums[1]
	for j :=0; j< nA; j++ {
		fmt.Println("schedA", in.Line())
	}
	for j :=0; j< nB; j++ {
		fmt.Println("schedB", in.Line())
	}
	return "dude"
}


func main() {
	ProblemReader.In.SolveProblems(solver)
}
