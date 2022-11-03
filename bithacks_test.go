package go_bithacks

import (
	"math/bits"
	"testing"
)

func TestDetectTwoIntegerOppositeSign(t *testing.T) {
	var cases [][]int32
	cases = append(cases, []int32{1, -1})
	cases = append(cases, []int32{1, 1})
	for _, v := range cases {
		t.Log(DetectTwoIntegerOppositeSign(v[0], v[1]))
	}
	for _, v := range cases {
		t.Log(DetectTwoIntegerSameSign(v[0], v[1]))
	}
}

func TestAbsoluteValueOfInt32(t *testing.T) {
	int32cases := []int32{1, -1, -2, -3, -111111}
	for _, v := range int32cases {
		t.Log(AbsoluteValueOfInt32(v))
	}
	int64cases := []int64{1, -1, -2, -3, -111111}
	for _, v := range int64cases {
		t.Log(AbsoluteValueOfInt64(v))
	}
}

func TestMinOfTwoInt32(t *testing.T) {
	var int32cases [][]int32
	int32cases = append(int32cases, []int32{1, -1})
	int32cases = append(int32cases, []int32{1, 2})
	for _, v := range int32cases {
		t.Log(MinOfTwoInt32(v[0], v[1]))
		t.Log(MaxOfTwoInt32(v[0], v[1]))
	}
	var int64cases [][]int64
	int64cases = append(int64cases, []int64{1, -1})
	int64cases = append(int64cases, []int64{1, 2})
	for _, v := range int64cases {
		t.Log(MinOfTwoInt64(v[0], v[1]))
		t.Log(MaxOfTwoInt64(v[0], v[1]))
	}
}

func TestCheckInt32PowerOfTwo(t *testing.T) {
	var cases []int32
	cases = append(cases, 1)
	cases = append(cases, 2)
	cases = append(cases, 3)
	cases = append(cases, 4)
	cases = append(cases, 5)
	for _, v := range cases {
		t.Log(CheckInt32PowerOfTwo(v))
	}
}

func TestCountBitsSet(t *testing.T) {
	var cases []uint32
	cases = append(cases, 1)
	cases = append(cases, 2)
	cases = append(cases, 3)
	cases = append(cases, 4)
	cases = append(cases, 5)
	for _, v := range cases {
		t.Log(CountBitsSet(v))
	}
}

func TestComputingParity(t *testing.T) {
	var cases []uint32
	cases = append(cases, 1)
	cases = append(cases, 2)
	cases = append(cases, 3)
	cases = append(cases, 4)
	cases = append(cases, 5)
	for _, v := range cases {
		t.Log(ComputingParity(v))
	}
}

func TestReverseBits(t *testing.T) {
	var cases []uint
	cases = append(cases, 1)
	cases = append(cases, 2)
	cases = append(cases, 3)
	cases = append(cases, 4)
	cases = append(cases, 5)
	for _, v := range cases {
		t.Log(ReverseBits(v))
		t.Log(bits.Reverse(v))
	}
}

func TestUpperCaseToLowerCase(t *testing.T) {
	var cases []rune
	cases = append(cases, 'A')
	cases = append(cases, 'B')
	cases = append(cases, 'C')
	cases = append(cases, 'D')
	cases = append(cases, 'E')
	for _, v := range cases {
		t.Logf("origin %c, transfered %c", v, UpperCaseToLowerCase(v))
	}
}

func TestLowerCaseToUpperCase(t *testing.T) {
	var cases []rune
	cases = append(cases, 'a')
	cases = append(cases, 'b')
	cases = append(cases, 'c')
	cases = append(cases, 'd')
	cases = append(cases, 'e')
	for _, v := range cases {
		t.Logf("origin %c, transfered %c", v, LowerCaseToUpperCase(v))
	}
}
