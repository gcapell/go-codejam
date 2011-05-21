/* See http://code.google.com/codejam/contest/dashboard?c=635101#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	nums := in.Nums(2)
	existing, toCreate := nums[0], nums[1]
	for j :=0; j<existing; j++ {
		fmt.Println("existing:", in.Line())
	}
	for j :=0; j<toCreate; j++ {
		fmt.Println("toCreate:", in.Line())
	}

	return fmt.Sprintf("Dude")
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
