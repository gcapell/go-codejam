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

func main() {
	in := bufio.NewReader(os.Stdin)

	cases := nextNum(in)

	for j:=0; j< cases; j++ {
		credit := nextNum(in)
		items := nextNum(in)
		prices := nextNums(in, items)
		
		p0, p1 := pairSum(credit, prices)
		fmt.Printf("Case #%d: %d %d\n", j+1, p0+1, p1+1)
	}
}

/* Return indices of two elements of 'n' which sum to 'sum' */
func pairSum(sum int, n []int) (p0, p1 int) {

	// map element value to element position
	pos := make( map[int] int)

	for p1, val := range(n) {
		p0, ok := pos[sum-val]
		if ok {
			return p0, p1
		}
		pos[val] = p1
	}
	return
}

/* Read n nums from in */
func nextNums(in *bufio.Reader, n int) (nums [] int) {
	line := nextLine(in)

	numStrings := strings.Split(line, " ", n)
	
	nums = make([]int, n)

	for pos, numString := range(numStrings) {
		
		_, err := fmt.Sscanf(numString, "%d", &nums[pos])
		if err != nil {
			log.Fatalln("Sscan", err, numString)
		}
	}

	return nums
}

func nextNum(in *bufio.Reader) (n int){
	line := nextLine(in)

	if _, err := fmt.Sscanln(line, &n); err != nil {
		log.Fatalln("scanf", err)
	}
	return n
}

func nextLine(in *bufio.Reader) string {
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatalln("readstring", err)
	}
	return line
}
