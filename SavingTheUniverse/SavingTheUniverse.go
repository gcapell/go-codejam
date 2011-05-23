/* See http://code.google.com/codejam/contest/dashboard?c=189252#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

var lastVal = -1 // last value assigned to a letter

func solver(in *ProblemReader.ProblemReader) string {
	nEngines := in.Num()
	
	engines := make ([]string, nEngines)
	for j:=0; j<nEngines; j++ {
		engines[j] = in.Line()
	}

	terms := in.Num()
	available := make(map[string] bool)
	reset(available, engines, "")
	changes := 0
	for j:=0; j<terms; j++ {
		term := in.Line()
		if _, ok := available[term]; ok {
			available[term] = false, false
			if len(available) == 0 {
				changes++
				reset(available, engines, term)
			}
		}
	}
	return fmt.Sprintf("%d", changes)
}

func reset(available map[string]bool, engines []string, lastUsed string ) {
	for _, s := range engines {
		if s != lastUsed {
			available[s] = true
		}
	}
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
