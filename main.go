package main

import (
	"fmt"
	"math"
)

type Decimal struct {
	precision int
	scale     int
}

func main() {
	dec := NewDecimal(123456, 2)
	dec2 := NewDecimal(123456, 3)

	fmt.Printf("%v", Add(dec, dec2))
}

func NewDecimal(precision int, scale int) *Decimal {
	return &Decimal{precision: precision, scale: scale}
}

func NewDecimalFromString(str string) *Decimal {
	fmt.Println(str)
	// for _, val := range byteArray {
	// 	fmt.Printf("%v", val)
	// }
	return NewDecimal(0, 0)
}

func Add(first *Decimal, second *Decimal) *Decimal {
	maxScale := max(first.scale, second.scale)
	precision := getRealValue(first, maxScale-first.scale) + getRealValue(second, maxScale-second.scale)

	return NewDecimal(precision, maxScale)
}

func getRealValue(decimal *Decimal, scaleDiff int) int {
	return decimal.precision * int(math.Pow10(scaleDiff))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
