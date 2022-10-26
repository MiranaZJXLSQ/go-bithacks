package go_bithacks

// NextTwoPowerOfInt32 获取下一个大于等于一个int32的2的幂
func NextTwoPowerOfInt32(x int32) int32 {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}

// NextTwoPowerOfInt64 获取下一个大于等于一个int64的2的幂
func NextTwoPowerOfInt64(x int64) int64 {
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
