/* See http://code.google.com/codejam/contest/dashboard?c=635101#s=p0
 */

package main

import (
	"fmt"
	"strings"
	"codejam/ProblemReader"
)

type Tree map[string]Tree

func newTree() Tree {
	return Tree(make(map[string]Tree))
}

/* Add path to tree, return number of components added */
func (tree Tree) add(path string) int {
	// fmt.Println("adding ", path, "to", tree)
	components := strings.Split(path[1:], "/", -1)
	additions := 0
	pos := tree
	for _, c := range components {
		child, ok := pos[c]
		if !ok {
			child = newTree()
			pos[c] = child
			additions++
		}
		pos = child
	}
	// fmt.Printf("add %s -> %#v : %d\n", path, components, additions)
	return additions
}

func solver(in *ProblemReader.ProblemReader) string {
	nums := in.NNums(2)
	existing, toCreate := nums[0], nums[1]

	tree := newTree()

	for j := 0; j < existing; j++ {
		tree.add(in.Line())
	}
	additions := 0
	for j := 0; j < toCreate; j++ {
		additions += tree.add(in.Line())
	}

	return fmt.Sprintf("%d", additions)
}

func main() {
	ProblemReader.In.SolveProblems(solver)
}
