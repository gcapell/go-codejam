/* See http://code.google.com/codejam/contest/dashboard?c=619102#s=p0
 */

package main

import (
	"fmt"
	"sort"
	"codejam/ProblemReader"
)

type Wire struct {
	a, b int
}

type WireSlice []Wire

func (ws WireSlice) Len() int {
	return len(ws)
}

func (ws WireSlice) Less(i, j int) bool {
	return ws[i].a < ws[j].a
}

func (ws WireSlice) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}


func solver(in *ProblemReader.ProblemReader) string {
	n := in.Num()
	wires := make([]Wire, n)

	for j := 0; j < n; j++ {
		wire := in.NNums(2)
		wires[j] = Wire{wire[0], wire[1]}
	}
	sort.Sort(WireSlice(wires))

	cross := 0
	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if wires[k].b < wires[j].b {
				cross++
			}
		}
	}
	return fmt.Sprintf("%d", cross)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
