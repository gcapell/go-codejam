/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p1
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	hw := in.NNums(2)
	height, width:= hw[0], hw[1]
	for j :=0; j< height; j++ {
		altitude := in.NNums(width)
		fmt.Println(altitude)
	}
	return "Dude"
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
