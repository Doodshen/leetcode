/*
 * @Descripttion:
 * @version:
 * @Author: wzs
 * @Date: 2023-10-14 12:22:42
 * @LastEditors: Andy
 * @LastEditTime: 2023-10-14 12:32:28
 */
/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] { //找到最小值
			min = prices[i]
		} else if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}
	return maxProfit
}

// @lc code=end

