/* See http://code.google.com/codejam/contest/dashboard?c=32013#s=p1
 */

package main

import (
	"time"
	"log"
	"sort"
	"fmt"
	"codejam/ProblemReader"
)

type trip struct {
	src string
	depart, arrive int
}

func (t trip) String() string {
	return fmt.Sprintf("%s %d -> %d", t.src, t.depart, t.arrive)
}

type tripArray [] trip

func solver(in *ProblemReader.ProblemReader) string {
	turnAround := in.Num()
	fmt.Println("turnAround", turnAround)
	nums := in.Nums(2)
	nA, nB := nums[0], nums[1]
	allTrips := make([]trip, nA + nB)
	
	for j :=0; j< nA; j++ {
		allTrips[j] = readTrip(in, "A")
	}
	for j :=0; j< nB; j++ {
		allTrips[nA +j ] = readTrip(in, "B")
	}
	sort.Sort(tripArray(allTrips))
	fmt.Println(allTrips)
	return "dude"
}

func readTrip(in *ProblemReader.ProblemReader, src string) trip {
	sched := in.Words()
	return trip { src , parseT(sched[0]), parseT(sched[1]) }
}

func parseT(s string) int {
	t, error := time.Parse("15:04", s)
	if error != nil {
		log.Fatalln("problem", error, "parsing", s)
	}
	return t.Hour*60 + t.Minute
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}

func (t tripArray) Len() int {
	return len(t)
}

func (t tripArray) Less(i,j int) bool {
	return t[i].depart < t[j].depart
}

func (t tripArray) Swap(i,j int) {
	t[i], t[j] = t[j], t[i]
}

