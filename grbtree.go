// Package grbtree is a Red-black tree implemented by go
// grbtree 是使用go实现的红黑树
//
package grbtree

import (
	"fmt"
	"strconv"
)

const (
	RED   bool = true
	BLACK bool = false
)

type RBTreeKey int64

type RBTreeNode struct {
	Key    RBTreeKey
	Value  interface{}
	Color  bool
	parent *RBTreeNode
	left   *RBTreeNode
	right  *RBTreeNode
}

type RBTree struct {
	Root    *RBTreeNode
	Len     uint32
	minNode *RBTreeNode
	maxNode *RBTreeNode
}


func (k *RBTreeKey)StrLen() int {
	return len(strconv.FormatInt(int64(*k), 10))
}

func (k *RBTreeKey)ToStr() string {
	return strconv.FormatInt(int64(*k), 10)
}


func NewRBTreeNode(key int, val interface{}) *RBTreeNode {
	return &RBTreeNode{
		Key:   RBTreeKey(key),
		Value: val,
		Color: RED,
	}
}


func (n *RBTreeNode) GetParent() *RBTreeNode {
	return n.parent
}


func (n *RBTreeNode) GetLeft() *RBTreeNode {
	return n.left
}


func (n *RBTreeNode) GetRight() *RBTreeNode {
	return n.right
}


// 节点是否是黑色
func (n *RBTreeNode) isBlack() bool {
	if n == nil {
		return true
	} else {
		return !n.Color
	}
}


// 替换子节点
func (n *RBTreeNode) replaceChild(old *RBTreeNode, new *RBTreeNode){
	if n.left == old {
		n.left = new
	}else{
		n.right = new
	}
	if new != nil {
		new.parent = n
	}
}


// 查找兄弟节点
func (n *RBTreeNode) findBroNode() (bro *RBTreeNode) {
	if n.parent == nil {
		return nil
	}
	if n.parent.left == n {
		bro = n.parent.right
	} else {
		bro = n.parent.left
	}
	return bro
}


// 以传入节点进行左旋转
func (t *RBTree) leftRotate(n *RBTreeNode) {
	//    5                     9
	//  /   \     左旋转       /  \
	// 3     9   --------->   5   11
	//      / \              /  \
	//     7  11            3    7
	retNode := n.right
	n.right = retNode.left

	retNode.left = n
	retNode.parent = n.parent
	n.parent = retNode
	
	if retNode.parent == nil {
		t.Root = retNode
	}else if retNode.parent.left == n {
		retNode.parent.left = retNode
	}else{
		retNode.parent.right = retNode
	}
}


// 以传入节点进行右旋转
func (t *RBTree) rightRotate(n *RBTreeNode) {
	//      9                     5
	//    /   \     右旋转       /  \
	//   5    11   ------->     3    9
	//  / \                        /  \
	// 3   7                      7    11
	retNode := n.left
	n.left = retNode.right

	retNode.right = n
	retNode.parent = n.parent
	n.parent = retNode

	if retNode.parent == nil {
		t.Root = retNode
	}else if retNode.parent.left == n {
		retNode.parent.left = retNode
	}else{
		retNode.parent.right = retNode
	}
}


// 查找k对应节点和k最接近节点，k 不存在则返回第一个返回值为 nil
func (t *RBTree) findNodeAndRecentNode(k RBTreeKey) (*RBTreeNode, *RBTreeNode) {
	var recent *RBTreeNode
	var fnode *RBTreeNode
	fnode = t.Root
	for fnode != nil {
		recent = fnode
		if k < fnode.Key {
			fnode = fnode.left
		} else if k > fnode.Key {
			fnode = fnode.right
		} else {
			recent = fnode.parent
			break
		}
	}
	return fnode, recent
}


