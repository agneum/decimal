package main

import "testing"

func TestAdd(t *testing.T) {
	first, second := NewDecimal(555, 1), NewDecimal(222, 1)

	sum := Add(first, second)
	if sum.precision != 777 && sum.scale != 1 {
		t.Error("Not Equals")
	}
}
