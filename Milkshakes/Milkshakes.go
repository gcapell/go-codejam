/* See http://code.google.com/codejam/contest/dashboard?c=32016#s=p1&a=2
 */

package main

import (
	"fmt"
	"strings"
	"codejam/ProblemReader"
)

type customer struct {
	plain       []int // flavours we like plain
	likesMalted bool
	malted      int // flavour we like malted (if we like any malted)
}

func (c *customer) String() string {
	if c.likesMalted {
		return fmt.Sprintf("(plain: %v, malted: %d)", c.plain, c.malted)
	} else {
		return fmt.Sprintf("(plain: %v)", c.plain)
	}
	panic("customer.String")
}

func (c *customer) happyWith(malted map[int]bool) bool {
	for _, flavour := range c.plain {
		if !malted[flavour] {
			return true
		}
	}
	return c.likesMalted && malted[c.malted]
}

func newCustomer(in []int) *customer {
	c := new(customer)
	flavours, in := in[0], in[1:]
	for j := 0; j < flavours; j++ {
		flavour, malted := in[0], in[1]
		in = in[2:]
		if malted == 0 {
			c.plain = append(c.plain, flavour)
		} else {
			c.likesMalted = true
			c.malted = flavour
		}
	}
	return c
}

func solver(in *ProblemReader.ProblemReader) string {
	flavours := in.Num()
	nCustomers := in.Num()

	customers := make([]*customer, nCustomers)
	for j := 0; j < nCustomers; j++ {
		customers[j] = newCustomer(in.Nums())
	}

	malted := make(map[int]bool, flavours)
	for j := 0; j < flavours; j++ {
		malted[j+1] = false
	}

	keepLooking := true
	for keepLooking {
		keepLooking = false
		// Is there anybody unhappy?
		for _, c := range customers {
			if !c.happyWith(malted) {
				// Unhappy, but can be made happy with malt
				if c.likesMalted && !malted[c.malted] {
					malted[c.malted] = true
					keepLooking = true
				} else {
					return "IMPOSSIBLE"
				}
			}
		}
	}

	reply := make([]string, flavours)
	for j := 0; j < flavours; j++ {
		reply[j] = ternary(malted[j+1], "1", "0")
	}
	return strings.Join(reply, " ")
}

func ternary(b bool, x string, y string) string {
	if b {
		return x
	}
	return y
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
