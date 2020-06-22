package trees

import (
	"testing"
)

func getNode() *Node {
	node := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
				Left: &Node{
					8,
					nil,
					nil,
				},
				Right: &Node{
					9,
					nil,
					nil,
				},
			},
			Right: &Node{
				Value: 5,
				Left: &Node{
					10,
					nil,
					nil,
				},
				Right: nil,
			},
		},
		Right: &Node{
			Value: 3,
			Left: &Node{
				6,
				nil,
				nil,
			},
			Right: &Node{
				7,
				nil,
				nil,
			},
		},
	}

	return node
}

func Test_Node(t *testing.T) {
	node := getNode()
	//node.PreOrder()
	//node.InOrder()
	//node.PostOrder()
	//node.LevelOrder()
	node.Invert()
	node.LevelOrder()
}