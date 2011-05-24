package ProblemReader

import (
	"fmt"
	"strings"
	"strconv"
	"log"
	"os"
	"encoding/line"
)

const (
	MAX_LINE = 1 << 15
)

type ProblemReader struct {
	r *line.Reader
}

var In = ProblemReader{line.NewReader(os.Stdin, MAX_LINE)}

func (in *ProblemReader) SolveProblems(solve func(*ProblemReader) string) {
	cases := in.Num()

	for j := 0; j < cases; j++ {
		fmt.Printf("Case #%d: %s\n", j+1, solve(in))
	}
}

/* Read nums from in */
func (in *ProblemReader) Nums() (nums []int) {
	words := in.Words()
	nums = make([]int, len(words))

	for pos, word := range words {
		nums[pos] = atoi(word)
	}

	return nums
}

/* Read n nums from in */
func (in *ProblemReader) NNums(n int) (nums []int) {
	nums = in.Nums()
	if len(nums) != n {
		log.Fatalln("expected %d nums, got %d (%v)", n, len(nums), nums)
	}
	return nums
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("atoi problem", s, err)
	}
	return n
}

func (in *ProblemReader) Words() []string {
	return strings.Fields(in.Line())
}

func (in *ProblemReader) Num() (n int) {
	words := in.Words()
	if len(words) != 1 {
		log.Fatalln("Expected one number, got: ", words)
	}
	return atoi(words[0])
}

func (in *ProblemReader) Line() string {
	line, isPrefix, err := in.r.ReadLine()
	if err != nil {
		log.Fatalln("readstring", err)
	}
	if isPrefix {
		log.Fatalln("line too long")
	}
	return string(line)
}
