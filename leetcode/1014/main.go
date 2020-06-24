package main

import (
	"fmt"
)

func maxScoreSightseeingPair(A []int) int {
	score, j, l, iMax := 0, 1, len(A), 0

	for j < l {
		if A[j-1]+j-1 > iMax {
			iMax = A[j-1] + j - 1
		}
		if iMax+A[j]-j > score {
			score = iMax + A[j] - j
		}
		j++
	}
	return score
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})})
}

/*
这一题，正向暴力解，会被直接判定超时。
正向暴力就是，先选的一个i，然后遍历i后面的j，并且暂存这个过程的最大值。
执行次数是 n的阶乘，算法上表达式 n^2 。

我们不妨换个思路，先定一个 j ，然后遍历 0 到 j-1，
这样区别在哪？我们不妨先看看得分公式
A[i] + A[j] + i - j
转换一下
(A[i] + i) + (A[j] - j)

当我们选定了 j， 那 A[j] - j 就是一个固定值，
那 score(i,j) 可以理解为：求最大的  A[i] + i 。
这一步思路转变很关键，成功想通了，才有下一步。

然后在我们求最大的  A[i] + i 这个过程里，每一次j的变化，
我们都可以复用上一个 j 时算出来的 iMax 结果，从而快速计算新的 iMax
这是第二个关键思路。


