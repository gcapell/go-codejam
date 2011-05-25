/* See http://code.google.com/codejam/contest/dashboard?c=90101#s=p1
 */

package main

import (
	"fmt"
	"strings"
	"codejam/ProblemReader"
)

type Cell struct {
	altitude int
	row, col	int

	tributaries [4]*Cell
	nTributaries int
	root *Cell
	name string
}

func (c *Cell) String() string {
	if c.root != nil {
		return fmt.Sprintf("%d(%d,%d)->(%d,%d)", c.altitude, c.row, c.col, c.root.row, c.root.col)
	} 
	return fmt.Sprintf("%d(%d,%d)", c.altitude, c.row, c.col)
}

func (c *Cell) addTributary (t *Cell) {
	c.tributaries[c.nTributaries] = t
	c.nTributaries++
}

func (c *Cell) markSink() {
	c.root = c
}

func (c *Cell) markTributaries() {
	for _, t := range c.tributaries[:c.nTributaries] {
		t.root = c.root
		t.markTributaries()
	}
}

type board struct {
	width, height int
	cell	[][]*Cell
}

func (b *board) String() string {
	lines := make([]string, b.height)
	for j :=0; j<b.height; j++ {
		lines[j] = fmt.Sprintf("%v", b.cell[j])
	}
	return strings.Join(lines, "\n")
}

func (b *board) display() string  {
	letters := "abcdefghijklmnopqrstuvwxyz"
	pos := 0

	reply := "\n"
	for j :=0; j<b.height; j++ {
		for k :=0; k<b.width; k++ {
			c := b.cell[j][k]
			if len(c.root.name) != 1 {
				c.root.name = string(letters[pos])
				pos++
			}
			reply += c.root.name + " "
		}
		reply += "\n"
	}
	return reply
}


type cellAndNeighbours struct {
	cell *Cell
	neighbours chan *Cell
}

func (b *board) cellsAndNeighbours() (chan cellAndNeighbours) {
	c := make(chan cellAndNeighbours)
	go func(){
		for j :=0; j < b.height; j++ {
			for k :=0; k < b.width; k++ {
				cell := b.cell[j][k]

				neighbours := make(chan *Cell)
				go func(j, k int){
					if j-1 >= 0 {
						neighbours <- b.cell[j-1][k]
					}
					if k-1 >= 0 {
						neighbours <- b.cell[j][k-1]
					}
					if k+1 <b.width {
						neighbours <- b.cell[j][k+1]
					}
					if j+1 <b.height {
						neighbours <- b.cell[j+1][k]
					}
					close(neighbours)
				}(j, k)
				c <- cellAndNeighbours{cell, neighbours}
			}
		}
		close(c)
	}()
	return c
}

func loadBoard(in *ProblemReader.ProblemReader) *board {
	hw := in.NNums(2)
	b := new(board)
	b.height, b.width = hw[0], hw[1]
	b.cell = make([][]*Cell, b.height)
	for j :=0; j< b.height; j++ {
		b.cell[j] = make([]*Cell, b.width)
		altitude := in.NNums(b.width)
		for k := 0 ; k < b.width; k++ {
			b.cell[j][k] = &Cell{altitude:altitude[k], row:j, col:k}
		}
	}
	return b
}
func solver(in *ProblemReader.ProblemReader) string {
	board := loadBoard(in)

	// fmt.Printf("board:\n%s\n", board)
	sinks := make ([]*Cell,0)

	for c := range board.cellsAndNeighbours() {
		lowest := c.cell
		
		for n := range c.neighbours {
			if n.altitude < lowest.altitude {
				lowest = n
			}
		}
		if lowest == c.cell {
			// fmt.Printf("Cell: %s sink\n", c.cell)
			c.cell.markSink()
			sinks = append(sinks, c.cell)
		} else {
			// fmt.Printf("Cell: %s -> %s\n", c.cell, lowest)
			lowest.addTributary(c.cell)
		}
	}
	for _, c := range sinks {
		c.markTributaries()
	}

	return board.display()
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
