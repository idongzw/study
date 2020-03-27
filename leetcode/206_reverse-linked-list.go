/*
 * @Author: dzw
 * @Date: 2020-03-16 21:06:26
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-16 21:20:56
 */

/*
 反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	showList(l)

	rst := reverseList(l)
	showList(rst)
}

func reverseList(head *ListNode) *ListNode {
	pre := (*ListNode)(nil)
	cur := head

	for cur != nil {
		tmp := cur.Next // 保存下个节点
		cur.Next = pre  // 指针反指
		pre = cur       // 更新前节点
		cur = tmp       // 更新当前节点
	}

	return pre
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Print("\n")
}
