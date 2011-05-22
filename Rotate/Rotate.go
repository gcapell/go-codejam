/* See http://code.google.com/codejam/contest/dashboard?c=544101#s=p0
 */

package main

import (
	"fmt"
	"codejam/ProblemReader"
)

func solver(in *ProblemReader.ProblemReader) string {
	nums := in.Nums(2)
	board, toWin := nums[0], nums[1]
	lines := make([] string, board)
	for j := 0; j < board; j++ {
		line := []byte(in.Line())
		shiftLine(line)
		lines[j] = string(line)
	}
	return winner(lines, toWin)
}

func shiftLine(line []byte)  {
	writePos := len(line)-1
	for readPos := len(line)-1; readPos >=0; readPos-- {
		if line[readPos] != '.' {
			line[writePos] = line[readPos]
			writePos--
		}
	}
	for ;writePos >= 0; writePos-- {
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
	case redWin && blueWin: reply = "Both"
	case redWin: reply = "Red"
	case blueWin: reply = "Blue"
	default: reply = "Neither"
	}
	return reply
}

func findWin(lines []string, toWin int, c byte) bool {
	// fmt.Println("findWin", string(c))
	
	board := len(lines)
	sequenceChan := make(chan sequence)
	go listSequences(board, sequenceChan)
	for s := range sequenceChan {
		// fmt.Println("sequence: ", s)
		p := s.forward(s.start, -(toWin-1))
		found := 0
		required := toWin - found

		thisSeq:
		for p.onBoard(board) {
			// fmt.Println("p: ", p, string(lines[p.row][p.col]))
			switch lines[p.row][p.col] {
			case c:
				found++
				p = s.forward(p, 1)
			case '.':
				if s.earlyExit {
					break thisSeq
				}
				required, found = toWin, 0
				p = s.forward(p, -toWin)
			default:
				required = toWin - found
				found = 0
				p = s.forward(p, -toWin)
			}
			if found == required {
				return true
			}
		}
	}
	return false
}

type point struct {
	row, col int
}

type sequence struct {
	name string
	start point
	earlyExit bool
	forward func(point, int) point
}

func(s sequence) String() string {
	return fmt.Sprintf("%s %v", s.name, s.start)
}

func(p point) onBoard(board int) bool {
	return p.row >= 0 && p.row < board && p.col >= 0 && p.col < board
}

func listSequences(board int, c chan sequence) {
	// Left
	forward := func(p point, n int) point {
		return point{p.row, p.col + n}
	}
	for j :=0; j<board; j++ {
		c <- sequence{ "left", point{board-1-j, board-1}, true, forward}
	}

	// Down
	forward = func(p point, n int) point {
		return point{p.row + n, p.col }
	}
	for j :=0; j<board; j++ {
		c <- sequence{ "down", point{board-1, j}, true, forward}
	}

	// Bottom left to top right
	forward = func(p point, n int) point {
		return point{p.row-n, p.col + n}
	}

	for j :=0; j<board; j++ {
		c <- sequence{ "updiag",point{0, j}, false, forward}
		if j != 0 {
			c <- sequence{ "updiag", point{j, board-1}, false, forward}
		}
	}

	// Top left to bottom right
	forward = func(p point, n int) point {
		return point{p.row+n, p.col + n}
	}

	for j :=0; j<board; j++ {
		
		c <- sequence{ "downdiag", point{j, board-1}, true, forward}
		if j != board-1 {
			c <- sequence{ "downdiag", point{board-1, j}, true, forward}
		}
	}
	close(c)
}
	
func main() {
	ProblemReader.In.SolveProblems(solver)
}
