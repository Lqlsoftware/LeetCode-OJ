/*
addOrSubOrMul
请对于下面式子进行填空，填入加减乘，使这个表达式成立。
1 ? 2 ? 3 ? 4 ? 5 ? 6 ? 7 ? 8 ? 9 ? 10 == 0
请输出一共有多少种方案可以使得表达式成立。
*/

package main

func addOrSubOrMul() int {
	add := func(a, b int) int {return a + b}
	sub := func(a, b int) int {return a - b}
	operate := []func(a, b int) int {add,sub}
	count := 0
	var calculate [][]int
	// 列举乘法情况
	for i := 0;i < 512;i++ {
		num := []int{1}
		for j := uint(0);j < 9;j++ {
			if (i >> j) & 1 == 1 {
				num[len(num) - 1] *= int(j + 2)
			} else {
				num = append(num, int(j + 2))
			}
		}
		calculate = append(calculate, num)
	}
	// 列举加减运算结果
	for _,v := range calculate {
		length := uint(len(v) - 1)
		total := 1 << length
		for i := 0;i < total;i++ {
			result := v[0]
			for j := uint(0);j < length;j++ {
				result = operate[(i >> j) & 1](result,v[j + 1])
			}
			if result == 0 {
				count++
			}
		}
	}
	return count
}