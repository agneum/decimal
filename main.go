// Package main allows to perform arithmetic operations with decimal numbers
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Decimal struct {
	precision int
	scale     int
}

func (d *Decimal) String() string {
	precisionString := strconv.Itoa(d.precision)
	length := len(precisionString)

	if length <= d.scale {
		precisionString = leftPad(precisionString, d.scale-length)
	}

	pointPosition := len(precisionString) - d.scale

	return fmt.Sprintf("%v.%v", precisionString[:pointPosition], precisionString[pointPosition:])
}

// NewDecimal creates a new Decimal struct by precision and scale numbers
func NewDecimal(precision int, scale int) *Decimal {
	return &Decimal{precision: precision, scale: scale}
}

// NewDecimalFromString creates a new Decimal struct from a string
func NewDecimalFromString(str string) (*Decimal, error) {
	var scale int

	str = cutSpaces(str)
	unsigedDecimalString := strings.Replace(str, "-", "", 1)
	precision, err := strconv.Atoi(strings.Replace(str, ".", "", 1))

	if err != nil {
		return NewDecimal(0, 0), err
	}

	pointIndex := strings.Index(unsigedDecimalString, ".")

	if pointIndex == -1 {
		scale = 0
	} else {
		scale = len(unsigedDecimalString) - pointIndex - 1
	}

	return NewDecimal(precision, scale), nil
}

// Add sums the two decimal numbers and returns the result in a Decimal struct
func Add(first *Decimal, second *Decimal) *Decimal {
	maxScale := max(first.scale, second.scale)
	precision := getRealValue(first, maxScale-first.scale) + getRealValue(second, maxScale-second.scale)

	return NewDecimal(precision, maxScale)
}

// Subtract finds the difference of two decimal numbers and returns the result in a Decimal struct
// Note the operation is not idempotent
func Subtract(first *Decimal, second *Decimal) *Decimal {
	maxScale := max(first.scale, second.scale)
	precision := getRealValue(first, maxScale-first.scale) - getRealValue(second, maxScale-second.scale)

	return NewDecimal(precision, maxScale)
}

// Multiply multiplies two decimal numbers and returns the result in a Decimal struct
func Multiply(first *Decimal, second *Decimal) *Decimal {
	return NewDecimal(first.precision*second.precision, first.scale+second.scale)
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

func cutSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func leftPad(str string, left int) string {
	a := make([]byte, left)
	for i := range a {
		a[i] = '0'
	}

	return string(a[:]) + str
}
