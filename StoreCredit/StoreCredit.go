/* See http://code.google.com/codejam/contest/dashboard?c=351101#s=p0
 */

package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	cases := nextNum(in)

	for j:=0; j< cases; j++ {
		credit := nextNum(in)
		items := nextNum(in)
		prices := nextLine(in)
		fmt.Println("credit:", credit, "items: ", items, "prices:" , prices)
	}
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
