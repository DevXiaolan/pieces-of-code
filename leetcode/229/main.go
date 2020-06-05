package main

import (
	"fmt"
)

/**
* 注意了，题目里所有的信息都是有它的意义的
* 能不能想到妙赞的解法，看你是否有注意到 n/3 这个条件
* 这意味着，在最终结果里的个数，不会超过2个！
 */
func majorityElement(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 只申请两个位置就够了
	flag := []int{0, 0}
	count := []int{0, 0}
	// 抵消阶段
	for i := 0; i < l; i++ {
		if flag[0] == nums[i] {
			// 非空位，计数加1
			count[0]++
		} else if flag[1] == nums[i] {
			// 非空位，计数加1
			count[1]++
		} else if count[0] == 0 {
			// 空位
			flag[0] = nums[i]
			count[0] = 1
		} else if count[1] == 0 {
			// 空位
			flag[1] = nums[i]
			count[1] = 1
		} else {
			// 没自己的位置？所有人计数-1
			// 因为题目只需要你找出成员，不需要统计次数
			count[0]--
			count[1]--
		}
	}
	// 计数阶段
	count[0] = 0
	count[1] = 0
	rs := []int{}
	for i := 0; i < l; i++ {
		if nums[i] == flag[0] {
			if count[0] > l/3 {
				continue
			}
			count[0]++
			if count[0] > l/3 {
				rs = append(rs, flag[0])
			}
		} else if nums[i] == flag[1] {
			if count[1] > l/3 {
				continue
			}
			count[1]++
			if count[1] > l/3 {
				rs = append(rs, flag[1])
			}
		}
	}
	return rs
}

func main() {
	in := []int{1, 2, 3, 4}
	fmt.Printf("%+v\n", []interface{}{majorityElement(in)})
}
