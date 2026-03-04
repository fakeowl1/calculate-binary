# Конвертор
Невеличка реалізація конвертація чисел у двійковий код для виконання лабораторної роботи №2 з Системного Програмування. Павлов В.Г заїбав.

## Приклад використання

```go
package main

import "fmt"

func main() {
	// Вхідні дані
	A_val := 13
	B_val := 1306
	C_val := 13061941
	N_val := 9365.0

	// Розрахунок дійсних чисел
	D := float64(A_val) / N_val
	E := float64(B_val) / N_val
	F := float64(C_val) / N_val

	fmt.Println("=== 1. ЦІЛІ ЧИСЛА (Послідовне представлення) ===")

	// --- BYTE (8 bit) ---
	bitsA8 := ConvertNumberToBinaryArray(uint(A_val), 8)
	fmt.Printf("Byte A (28):   %v\n", bitsA8)
	fmt.Printf("Byte -A (-28): %v\n", GetTwosComplement(bitsA8, 8))

	// --- WORD (16 bit) ---
	bitsA16 := ConvertNumberToBinaryArray(uint(A_val), 16)
	bitsB16 := ConvertNumberToBinaryArray(uint(B_val), 16)
	fmt.Printf("\nWord A (28):    %v\n", bitsA16)
	fmt.Printf("Word -A (-28):  %v\n", GetTwosComplement(bitsA16, 16))
	fmt.Printf("Word B (2810):  %v\n", bitsB16)
	fmt.Printf("Word -B (-2810):%v\n", GetTwosComplement(bitsB16, 16))

	// --- SHORTINT (32 bit) ---
	bitsA32 := ConvertNumberToBinaryArray(uint(A_val), 32)
	bitsB32 := ConvertNumberToBinaryArray(uint(B_val), 32)
	bitsC32 := ConvertNumberToBinaryArray(uint(C_val), 32)
	fmt.Printf("\nShortInt A (28):    %v\n", bitsA32)
	fmt.Printf("ShortInt -A (-28):  %v\n", GetTwosComplement(bitsA32, 32))
	fmt.Printf("ShortInt B (2810):  %v\n", bitsB32)
	fmt.Printf("ShortInt -B (-2810):%v\n", GetTwosComplement(bitsB32, 32))
	fmt.Printf("ShortInt C:         %v\n", bitsC32)
	fmt.Printf("ShortInt -C:        %v\n", GetTwosComplement(bitsC32, 32))

	// --- LONGINT (64 bit) ---
	bitsA64 := ConvertNumberToBinaryArray(uint(A_val), 64)
	bitsB64 := ConvertNumberToBinaryArray(uint(B_val), 64)
	bitsC64 := ConvertNumberToBinaryArray(uint(C_val), 64)
	fmt.Printf("\nLongInt A (28):    %v\n", bitsA64)
	fmt.Printf("LongInt -A (-28):  %v\n", GetTwosComplement(bitsA64, 64))
	fmt.Printf("LongInt B (2810):  %v\n", bitsB64)
	fmt.Printf("LongInt -B (-2810):%v\n", GetTwosComplement(bitsB64, 64))
	fmt.Printf("LongInt C:         %v\n", bitsC64)
	fmt.Printf("LongInt -C:        %v\n", GetTwosComplement(bitsC64, 64))

	fmt.Println("\n=== 2. ДІЙСНІ ЧИСЛА (Послідовне представлення) ===")

	// --- SINGLE (4 байти) ---
	sD := NewFloatBinary(D, 8, 23, 127)
	sMD := NewFloatBinary(-D, 8, 23, 127)
	fmt.Printf("Single D (%.6f):  %s\n", D, sD.toString())
	fmt.Printf("Single -D:             %s\n", sMD.toString())

	// --- DOUBLE (8 байтів) ---
	dE := NewFloatBinary(E, 11, 52, 1023)
	dME := NewFloatBinary(-E, 11, 52, 1023)
	fmt.Printf("\nDouble E (%.6f):  %s\n", E, dE.toString())
	fmt.Printf("Double -E:             %s\n", dME.toString())

	// --- EXTENDED (10 байтів) ---
	exF := NewFloatBinary(F, 15, 64, 16383)
	exMF := NewFloatBinary(-F, 15, 64, 16383)
	fmt.Printf("\nExtended F (%.6f): %s\n", F, exF.toString())
	fmt.Printf("Extended -F:            %s\n", exMF.toString())
}
```

## Запуск тестів

```bash
go test
```
