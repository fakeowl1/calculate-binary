package main

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestConvertFractionToBinary(t *testing.T) {
	tests := []struct {
			name string
			fraction float64 
			length int
			want	 []uint8
	}{
		{"Zero", 0.0, 11, []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"Length: 11", 0.66, 11, []uint8{1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1}},
		{"Length: 20", 0.66, 20, []uint8{1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0}},
		{"Small number", 0.001, 11, []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			got := convertFractionToBinary(tt.fraction, tt.length)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFloatBinary(t *testing.T) {
	tests := []struct {
		name string
		fraction float64
		exponentLength int
		lengthMantissa int
		bias int
		want string
	}{
		{"float", 1.055, 8, 23, 127, fmt.Sprintf("%032b", math.Float32bits(1.055))},
		{"double", 0.00000858583, 11, 52, 1023, fmt.Sprintf("%064b", math.Float64bits(0.00000858583))},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			got := NewFloatBinary(tt.fraction, tt.exponentLength, tt.lengthMantissa, tt.bias)
			gotString := got.toString()
			if (gotString != tt.want) {
				t.Errorf("got %v, want %s", got, tt.want)
			}
		})
	}
}
