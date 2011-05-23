/* See http://code.google.com/codejam/contest/dashboard?c=189252#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

var lastVal = -1	// last value assigned to a letter

func solver(in *ProblemReader.ProblemReader) string {
	word := []byte(in.Line())
	lastVal = -1

	letterVal := make (map[byte] int)
	for _,c := range word {
		if _, ok := letterVal[c]; !ok {
			letterVal[c] = nextVal()
		}
	}

	base := int64(lastVal + 1)
	if base == 1 {
		base++
	}
	sum := int64(0)
	mul := int64(1)
	
	//fmt.Println("letterVal: ", letterVal, "word:", word)

	for j := len(word)-1; j>=0; j-- {
		sum += mul * int64(letterVal[word[j]])
		//fmt.Println("letter: ", word[j], "val: ", letterVal[word[j]], "mul: ", mul, "sum: ", sum)
		mul *= base
	}
	return fmt.Sprintf("%d", sum)
}

func nextVal() int {
	switch lastVal {
	case -1: lastVal = 1
	case 1: lastVal = 0
	case 0: lastVal = 2
	default: lastVal++
	}
	return lastVal
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
