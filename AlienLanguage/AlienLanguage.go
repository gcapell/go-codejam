/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func main() {
	in := &ProblemReader.In
	nums := in.Nums(3)
	wordLength, nWords, nCases := nums[0], nums[1], nums[2]
	fmt.Println("wordLength: ", wordLength)
	for j:=0; j<nWords; j++ {
		fmt.Println("word: ", in.Line())
	}
	for j:=0; j<nCases; j++ {
		fmt.Println("case: ", in.Line())
	}
}
