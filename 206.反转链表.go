/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(cur *ListNode) *ListNode {
	if cur == nil {
		return nil
	}
	var pre *ListNode
	for cur!=nil {
		temp:=cur.Next

		cur.Next=pre

		pre=cur
		
		cur=temp
		
	}
	return pre
}
	
// @lc code=end

