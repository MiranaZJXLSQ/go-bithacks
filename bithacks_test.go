package go_bithacks

import "testing"

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
