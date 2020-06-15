package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}

	var l1Value int
	var l2Value int

	var next1 *ListNode
	var next2 *ListNode

	next1 = nil
	next2 = nil
	hasAdd := false

	if l1 == nil {
		l1Value = 0
	} else {
		l1Value = l1.Val
		next1 = l1.Next
	}

	if l2 == nil {
		l2Value = 0
	} else {
		l2Value = l2.Val
		next2 = l2.Next
	}

	node.Val = l1Value + l2Value

	if next1 != nil || next2 != nil {
		if node.Val >= 10 {
			if next1 != nil {
				l1.Next.Val += 1
				hasAdd = true
			}

			if !hasAdd && next2 != nil {
				l2.Next.Val += 1
			}

			node.Val = node.Val - 10
		}

		node.Next = addTwoNumbers(next1, next2)
	} else {
		if node.Val >= 10 {
			node.Val = node.Val - 10
			node.Next = &ListNode{1, nil}
		}
	}

	return node
}

func main() {
	l1 := &ListNode{
		9,
		&ListNode{
			1,
			&ListNode{
				6,
				nil,
			},
		},
	}
	l2 := &ListNode{
		0,
		&ListNode{
			1, nil,
		},
	}
	result := addTwoNumbers(l1, l2)

	result.toString()
}

func (v *ListNode) toString() {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("%+v", v)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", " ")
	if err != nil {
		fmt.Printf("%+v", v)
	}
	fmt.Println(out.String())
}
