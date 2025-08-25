package main

import "testing"

func TestGroupTemp(t *testing.T) {
	tests := []struct {
		num      int64
		pos      uint64
		bit      bool
		expected int64
	}{
		{5, 0, false, 4},
		{5, 2, true, 5},
		{5, 3, true, 13},
		{15, 4, true, 31},
		{16, 4, false, 0},
	}

	for _, tc := range tests {
		res := setBit(tc.num, tc.pos, tc.bit)
		if res != tc.expected {
			t.Errorf("SetBit(%d, %d, %t) = %d, want %d",
				tc.num, tc.pos, tc.bit, res, tc.expected)
		}
	}
}