// 添加节点后的调整
func (t *RBTree) insertFixUp(n *RBTreeNode) {
	for !n.parent.isBlack() {
		uncleanNode := n.parent.findBroNode()
		if !uncleanNode.isBlack() {
			// 插入的节点的父节点和叔叔节点为红色，则：
			// 1）把父节点和叔叔节点设为黑色；2）把爷爷节点设为红色；
			// 3）把指针定位到爷爷节点作为当前需要操作的节点，再根据变换规则来进行判断操作
			n.parent.Color = BLACK
			uncleanNode.Color = BLACK
			n.parent.parent.Color = RED
			// 对爷爷节点进行调整操作
			n = n.parent.parent
		} else if n == n.parent.left {
			if n.parent == n.parent.parent.left{
				// LL 插入情况
				// LL情况：父节点为爷爷节点的左节点，插入节点为父节点的左节点
				// 1.把父节点变为黑色. 2.把爷爷节点变为红色. 3.以爷爷节点右旋转.
				n.parent.Color = BLACK
				n.parent.parent.Color = RED
				n = n.parent.parent
				t.rightRotate(n)
			}else{
				// RL 情况: 父节点为爷爷节点的右节点，插入节点为父节点的左节点
				// 1.插入节点变为黑色. 2.把爷爷节点变为红色. 3.以父节点右旋转. 4.以爷爷节点左旋转.
				n.Color = BLACK
				n.parent.parent.Color = RED
				t.rightRotate(n.parent)
				n = n.parent
				t.leftRotate(n)
			}
		} else {
			if n.parent == n.parent.parent.right {
				// RR情况：父节点为爷爷节点的右节点，插入节点为父节点的右节点
				// 1.把父节点变为黑色. 2.把爷爷节点变为红色. 3.以爷爷节点左旋转.
				n.parent.Color = BLACK
				n.parent.parent.Color = RED
				n = n.parent.parent
				t.leftRotate(n)
			}else{
				// LR情况：父节点为爷爷节点的左节点，插入节点为父节点的右节点
				// 1.插入节点变为黑色. 2.把爷爷节点变为红色. 3.以父节点左旋转. 4.以爷爷节点右旋转.
				n.Color = BLACK
				n.parent.parent.Color = RED
				t.leftRotate(n.parent)
				n = n.parent
				t.rightRotate(n)
			}
		}
	}
	t.Root.Color = BLACK
}


func (t *RBTree) insert(i_node *RBTreeNode) error {
	kn, nf := t.findNodeAndRecentNode(i_node.Key)
	if kn != nil {
		return errKeyAlreadyExists
	}
	if i_node.Key < nf.Key {
		nf.left = i_node
	} else {
		nf.right = i_node
	}
	i_node.parent = nf
	t.Len++
	if i_node.Key < t.minNode.Key {
		t.minNode = i_node
	}
	if i_node.Key > t.maxNode.Key {
		t.maxNode = i_node
	}
	t.insertFixUp(i_node)
	return nil
}


// 原先 n 父节点指向 n 的子节点替换为传入的节点
func (t *RBTree) parentReplaceChild(n *RBTreeNode, new *RBTreeNode){
	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = new
		}else{
			n.parent.right = new
		}
	}else{
		t.Root = new
	}
	if new != nil {
		new.parent = n.parent
	}
}


func (t *RBTree) delete(n *RBTreeNode) {
	if n.left == nil && n.right == nil {
		// 删除节点没有子节点，即为叶节点，
		if n.Color {
			// 叶子节点为红色直接将该节点删除（替换父节点指向nil）
			n.parent.replaceChild(n, nil)
		}else{
			if n.parent != nil {
				// 叶子节点为黑色，需要特殊处理
				t.deleteFixUp(n)
				if n.parent.left == n {
					n.parent.left = nil
				}else if n.parent.right == n{
					n.parent.right = nil
				}
			}else{
				// 根节点直接删除
				t.Root = nil
			}
		}
		if t.minNode == n {
			t.minNode = n.parent
		}
		if t.maxNode == n {
			t.maxNode = n.parent
		}
	}else if n.left == nil {
		// 删除的节点只有一个子节点
		// 1. 删除节点(父节点指向删除节点的子节点)，删除节点只能是黑色，子节点也只能是红色（删除节点只一个子节点，若删除节点子节点不是红色则到叶子节点的黑色节点数将不一样）
		// 2. 修改子节点为黑色
		t.parentReplaceChild(n, n.right)
		n.right.Color = BLACK
		if t.minNode.Key == n.Key {
			t.minNode = n.right
		}
	}else if n.right == nil {
		t.parentReplaceChild(n, n.left)
		n.left.Color = BLACK
		if t.maxNode.Key == n.Key {
			t.maxNode = n.left
		}
	}else {
		// 找到n的后继节点，把n替换成后继节点的k,值，转换为删除n的后继节点
		// n 有两个子节点，所以不可能是最大\最小节点
		var nextNode *RBTreeNode
		if n.right.left == nil {
			nextNode = n.right
		}else {
			nextNode = n.right.left
		}
		n.Key = nextNode.Key
		n.Value = nextNode.Value
		t.delete(nextNode)
		return
	}
	t.Len -= 1
}


