/* See http://code.google.com/codejam/contest/dashboard?c=351101#s=p0
 */

package main

import (
	"codejam/ProblemReader"
	"fmt"
)

/* Return indices of two elements of 'n' which sum to 'sum' */
func pairSum(sum int, n []int) (p0, p1 int) {

	// map element value to element position
	pos := make(map[int]int)

	for p1, val := range n {
		p0, ok := pos[sum-val]
		if ok {
			return p0, p1
		}
		pos[val] = p1
	}
	return
}

func solver(in *ProblemReader.ProblemReader)(string) {
	credit := in.NextNum()
	items := in.NextNum()
	prices := in.NextNums( items)

	p0, p1 := pairSum(credit, prices)
	return fmt.Sprintf("%d %d", p0+1, p1+1)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
