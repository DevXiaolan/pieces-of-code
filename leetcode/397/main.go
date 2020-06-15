package main

import (
	"fmt"
)

func integerReplacement(n int) int {
	count:=0
	for n!=1 {
		if n&1==0 {
			n/=2
		}else{
			if (n & 2) == 0 || n == 3  {
				n--
			}else{
				n++
			}
		}
		count++
	}
	return count
}

func main() {
	fmt.Printf("%+v\n",[]interface{}{integerReplacement(65535)})
}

/**
这里主要是位运算
题目的目标是把n变成1，可以理解为，向着把数字变小的方向，是效率最高的。
所以，我们应该尽量制造机会让n连续除以2。
于是我们这里在决定加一还是减一的问题是，尽量让结果是4的倍数。
举个例子，5、13.
n=5的情况下，我们会选择n--, 因为4可以连续➗2.
同理，n=11时，采取n++, 12比10有更多的➗2的机会
*/