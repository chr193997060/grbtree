package tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/chr193997060/grbtree"
)

func findMinMax(t []int) ([2]int, error) {
	if len(t) == 0 {
		return [2]int{}, errors.New("t no value")
	}
	max := t[0]
	min := t[0]
	if len(t) != 1{
		for _, i := range(t[1:]){
			if i > max {
				max = i
			}
			if i < min {
				min = i
			}
		}
	}
	return [2]int{min, max}, nil
}

func treeAddTestKs(t *grbtree.RBTree, ks []int){
	for _, k := range(ks){
		t.Add(k, 1)
	}
}

func addCase1(t *testing.T, tree *grbtree.RBTree){
	addks := []int{55, 38, 80, 25, 46, 76, 72,  }
	treeAddTestKs(tree, addks)
	if int(tree.Len) != len(addks) {
		t.Fail()
	}
	min_max, _ := findMinMax(addks)
	min := min_max[0]
	max := min_max[1]
	t_max, _, _ := tree.GetMax()
	if t_max != grbtree.RBTreeKey(max){
		fmt.Println("max error")
		t.Fail()
	}
	t_min, _, _ := tree.GetMin()
	if t_min != grbtree.RBTreeKey(min){
		fmt.Println("min error")
		fmt.Println(t_min)
		t.Fail()
	}
	tree.PrintTree(5)
}

func TestAdd(t *testing.T) {
	rbt := grbtree.NewRBTree()
	addCase1(t, rbt)
}

func delCase1(t *testing.T, tree *grbtree.RBTree){
	tree.Clear()
	tree.Add(55, 1)
	tree.PrintTree(3)
	tree.Del(55)
	tree.PrintTree(3)
}

func delCase2(t *testing.T, tree *grbtree.RBTree){
	tree.Clear()

	treeAddTestKs(tree, []int{55, 38, 80, 25, 46, 76, 72,  })

	tree.PrintTree(5)

	tree.Del(80)
	tree.PrintTree(5)

	tree.Del(72)
	tree.PrintTree(5)

	tree.Del(76)
	tree.PrintTree(5)
}


func TestDel(t *testing.T){
	rbt := grbtree.NewRBTree()
	delCase1(t, rbt)
	delCase2(t, rbt)
}
