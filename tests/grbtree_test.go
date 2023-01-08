package tests

import (
	"errors"
	"fmt"
	"testing"
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

func treeAddTestKs(t *RBTree, ks []int){
	for _, k := range(ks){
		t.Add(k, 1)
	}
}

func addCase1(t *testing.T, tree *RBTree){
	addks := []int{55, 38, 80, 25, 46, 76, 72,  }
	treeAddTestKs(tree, addks)
	if int(tree.Len) != len(addks) {
		t.Fail()
	}
	min_max, _ := findMinMax(addks)
	min := min_max[0]
	max := min_max[1]
	if tree.maxNode.Key != RBTreeKey(max){
		fmt.Println("max error")
		t.Fail()
	}
	if tree.minNode.Key != RBTreeKey(min){
		fmt.Println("min error")
		fmt.Println(tree.minNode.Key)
		t.Fail()
	}
	tree.BFSPrint()
}

func TestAdd(t *testing.T) {
	rbt := NewRBTree()
	addCase1(t, rbt)
}

func delCase1(t *testing.T, tree *RBTree){
	tree.Clear()
	tree.Add(55, 1)
	fmt.Println("rb len", tree.Len, "max:", tree.maxNode.Key, "min:", tree.minNode.Key)
	tree.BFSPrint()
	tree.Del(55)
	tree.BFSPrint()
}

func delCase2(t *testing.T, tree *RBTree){
	tree.Clear()

	treeAddTestKs(tree, []int{55, 38, 80, 25, 46, 76, 72,  })
	fmt.Println("rb len", tree.Len, "max:", tree.maxNode, "min:", tree.minNode)

	tree.BFSPrint()

	tree.Del(80)
	tree.BFSPrint()

	tree.Del(72)
	tree.BFSPrint()

	tree.Del(76)
	tree.BFSPrint()
}


func TestDel(t *testing.T){
	rbt := NewRBTree()
	delCase1(t, rbt)
	delCase2(t, rbt)
}
