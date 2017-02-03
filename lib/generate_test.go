package lib

import "testing"

func TestRandIntFromRangeIsInRange(t *testing.T) {
	type rge struct {
		min int
		max int
	}

	tests := []rge{
		{min: 2, max: 5},
		{min: 11, max: 19},
	}

	for _, tt := range tests {
		i := RandIntFromRange(tt.min, tt.max)
		if !(tt.min <= i && i <= tt.max) {
			t.Errorf("RandIntFromRange(%d, %d) does not produce value inbetween range.", tt.min, tt.max)
		}
	}
}
