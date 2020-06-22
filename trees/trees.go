package trees

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

//PrintNode print value
func (n *Node) PrintNode() {
	if n == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(n.Value)
}

//PreOrder 前序遍历 父节点->左子树->右子树
func (n  *Node)PreOrder() {
	if n == nil {
		return
	}
	n.PrintNode()
	n.Left.PreOrder()
	n.Right.PreOrder()
}

//InOrder 中序遍历 左子树->父节点->右子树
func (n *Node)InOrder() {
	if (n == nil) {
		return
	}
	n.Left.InOrder()
	n.PrintNode()
	n.Right.InOrder()
}

//PostOrder 后序遍历 左子树 -> 右子树 ->父节点
func (n *Node)PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	n.PrintNode()
}
//深度优先DFS 栈先进后出 前序中序后序

//层次遍历 广度优先 bfs  => 队列 先进先出
func (n *Node)LevelOrder() {
	if n == nil {
		return
	}

	var result  []int
	queue := []*Node{n}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		result = append(result, n.Value)
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}

	for _,v:= range result {
		fmt.Println(v)
	}
}

// 反转二叉树
func (n *Node) Invert() {
	if n  == nil {
		return
	}
	tmp := n.Left
	n.Left = n.Right
	n.Right = tmp
	n.Left.Invert()
	n.Right.Invert()
}