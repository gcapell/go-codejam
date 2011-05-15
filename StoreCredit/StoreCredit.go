/* See http://code.google.com/codejam/contest/dashboard?c=351101#s=p0
 */

package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"strings"
)

type ProblemReader bufio.Reader

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

/* Read n nums from in */
func (in *ProblemReader) nextNums(n int) (nums []int) {
	line := nextLine(in)

	numStrings := strings.Split(line, " ", n)

	nums = make([]int, n)

	for pos, numString := range numStrings {
		_, err := fmt.Sscanf(numString, "%d", &nums[pos])
		if err != nil {
			log.Fatalln("Sscan", err, numString)
		}
	}

	return nums
}

func (in *ProblemReader)nextNum() (n int) {
	line := nextLine(in)
	if _, err := fmt.Sscanln(line, &n); err != nil {
		log.Fatalln("scanf", err)
	}
	return n
}

func nextLine(in *ProblemReader) string {
	line, err := (*bufio.Reader)(in).ReadString('\n')
	if err != nil {
		log.Fatalln("readstring", err)
	}
	return line
}

func (in *ProblemReader) solveProblems( solve func(*ProblemReader)string) {
	cases := in.nextNum()

	for j := 0; j < cases; j++ {
		fmt.Printf("Case #%d: %s\n", j+1, solve(in))
	}
}

func solver(in *ProblemReader)(string) {
	credit := in.nextNum()
	items := in.nextNum()
	prices := in.nextNums( items)

	p0, p1 := pairSum(credit, prices)
	return fmt.Sprintf("%d %d", p0+1, p1+1)
}

func main() {
	in := (*ProblemReader)(bufio.NewReader(os.Stdin))

	in.solveProblems(solver)
}
