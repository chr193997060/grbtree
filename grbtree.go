package grbtree

import "fmt"

const (
	RED   bool = true
	BLACK bool = false
)

type RBTreeKey int64

type RBTreeNode struct {
	Key    RBTreeKey
	Value  interface{}
	Color  bool
	Parent *RBTreeNode
	Left   *RBTreeNode
	Right  *RBTreeNode
}

type RBTree struct {
	Root    *RBTreeNode
	Len     uint32
	minNode *RBTreeNode
	maxNode *RBTreeNode
}


func NewRBTree() *RBTree {
	return &RBTree{}
}


func NewRBTreeNode(key int, val interface{}) *RBTreeNode {
	return &RBTreeNode{
		Key:   RBTreeKey(key),
		Value: val,
		Color: RED,
	}
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
	if n.Left == old {
		n.Left = new
	}else{
		n.Right = new
	}
	if new != nil {
		new.Parent = n
	}
}


// 以传入节点进行左旋转
//    5                     9
//  /   \     左旋转       /  \
// 3     9   --------->   5   11
//      / \              /  \
//     7  11            3    7
func (t *RBTree) LeftRotate(n *RBTreeNode) {
	retNode := n.Right
	n.Right = retNode.Left

	retNode.Left = n
	retNode.Parent = n.Parent
	n.Parent = retNode
	
	if retNode.Parent == nil {
		t.Root = retNode
	}else if retNode.Parent.Left == n {
		retNode.Parent.Left = retNode
	}else{
		retNode.Parent.Right = retNode
	}
}


// 以传入节点进行右旋转
//      9                     5
//    /   \     右旋转       /  \
//   5    11   ------->     3    9
//  / \                        /  \
// 3   7                      7    11
func (t *RBTree) RightRotate(n *RBTreeNode) {
	retNode := n.Left
	n.Left = retNode.Right

	retNode.Right = n
	retNode.Parent = n.Parent
	n.Parent = retNode

	if retNode.Parent == nil {
		t.Root = retNode
	}else if retNode.Parent.Left == n {
		retNode.Parent.Left = retNode
	}else{
		retNode.Parent.Right = retNode
	}
}


// 查找k对应节点和k最接近节点，k 不存在则返回第一个返回值为 nil
func (tree *RBTree) findNodeAndRecentNode(k RBTreeKey) (*RBTreeNode, *RBTreeNode) {
	var recent *RBTreeNode
	var fnode *RBTreeNode
	fnode = tree.Root
	for fnode != nil {
		recent = fnode
		if k < fnode.Key {
			fnode = fnode.Left
		} else if k > fnode.Key {
			fnode = fnode.Right
		} else {
			recent = fnode.Parent
			break
		}
	}
	return fnode, recent
}


// 查找兄弟节点
func (n *RBTreeNode) findBroNode() (bro *RBTreeNode) {
	if n.Parent == nil {
		return nil
	}
	if n.Parent.Left == n {
		bro = n.Parent.Right
	} else {
		bro = n.Parent.Left
	}
	return bro
}


// 添加节点后的调整
func (tree *RBTree) insertFixUp(n *RBTreeNode) {
	for !n.Parent.isBlack() {
		uncleanNode := n.Parent.findBroNode()
		if !uncleanNode.isBlack() {
			// 插入的节点的父节点和叔叔节点为红色，则：
			// 1）把父节点和叔叔节点设为黑色；2）把爷爷节点设为红色；
			// 3）把指针定位到爷爷节点作为当前需要操作的节点，再根据变换规则来进行判断操作
			n.Parent.Color = BLACK
			uncleanNode.Color = BLACK
			n.Parent.Parent.Color = RED
			// 对爷爷节点进行调整操作
			n = n.Parent.Parent
		} else if n == n.Parent.Left {
			if n.Parent == n.Parent.Parent.Left{
				// LL 插入情况
				// LL情况：父节点为爷爷节点的左节点，插入节点为父节点的左节点
				// 1.把父节点变为黑色. 2.把爷爷节点变为红色. 3.以爷爷节点右旋转.
				n.Parent.Color = BLACK
				n.Parent.Parent.Color = RED
				n = n.Parent.Parent
				tree.RightRotate(n)
			}else{
				// RL 情况: 父节点为爷爷节点的右节点，插入节点为父节点的左节点
				// 1.插入节点变为黑色. 2.把爷爷节点变为红色. 3.以父节点右旋转. 4.以爷爷节点左旋转.
				n.Color = BLACK
				n.Parent.Parent.Color = RED
				tree.RightRotate(n.Parent)
				n = n.Parent
				tree.LeftRotate(n)
			}
		} else {
			if n.Parent == n.Parent.Parent.Right {
				// RR情况：父节点为爷爷节点的右节点，插入节点为父节点的右节点
				// 1.把父节点变为黑色. 2.把爷爷节点变为红色. 3.以爷爷节点左旋转.
				n.Parent.Color = BLACK
				n.Parent.Parent.Color = RED
				n = n.Parent.Parent
				tree.LeftRotate(n)
			}else{
				// LR情况：父节点为爷爷节点的左节点，插入节点为父节点的右节点
				// 1.插入节点变为黑色. 2.把爷爷节点变为红色. 3.以父节点左旋转. 4.以爷爷节点右旋转.
				n.Color = BLACK
				n.Parent.Parent.Color = RED
				tree.LeftRotate(n.Parent)
				n = n.Parent
				tree.RightRotate(n)
			}
		}
	}
	tree.Root.Color = BLACK
}


