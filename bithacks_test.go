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
