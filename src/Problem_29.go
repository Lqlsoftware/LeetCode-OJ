/*
Divide Two Integers

Divide two integers without using multiplication, division and mod operator.
If it is overflow, return MAX_INT.
*/

package main

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0{
		return 0
	} else if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	res := 0
	flag := (uint32(dividend & 0x80000000) ^ uint32(divisor & 0x80000000)) >> 31
	z := 31
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for z > 0 && (divisor & (1 << uint32(z))) == 0 {
		z--
	}
	for i := 31 - z;i >= 0;i-- {
		x := divisor << uint32(i)
		if dividend >= x {
			dividend -= x
			res += 1 << uint32(i)
		}
	}

	if flag == 1 {
		return -res
	}
	return res
}