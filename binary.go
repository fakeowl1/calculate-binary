package main

import "fmt"

func PrintBitArray(arr []uint8) {
	for _, value := range arr {
		fmt.Printf("%d", value)
	}
}

func ConvertNumberToBinaryArray(number uint, length int) ([]uint8) {
	binaryNumber := make([]uint8, 0, length)

	for i := 0; i < length; i++ {
		binaryNumber = append(binaryNumber, uint8(number&1))
		number >>= 1
	}

	for i, j := 0, len(binaryNumber)-1; i < j; i, j = i+1, j-1 {
		binaryNumber[i], binaryNumber[j] = binaryNumber[j], binaryNumber[i]
 	}

	return binaryNumber
}

func GetTwosComplement(bitArray []uint8, length int) ([]uint8) {
	newArray := make([]uint8, 0, length)

	for _, value := range bitArray {
		newArray = append(newArray, value ^ 1)
	}

	fmt.Printf("Зворотний код: ")
	PrintBitArray(newArray)
	fmt.Printf("\n")
	fmt.Printf("Додана одиниця: ")
	carry := uint8(1)
	for i := len(newArray) - 1; i >= 0 && carry == 1; i-- {
		sum := newArray[i] + carry
		newArray[i] = sum % 2
		carry = sum / 2
	}
	PrintBitArray(newArray)
	fmt.Printf("\n")

	return newArray
}
