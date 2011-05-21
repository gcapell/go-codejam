/* See http://code.google.com/codejam/contest/dashboard?c=544101#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	nums := in.Nums(2)
	board, toWin := nums[0], nums[1]
	fmt.Println("board:", board, "toWin:", toWin)
	for j := 0; j < board; j++ {
		fmt.Println("line", in.Line())
	}
	return "Dude!"
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
