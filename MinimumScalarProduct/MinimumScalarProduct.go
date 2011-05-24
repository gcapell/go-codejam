/* See http://code.google.com/codejam/contest/dashboard?c=32016#s=p0
 */

package main

import (
	"fmt"
	"sort"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	n := in.Num()
	v1 := in.NNums(n)
	v2 := in.NNums(n)

	sort.SortInts(v1)
	sort.SortInts(v2)

	var sum int64 = 0

	for j := 0; j < n; j++ {
		sum += int64(v1[j]) * int64(v2[n-j-1])
	}
	return fmt.Sprint(sum)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
