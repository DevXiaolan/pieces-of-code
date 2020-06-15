package main

import (
	"fmt"
)

func rangeBitwiseAnd(m int, n int) int {
	var c uint
	for m!=n{
		m=m>>1
		n=n>>1
		c++
	}
	return m<<c
}

func main() {
	fmt.Printf("%+v\n",[]interface{}{rangeBitwiseAnd(4000000,2147483646)})
}

/**
对于递增数列，总是由低位往高位在0和1之间变化，
由于只要出现一个0，那这一位就一直会是0，这个是最优解的关键。
所以我们只需要找到m,n两个数的最大前缀，就是最终答案了。
比起迭代每一个数，快了xxxx个数量级。
*/