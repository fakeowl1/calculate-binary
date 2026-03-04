package main

import (
	"fmt"
	"math"
	"strings"
)

type FloatBinary struct {
	Sign uint8
	Exponent []uint8
	BiasedExp int 
	Mantissa []uint8
}

func (fb *FloatBinary) toString() (string) {
	var sb strings.Builder
	
	fmt.Fprintf(&sb, "%d", fb.Sign)

	for _, value := range fb.Exponent {
		fmt.Fprintf(&sb, "%d", value)
	}
	
	for _, value := range fb.Mantissa {
		fmt.Fprintf(&sb, "%d", value)
	}

	return sb.String()
}

func RoundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func convertFractionToBinary(fraction float64, length int) ([]uint8){
	binaryNumber := make([]uint8, 0, length)
	
	for i := range length {
		fraction *= 2
		
		if fraction >= 1 {
			binaryNumber = append(binaryNumber, 1)
		} else {
			binaryNumber = append(binaryNumber, 0)
		}

		fmt.Printf("Step %d: %F * 2 = %F (%d)\n", i+1, fraction/2, fraction, binaryNumber[i])
		_, fraction = math.Modf(fraction)
	}

	return binaryNumber
}

func NewFloatBinary(number float64, exponentLength int, lengthMantissa int, bias int) (*FloatBinary) {
    if number == 0 {
        return &FloatBinary{
					Sign: 0,
					Exponent: nil,
					BiasedExp: 127,
					Mantissa: nil,
				}
    }

		var sign = 0
		if (number < 0) {
			sign = 1
		}

    var mainPart, fractionPart = math.Modf(math.Abs(number))

		mainPartBitArray := ConvertNumberToBinaryArray(uint(mainPart), exponentLength)
		fractionPartBitArray := convertFractionToBinary(fractionPart, lengthMantissa + 32)
		allBits := append(mainPartBitArray, fractionPartBitArray...)
    
		firstOneIdx := -1
		for i, bit := range allBits {
			if bit == 1 {
				firstOneIdx = i
				break
			}
		}

		commaIdx := len(mainPartBitArray)
		exponent := commaIdx - 1 - firstOneIdx

		biasedExp := bias + exponent
		expBits := ConvertNumberToBinaryArray(uint(biasedExp), exponentLength)
		mantissa := allBits[firstOneIdx+1:]
		
		if (len(mantissa) > lengthMantissa) {
			mantissa = mantissa[:lengthMantissa]
		}

		for len(mantissa) < lengthMantissa {
				mantissa = append(mantissa, 0)
		}

		return &FloatBinary{
			Sign: uint8(sign),
			Exponent: expBits,
			BiasedExp: biasedExp,
			Mantissa: mantissa,
		}
}
