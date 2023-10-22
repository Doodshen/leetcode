/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-03-15 19:51:48
 * @LastEditors: Andy
 * @LastEditTime: 2023-03-15 20:46:12
 */
/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	result := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return result.Next
}

// @lc code=end

