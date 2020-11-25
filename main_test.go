package main

import (
	"testing"
)

func Test_RobMoney(t *testing.T) {

	var testCases = []struct {
		name         string
		houses       []uint
		robbedAmount uint
	}{
		{"valid", []uint{}, 0},
		{"valid", []uint{1, 2, 3, 4}, 6},
		{"valid", []uint{2, 3, 2}, 4},
		{"valid", []uint{2, 1, 1, 3}, 5},
		{"valid", []uint{1, 2, 3, 4, 5, 6, 7, 100}, 112},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {
			if robbedAmount := robMoney(tt.houses...); robbedAmount != tt.robbedAmount {
				t.Errorf("worng value returned for input: %v\n output: %d; expected output:%d\n", tt.houses, robbedAmount, tt.robbedAmount)
			}
		})
	}
}
