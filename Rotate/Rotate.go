/* See http://code.google.com/codejam/contest/dashboard?c=544101#s=p0
 */

package main

import (
	"fmt"
	"log"
	"codejam/ProblemReader"
)

func assert(b bool) {
	if !b {
		panic("assertion")
	}
}
func solver(in *ProblemReader.ProblemReader) string {
	nums := in.Nums(2)
	board, toWin := nums[0], nums[1]
	assert(toWin <= board)
	lines := make([]string, board)
	for j := 0; j < board; j++ {
		line := []byte(in.Line())
		if len(line) != board {
			log.Fatalf("Expected %#v to be %d long", string(line), board)
		}
		shiftLine(line)
		lines[j] = string(line)
	}
	return winner(lines, toWin)
}

func shiftLine(line []byte) {
	writePos := len(line) - 1
	for readPos := len(line) - 1; readPos >= 0; readPos-- {
		if line[readPos] != '.' {
			assert(line[readPos] == 'R' || line[readPos] == 'B')
			line[writePos] = line[readPos]
			writePos--
		}
	}
	for ; writePos >= 0; writePos-- {
		line[writePos] = '.'
	}
}

func winner(lines []string, toWin int) (reply string) {
	if false {
		for row, line := range lines {
			fmt.Println(line, row)
		}
	}
	redWin := findWin(lines, toWin, 'R')
	blueWin := findWin(lines, toWin, 'B')

	switch {
	case redWin && blueWin:
		reply = "Both"
	case redWin:
		reply = "Red"
	case blueWin:
		reply = "Blue"
	default:
		reply = "Neither"
	}
	return reply
}

func findWin(lines []string, toWin int, c byte) bool {
	//fmt.Println("findWin", string(c))

	board := len(lines)
	sequenceChan := make(chan sequence)
	go listSequences(board, sequenceChan)
	for s := range sequenceChan {
		//fmt.Println("sequence: ", s)
		found := 0

		for p := s.start; p.onBoard(board); p = s.next(p) {
			//fmt.Println("p: ", p, found, string(lines[p.row][p.col]))
			if lines[p.row][p.col] == c {
				found++
				if found == toWin {
					return true
				}
			} else {
				found = 0
			}
		}
	}
	return false
}

type point struct {
	row, col int
}

type sequence struct {
	name  string
	start point
	next  func(point) point
}

func (s sequence) String() string {
	return fmt.Sprintf("%s %v", s.name, s.start)
}

func (p point) onBoard(board int) bool {
	return p.row >= 0 && p.row < board && p.col >= 0 && p.col < board
}

func listSequences(board int, c chan sequence) {
	// Right
	next := func(p point) point {
		return point{p.row, p.col + 1}
	}
	for j := 0; j < board; j++ {
		c <- sequence{"right", point{j, 0}, next}
	}

	// Down
	next = func(p point) point {
		return point{p.row + 1, p.col}
	}
	for j := 0; j < board; j++ {
		c <- sequence{"down", point{0, j}, next}
	}

	// Bottom left to top right
	next = func(p point) point {
		return point{p.row - 1, p.col + 1}
	}

	for j := 0; j < board; j++ {
		c <- sequence{"updiag", point{board - 1 - j, 0}, next}
		c <- sequence{"updiag", point{board - 1, j}, next}
	}

	// Top left to bottom right
	next = func(p point) point {
		return point{p.row + 1, p.col + 1}
	}

	for j := 0; j < board; j++ {
		c <- sequence{"downdiag", point{0, j}, next}
		c <- sequence{"downdiag", point{j, 0}, next}
	}
	close(c)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
