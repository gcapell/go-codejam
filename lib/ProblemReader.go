package ProblemReader

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"os"
)

type ProblemReader bufio.Reader

var	In = (*ProblemReader)(bufio.NewReader(os.Stdin))

func (in *ProblemReader) SolveProblems( solve func(*ProblemReader)string) {
	cases := in.Num()

	for j := 0; j < cases; j++ {
		fmt.Printf("Case #%d: %s\n", j+1, solve(in))
	}
}

/* Read n nums from in */
func (in *ProblemReader) Nums(n int) (nums []int) {
	line := in.Line()

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

func (in *ProblemReader) Words() ([]string) {
	line := in.Line()
	return strings.Split(line, " ", -1)
}

func (in *ProblemReader) Num() (n int) {
	line := in.Line()
	if _, err := fmt.Sscanln(line, &n); err != nil {
		log.Fatalln("scanf", err)
	}
	return n
}

func (in *ProblemReader)Line() string {
	line, err := (*bufio.Reader)(in).ReadString('\n')
	if err != nil {
		log.Fatalln("readstring", err)
	}
	return line
}
