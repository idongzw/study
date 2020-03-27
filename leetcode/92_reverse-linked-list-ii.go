/*
 * @Author: dzw
 * @Date: 2020-03-16 21:27:17
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-17 22:22:08
 */

/*
 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
	showList(reverseBetween(l, 2, 5))
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n || head.Next == nil {
		return head
	}

	// 反转之前的节点
	preTmp := (*ListNode)(nil)
	// 反转开始的节点
	startTmp := head

	// 遍历到反转处
	pre := (*ListNode)(nil)
	cur := head
	for i := 0; i < m-1; i++ {
		preTmp = cur // 保存开始之前的节点
		cur = cur.Next
		startTmp = cur // 保存开始的节点
	}

	// 反转
	pre = cur
	cur = pre.Next
	for i := 0; i < n-m; i++ {
		tmp := cur.Next // 保存下一个节点
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	// 链接剩下的节点
	startTmp.Next = cur
	// 从头开始反转，直接返回
	if m == 1 {
		return pre
	}
	// 链接到之前的未反转节点
	preTmp.Next = pre
	return head
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Print("\n")
}
