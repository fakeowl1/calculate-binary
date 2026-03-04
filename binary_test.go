package main

import (
	"slices"
	"testing"
)

func TestConvertNumberToBinaryArray(t *testing.T) {
		tests := []struct {
				name	 string
				num		 uint
				length int
				want	 []uint8
		}{
				{"Positive 28", 28, 8, []uint8{0, 0, 0, 1, 1, 1, 0, 0}},
				{"Zero", 0, 8, []uint8{0, 0, 0, 0, 0, 0, 0, 0}},
				{"Max Byte", 255, 8, []uint8{1, 1, 1, 1, 1, 1, 1, 1}},
		}

		for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
						got := ConvertNumberToBinaryArray(tt.num, tt.length)
						if !slices.Equal(got, tt.want) {
							t.Errorf("got %v, want %v", got, tt.want)
						}
				})
		}
}


func TestGetTwosComplement(t *testing.T) {
		input := []uint8{0, 0, 0, 1, 1, 1, 0, 0}
		want := []uint8{1, 1, 1, 0, 0, 1, 0, 0}

		got := GetTwosComplement(input, 8)

		if !slices.Equal(got, want) {
			t.Errorf("GetTwosComplement() = %v, want %v", got, want)
		}
}
