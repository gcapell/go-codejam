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
	return winner(lines, board, toWin)
}

func winner(lines []string, board int, toWin int) (reply string) {
	fmt.Println("board:", board, "toWin:", toWin)
	for row, line := range lines {
		fmt.Println(line, row)
	}
	redWin := findWin(lines, board, toWin, 'R')
	blueWin := findWin(lines, board, toWin, 'B')
	
	switch {
	case redWin && blueWin: reply = "Both"
	case redWin: reply = "Red"
	case blueWin: reply = "Blue"
	default: reply = "Neither"
	}
	return reply
}

func findWin(lines []string, board int, toWin int, c byte) bool {
	fmt.Println("findWin", string(c))
	
	for row :=len(lines)-1; row >=0; row-- {
		first := len(lines)-toWin
		found := 0

		for first >=0 && lines[row][first]!='.' {
			required := toWin - found
			found = 0
			for pos := first; pos < first + required; pos++ {
				fmt.Println(row, pos, string(lines[row][pos]))
				if lines[row][pos] == c {
					found++
				} else {
					break
				}
			}
			if found == required {
				fmt.Println("success!")
				return true
			}
		}
	}

	// Vertical
	fmt.Println("vertical")
	for col :=len(lines)-1; col >=0; col-- {
		first := len(lines)-toWin

		vertical:
		for first >=0 && lines[first][col]!='.' {
			for pos := first; pos < first + toWin; pos++ {
				fmt.Println(pos, col, string(lines[pos][col]))
				if lines[pos][col] != c {
					first = pos - toWin
					continue vertical
				}
			}
			fmt.Println("success!")
			return true
		}
	}
	
	return false
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

func main() {
	ProblemReader.In.SolveProblems(solver)
}
