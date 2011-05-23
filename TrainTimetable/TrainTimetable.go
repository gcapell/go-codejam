/* See http://code.google.com/codejam/contest/dashboard?c=32013#s=p1
 */

package main

import (
	"container/heap"
	"container/vector"
	"time"
	"log"
	"sort"
	"fmt"
	"codejam/ProblemReader"
)

const (
	A= "A"
	B= "B"
)

type trip struct {
	src, dst string
	depart, arrive int
}

type myHeap struct {
	vector.IntVector
}
func (h *myHeap) Less(i, j int) bool { return h.At(i) < h.At(j) }

type station struct {
	reserved int
	waiting myHeap
}

var (
	allStations = map[string] station {A: station{},B: station{}}
)

func newStation() station {
	return station{}
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
		allTrips[j] = readTrip(in, A, B)
	}
	for j :=0; j< nB; j++ {
		allTrips[nA +j ] = readTrip(in, B, B)
	}
	sort.Sort(tripArray(allTrips))
	
	for _, t := range allTrips {
		fmt.Println(t)
		src, dst := allStations[t.src], allStations[t.dst]
		src.getTrain(t.depart)
		dst.addTrain(t.arrive + turnAround)
	}

	sum := 0
	for _,v := range allStations {
		sum += v.reserved
	}

	return fmt.Sprintf("%d", sum)
}

func (s station) getTrain(departure int) {
	h := s.waiting
	if h.Len() > 0  && h.At(0) < departure {
		heap.Pop(&h)
	} else {
		s.reserved++
	}
}

func (s station) addTrain(arrival int) {
	heap.Push(s.waiting, interval)
}

func readTrip(in *ProblemReader.ProblemReader, src string, dst string) trip {
	sched := in.Words()
	return trip { src , dst, parseT(sched[0]), parseT(sched[1]) }
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