func (tree *RBTree) insert(i_node *RBTreeNode) error {
	kn, nf := tree.findNodeAndRecentNode(i_node.Key)
	if kn != nil {
		return errKeyAlreadyExists
	}
	if i_node.Key < nf.Key {
		nf.Left = i_node
	} else {
		nf.Right = i_node
	}
	i_node.Parent = nf
	tree.Len++
	if i_node.Key < tree.minNode.Key {
		tree.minNode = i_node
	}
	if i_node.Key > tree.maxNode.Key {
		tree.maxNode = i_node
	}
	tree.insertFixUp(i_node)
	return nil
}


// 添加节点到树中
func (tree *RBTree) Add(k int, v interface{}) {
	if tree.Root == nil {
		tree.Root = &RBTreeNode{
			Key:   RBTreeKey(k),
			Value: v,
			Color: BLACK,
		}
		tree.minNode = tree.Root
		tree.maxNode = tree.Root
		tree.Len = 1
		return
	}
	node := NewRBTreeNode(k, v)
	tree.insert(node)
}


func (tree *RBTree) Get(k int) (interface{}, error) {
	n, _ := tree.findNodeAndRecentNode(RBTreeKey(k))
	if n == nil {
		return nil, errKeyNotExists
	}
	return n.Value, nil
}



// 原先 n 父节点指向 n 的子节点替换为传入的节点
func (tree *RBTree) parentReplaceChild(n *RBTreeNode, new *RBTreeNode){
	if n.Parent != nil {
		if n.Parent.Left == n {
			n.Parent.Left = new
		}else{
			n.Parent.Right = new
		}
	}else{
		tree.Root = new
	}
	if new != nil {
		new.Parent = n.Parent
	}
}


func (tree *RBTree) delete(n *RBTreeNode) {
	if n.Left == nil && n.Right == nil {
		// 删除节点没有子节点，即为叶节点，
		if n.Color {
			// 叶子节点为红色直接将该节点删除（替换父节点指向nil）
			n.Parent.replaceChild(n, nil)
		}else{
			if n.Parent != nil {
				// 叶子节点为黑色，需要特殊处理
				tree.deleteFixUp(n)
				if n.Parent.Left == n {
					n.Parent.Left = nil
				}else if n.Parent.Right == n{
					n.Parent.Right = nil
				}
			}else{
				// 根节点直接删除
				tree.Root = nil
			}
		}
		if tree.minNode == n {
			tree.minNode = n.Parent
		}
		if tree.maxNode == n {
			tree.maxNode = n.Parent
		}
	}else if n.Left == nil {
		// 删除的节点只有一个子节点
		// 1. 删除节点(父节点指向删除节点的子节点)，删除节点只能是黑色，子节点也只能是红色（删除节点只一个子节点，若删除节点子节点不是红色则到叶子节点的黑色节点数将不一样）
		// 2. 修改子节点为黑色
		tree.parentReplaceChild(n, n.Right)
		n.Right.Color = BLACK
		if tree.minNode.Key == n.Key {
			tree.minNode = n.Right
		}
	}else if n.Right == nil {
		tree.parentReplaceChild(n, n.Left)
		n.Left.Color = BLACK
		if tree.maxNode.Key == n.Key {
			tree.maxNode = n.Left
		}
	}else {
		// 找到n的后继节点，把n替换成后继节点的k,值，转换为删除n的后继节点
		// n 有两个子节点，所以不可能是最大\最小节点
		var nextNode *RBTreeNode
		if n.Right.Left == nil {
			nextNode = n.Right
		}else {
			nextNode = n.Right.Left
		}
		n.Key = nextNode.Key
		n.Value = nextNode.Value
		tree.delete(nextNode)
		return
	}
	tree.Len -= 1
}