// 删除节点的兄弟节点右红色子节点的情况下的颜色操作
func (t *RBTree) deleteNodeRedBroChildColorRevise(n *RBTreeNode) {
	c := n.parent.parent
	c.Color = n.parent.Color
	c.left.Color = BLACK
	c.right.Color = BLACK
}


// 黑色叶子节点删除后的调整操作
func (t *RBTree) deleteFixUp(n *RBTreeNode) {
	if n.parent == nil {
		return
	}
	if n.parent.left == n {
		broNode := n.findBroNode()
		if broNode.isBlack(){
			// 兄弟节点为黑色
			blColorIsBlack := broNode.left.isBlack()
			brColorIsBlack := broNode.right.isBlack()
			if !blColorIsBlack && !brColorIsBlack {
				// 兄弟节点有2个红色子节点
				// RR 
				t.leftRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else if !blColorIsBlack {
				// 兄弟节点左子节点为红色
				// RL
				// 先调整成RR模式，然后走RR模式的操作
				t.rightRotate(broNode)
				t.leftRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else if !brColorIsBlack {
				// 兄弟节点右子节点为红色
				// 删除节点后，经过删除节点的子节点的黑色路径会减1，就需要补充黑色节点，可以把兄弟节点的红色节点移到删除节点这边并修改颜色;
				// 旋转父节点相当于补充节点，然后把旋转后的父节点的父节点的颜色修改成父节点的颜色，就相当于把兄弟节点的移到了删除节点，
				// 这样就相当于补充上了删除节点，让后在把开始时兄弟节点的红色节点变成黑色就完成了调整
				// RR
				t.leftRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else{
				// 兄弟节点没有红色子节点
				broNode.Color = RED
				if !n.parent.isBlack(){
					// 父节点是红色的，把父节点颜色替换成黑色
					n.parent.Color = BLACK
				}else{
					// 父节点不是红色，则需要以父节点为删除节点进行调整
					t.deleteFixUp(n.parent)
				}
			}
		}else{
			// 兄弟节点为红色
			// 以父节点旋转，父节点和兄弟节点替换颜色，这样父节点就变成和黑色，变成了上面的情况
			// RR
			t.leftRotate(n.parent)
			n.parent.Color = RED
			broNode.Color = BLACK
			t.deleteFixUp(n)
		}
	}else {
		broNode := n.findBroNode()
		if broNode.isBlack() {
			blColorIsBlack := broNode.left.isBlack()
			brColorIsBlack := broNode.right.isBlack()
			if !blColorIsBlack && !brColorIsBlack {
				// LL
				t.rightRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else if !blColorIsBlack {
				// LL
				t.rightRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else if !brColorIsBlack {
				// LR
				t.leftRotate(broNode)
				t.rightRotate(n.parent)
				t.deleteNodeRedBroChildColorRevise(n)
			}else{
				broNode.Color = RED
				if !n.parent.isBlack() {
					n.parent.Color = BLACK
				}else{
					t.deleteFixUp(n.parent)
				}
			}
		}else{
			// LL
			t.rightRotate(n.parent)
			n.parent.Color = RED
			broNode.Color = BLACK
			t.deleteFixUp(n)
		}
	}
}

// tree := grbtree.NewRBTree()
func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) Get(k int) (interface{}, error) {
	n, _ := t.findNodeAndRecentNode(RBTreeKey(k))
	if n == nil {
		return nil, errKeyNotExists
	}
	return n.Value, nil
}

func (t *RBTree) GetMax() (k RBTreeKey, v interface {}, err error){
	if t.Root == nil {
		return k, nil, errNotNode
	}
	k = t.maxNode.Key
	v = t.maxNode.Value
	return k, v, err
}

func (t *RBTree) GetMin() (k RBTreeKey, v interface {}, err error){
	if t.Root == nil {
		return k, nil, errNotNode
	}
	k = t.minNode.Key
	v = t.minNode.Value
	return k, v, err
}

// 添加节点到树中
func (t *RBTree) Add(k int, v interface{}) {
	if t.Root == nil {
		t.Root = &RBTreeNode{
			Key:   RBTreeKey(k),
			Value: v,
			Color: BLACK,
		}
		t.minNode = t.Root
		t.maxNode = t.Root
		t.Len = 1
		return
	}
	node := NewRBTreeNode(k, v)
	t.insert(node)
}

// 删除树中的节点
func (t *RBTree) Del(k int) {
	if t.Root == nil {
		return
	}
	n, _ := t.findNodeAndRecentNode(RBTreeKey(k))
	if n == nil {
		return
	}
	t.delete(n)
}


// 清除树的节点
func (t *RBTree) Clear() {
	if t.Root == nil {
		return
	}else{
		t.Root = nil
		t.minNode = nil
		t.maxNode = nil
		t.Len = 0
	}
}

// 用于处理树多个nil节点邻近情况
type nodeBox struct {
	n *RBTreeNode
	c int // nil邻近节点的数量
}


// 广度查找节点. 多个邻近 nil 节点将合并成一个，通过计数来表示有多少nil节点
func (t *RBTree) bfs(layer int) [][]*nodeBox {
	if layer == 0 {
		return [][]*nodeBox{}
	}
	if t.Root == nil {
		return [][]*nodeBox{ {{c:1}, }}
	}
	queue := make([][]*nodeBox, 0)
	queue = append(queue, []*nodeBox{ {n: t.Root}})
	for i := 0; i < layer - 1; i++ {
		q := queue[i]
		next := make([]*nodeBox, 0)
		for _, nBox := range(q){
			if nBox.n == nil {
				for cc := 0; cc < nBox.c; cc++{
					next_len := len(next)
					if next_len > 0 && next[next_len-1].n == nil {
						next[next_len-1].c += 2
					}else{
						next = append(next, &nodeBox{c:2})
					}
				} 
			}else{
				if nBox.n.left != nil{
					next = append(next, &nodeBox{n: nBox.n.left})
				}else{
					next_len := len(next)
					if next_len > 0 && next[next_len-1].n == nil {
						next[next_len-1].c++
					}else{
						next = append(next, &nodeBox{c:1})
					}
				}
				if nBox.n.right != nil{
					next = append(next, &nodeBox{n: nBox.n.right})
				}else{
					next_len := len(next)
					if next_len > 0 && next[next_len-1].n == nil {
						next[next_len-1].c++
					}else{
						next = append(next, &nodeBox{c:1})
					}
				}
			}
		}
		if len(next) == 1 && next[0].n == nil {
			break
		}
		queue = append(queue, next)
	}
	return queue
}


// 显示树结构，()包裹的节点代表红色，[]代表黑色
//             [0x6]
//        /------|------\
//     (0x4)           (0x8)
//    /--|--\         /--|--\
// [0x1]   [0x2]   [0x7]   [0x9]
func (t *RBTree) PrintTree(layer int){
	if t.Len == 0 {
		fmt.Println(nil)
		return
	}else if t.Len == 1 || layer == 1 {
		fmt.Printf("[%v]\n", Int64ToHexStr(int64(t.Root.Key)))
		return
	}
	var node_width int
	queue := t.bfs(layer)
	layer = len(queue) - 1
	to16 := false
	downLayerMaxNodeCount := 1 << layer // 最底层的节点数量
	node_additional_width := len("[]")  // 节点额外信息宽度
	if to16 {
		node_width = node_additional_width + len(Int64ToHexStr(int64(t.maxNode.Key)))// 显示的节点总宽度
	}else{
		node_width = node_additional_width + t.maxNode.Key.StrLen() // 显示的节点总宽度
	}
	node_width_half := node_width >> 1
	down_node_interval := node_width - 2 // 最底层节点间的间隔, 即去掉()或[]后的节点宽度

	q_len := len(queue)
	_total := StrCopy("-", downLayerMaxNodeCount * node_width + (downLayerMaxNodeCount - 1) * down_node_interval)
	// 开始显示
	fmt.Println(_total)
	for layer, nBoxs := range(queue) {
		layer_node_interval_down_node_count := downLayerMaxNodeCount / (1 << layer)  // 本层 两个节点间 最下层节点数
		node_interval := (layer_node_interval_down_node_count * node_width + (layer_node_interval_down_node_count - 1) * down_node_interval) - 2  // 本层两个节点间的间隔
		a_node_down_node_count := layer_node_interval_down_node_count >> 1  // 开头到第一个节点间 最下层的节点数
		a_node_point := a_node_down_node_count * node_width  // 第一个节点的坐标
		if a_node_down_node_count > 1 {
			a_node_point += (a_node_down_node_count - 1) * down_node_interval
		}
		a_space_count := a_node_point - 1  // 开始位置到第一个节点的空格数

		_count := int((a_node_point) / 2)  // - 显示的数量

		print_node_count := 0  // 本层已经输出的节点数

		var nodeK string
		var s1 string
		var s2 string

		a_space_count_space := StrCopy(" ", a_space_count)
		node_interval_space := StrCopy(" ", node_interval)
		_count_str := StrCopy("-", _count)
		if len(nBoxs) == 1 && nBoxs[0].n == nil {
			break
		}
		for _, nBox := range(nBoxs) {
			if nBox.n != nil {
				var k string
				if to16 {
					k = Int64ToHexStr(int64(nBox.n.Key))
					k = StrRightFilling(k, down_node_interval, " ")
				}else{
					k = nBox.n.Key.ToStr()
					k = StrLeftFilling(k, down_node_interval, "0")
				}
				if nBox.n.isBlack() {
					nodeK = fmt.Sprintf("[%v]", k)  // 节点将以16进制显示
				}else{
					nodeK = fmt.Sprintf("(%v)", k)
				}
				if print_node_count == 0 {
					s1 += fmt.Sprintf("%v%v", a_space_count_space, nodeK)
				}else{
					s1 += fmt.Sprintf("%v%v", node_interval_space, nodeK)
				}
				if layer < q_len - 1 {
					a := len(s1) - node_width_half - _count - 1 - len(s2) - (node_width % 2)
					if nBox.n.left != nil {
						s2 += fmt.Sprintf("%v/%v", StrCopy(" ", a), _count_str)
					}else{
						s2 += StrCopy(" ", a + _count + 1)
					}
					if nBox.n.left != nil || nBox.n.right != nil{
						s2 += "|"
					}
					if nBox.n.right != nil {
						s2 += fmt.Sprintf("%v\\", _count_str)
					}	
				}
				print_node_count ++
			}else if nBox.c > 0 {
				nodeK = StrCopy(" ", node_width)
				for cc := nBox.c; cc > 0; cc-- {
					if print_node_count == 0 {
						s1 += fmt.Sprintf("%v%v", a_space_count_space, nodeK)
					}else{
						s1 += fmt.Sprintf("%v%v", node_interval_space, nodeK)
					}
					print_node_count ++
				} 
			}
		}
		fmt.Print(s1, "\n")
		if s2 != ""{
			fmt.Print(s2, "\n")
		}
	}
	fmt.Println(_total)
}