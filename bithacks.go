package go_bithacks

import (
	"math/bits"
)

// NextPowerOfTwoInt32 获取下一个大于等于一个int32的2的幂
func NextPowerOfTwoInt32(x int32) int32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}

// NextPowerOfTwoInt64 获取下一个大于等于一个int64的2的幂
func NextPowerOfTwoInt64(x int64) int64 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	x++
	return x
}

// DetectTwoIntegerOppositeSign 检查两个int32是否符号相反
func DetectTwoIntegerOppositeSign(x, y int32) bool {
	return (x ^ y) < 0
}

// DetectTwoIntegerSameSign 检查两个int32是否符号相同
func DetectTwoIntegerSameSign(x, y int32) bool {
	return (x ^ y) >= 0
}

// AbsoluteValueOfInt32 求一个int32的绝对值
func AbsoluteValueOfInt32(x int32) int32 {
	return x ^ (x >> 31) - (x >> 31)
}

// AbsoluteValueOfInt64 求一个int64的绝对值
func AbsoluteValueOfInt64(x int64) int64 {
	return x ^ (x >> 63) - (x >> 63)
}

// MinOfTwoInt32 找到两个int32中较小的那个
func MinOfTwoInt32(x, y int32) int32 {
	return y + ((x - y) & ((x - y) >> 31))
}

// MinOfTwoInt64 找到两个int64中较小的那个
func MinOfTwoInt64(x, y int64) int64 {
	return y + ((x - y) & ((x - y) >> 63))
}

// MaxOfTwoInt32 找到两个int32中较大的那个
func MaxOfTwoInt32(x, y int32) int32 {
	return x - ((x - y) & ((x - y) >> 31))
}

// MaxOfTwoInt64 找到两个int64中较大的那个
func MaxOfTwoInt64(x, y int64) int64 {
	return x - ((x - y) & ((x - y) >> 63))
}

// CheckInt32PowerOfTwo 检查数字是不是2的幂
func CheckInt32PowerOfTwo(x int32) bool {
	return (x & (x - 1)) == 0
}

// CountBitsSet 计算uint32中位数为1的数量
// 这个又叫做 hammingWeight -- LT 191
func CountBitsSet(x uint32) uint32 {
	var r uint32
	for ; x > 0; x >>= 1 {
		r += x & 1
	}
	return r
}

// ComputingParity uint32的二进制中1的数量是否为奇数
func ComputingParity(x uint32) bool {
	var flag bool
	for ; x > 0; x = x & (x - 1) {
		flag = !flag
	}
	return flag
}

// ReverseBits uint位反转
func ReverseBits(x uint) uint {
	r, s := x, bits.UintSize-1
	for x = x >> 1; x > 0; x >>= 1 {
		r <<= 1
		r |= x & 1
		s--
	}
	r <<= s
	return r
}