// 删除节点的兄弟节点右红色子节点的情况下的颜色操作
func (tree *RBTree) deleteNodeRedBroChildColorRevise(n *RBTreeNode) {
	c := n.Parent.Parent
	c.Color = n.Parent.Color
	c.Left.Color = BLACK
	c.Right.Color = BLACK
}


// 黑色叶子节点删除后的调整操作
func (tree *RBTree) deleteFixUp(n *RBTreeNode) {
	if n.Parent == nil {
		return
	}
	if n.Parent.Left == n {
		broNode := n.findBroNode()
		if broNode.isBlack(){
			// 兄弟节点为黑色
			blColorIsBlack := broNode.Left.isBlack()
			brColorIsBlack := broNode.Right.isBlack()
			if !blColorIsBlack && !brColorIsBlack {
				// 兄弟节点有2个红色子节点
				// RR 
				tree.LeftRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else if !blColorIsBlack {
				// 兄弟节点左子节点为红色
				// RL
				// 先调整成RR模式，然后走RR模式的操作
				tree.RightRotate(broNode)
				tree.LeftRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else if !brColorIsBlack {
				// 兄弟节点右子节点为红色
				// 删除节点后，经过删除节点的子节点的黑色路径会减1，就需要补充黑色节点，可以把兄弟节点的红色节点移到删除节点这边并修改颜色;
				// 旋转父节点相当于补充节点，然后把旋转后的父节点的父节点的颜色修改成父节点的颜色，就相当于把兄弟节点的移到了删除节点，
				// 这样就相当于补充上了删除节点，让后在把开始时兄弟节点的红色节点变成黑色就完成了调整
				// RR
				tree.LeftRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else{
				// 兄弟节点没有红色子节点
				broNode.Color = RED
				if !n.Parent.isBlack(){
					// 父节点是红色的，把父节点颜色替换成黑色
					n.Parent.Color = BLACK
				}else{
					// 父节点不是红色，则需要以父节点为删除节点进行调整
					tree.deleteFixUp(n.Parent)
				}
			}
		}else{
			// 兄弟节点为红色
			// 以父节点旋转，父节点和兄弟节点替换颜色，这样父节点就变成和黑色，变成了上面的情况
			// RR
			tree.LeftRotate(n.Parent)
			n.Parent.Color = RED
			broNode.Color = BLACK
			tree.deleteFixUp(n)
		}
	}else {
		broNode := n.findBroNode()
		if broNode.isBlack() {
			blColorIsBlack := broNode.Left.isBlack()
			brColorIsBlack := broNode.Right.isBlack()
			if !blColorIsBlack && !brColorIsBlack {
				// LL
				tree.RightRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else if !blColorIsBlack {
				// LL
				tree.RightRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else if !brColorIsBlack {
				// LR
				tree.LeftRotate(broNode)
				tree.RightRotate(n.Parent)
				tree.deleteNodeRedBroChildColorRevise(n)
			}else{
				broNode.Color = RED
				if !n.Parent.isBlack() {
					n.Parent.Color = BLACK
				}else{
					tree.deleteFixUp(n.Parent)
				}
			}
		}else{
			// LL
			tree.RightRotate(n.Parent)
			n.Parent.Color = RED
			broNode.Color = BLACK
			tree.deleteFixUp(n)
		}
	}
}


// 删除树中的节点
func (tree *RBTree) Del(k int) {
	if tree.Root == nil {
		return
	}
	n, _ := tree.findNodeAndRecentNode(RBTreeKey(k))
	if n == nil {
		return
	}
	tree.delete(n)
}


// 清除树的节点
func (tree *RBTree) Clear() {
	if tree.Root == nil {
		return
	}else{
		tree.Root = nil
		tree.minNode = nil
		tree.maxNode = nil
		tree.Len = 0
	}
}


// 广度输出
func (tree *RBTree) BFSPrint() {
	fmt.Println("-----------------")
	fmt.Println(tree.Len)
	if tree.Root == nil {
		fmt.Println(nil)
		return
	}
	queue := make([][]*RBTreeNode, 0)
	queue = append(queue, []*RBTreeNode{tree.Root})
	for i := 0; i <int(tree.Len); i++ {
		q := queue[i]
		nextQ := make([]*RBTreeNode, 0)
		for _, n := range(q){
			fmt.Print(n.Key, n.Color, " ")
			if n.Left != nil{
				nextQ = append(nextQ, n.Left)
			}
			if n.Right != nil {
				nextQ = append(nextQ, n.Right)
			}
		}
		fmt.Print("\n")
		if len(nextQ) == 0 {
			break
		}
		queue = append(queue, nextQ)
	}
}