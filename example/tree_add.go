package example

import (
	"fmt"

	"github.com/chr193997060/grbtree"
)

func main() {
	tree := grbtree.NewRBTree()
	tree.Add(1000, 1)
	tree.Add(2000, 2)
	v, err := tree.Get(1000)
	if err != nil {
		fmt.Println(err.Error())
	}
	if v.(int) != 1 {
		fmt.Println(v)
	}

	tree.Del(1000)
	_, err = tree.Get(1000)
	if err != nil {
		fmt.Println(err.Error())
	}
	tree.PrintTree(5)
}