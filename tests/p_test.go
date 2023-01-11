package tests

import (
	"testing"

	"github.com/chr193997060/grbtree"
)


func TestXxx(t *testing.T) {
	tree := grbtree.NewRBTree()
	// tree.Add(6, 1)
	// tree.Add(4, 1)
	// tree.Add(8, 1)
	// tree.Add(1, 1)
	// tree.Add(2, 1)
	// tree.Add(7, 1)
	// tree.Add(9, 1)
	// tree.Add(10, 1)
	// tree.Add(3, 1)
	tree.PrintTree(3)

	for i := 1; i < 20; i++ {
		tree.Add(i, i)
	}

	tree.PrintTree(7)
}


