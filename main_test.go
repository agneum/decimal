package decimal

import (
	"reflect"
	"testing"
)

var AddTests = []struct {
	first  Decimal
	second Decimal
	result Decimal
}{
	{Decimal{555, 1}, Decimal{222, 1}, Decimal{777, 1}},
	{Decimal{555, 1}, Decimal{225, 1}, Decimal{780, 1}},
	{Decimal{555, 1}, Decimal{222, 2}, Decimal{5772, 2}},
	{Decimal{555, 0}, Decimal{222, 1}, Decimal{5772, 1}},
	{Decimal{555, 1}, Decimal{-222, 1}, Decimal{333, 1}},
	{Decimal{-555, 1}, Decimal{-222, 1}, Decimal{-777, 1}},
	{Decimal{55, 1}, Decimal{-222, 1}, Decimal{-167, 1}},
	{Decimal{-555, 3}, Decimal{222, 1}, Decimal{21645, 3}},
}

var SubstractTests = []struct {
	first  Decimal
	second Decimal
	result Decimal
}{
	{Decimal{555, 1}, Decimal{222, 1}, Decimal{333, 1}},
	{Decimal{555, 1}, Decimal{222, 2}, Decimal{5328, 2}},
	{Decimal{555, 1}, Decimal{-222, 1}, Decimal{777, 1}},
	{Decimal{-555, 1}, Decimal{-222, 1}, Decimal{-333, 1}},
	{Decimal{55, 1}, Decimal{-222, 1}, Decimal{277, 1}},
	{Decimal{-555, 3}, Decimal{222, 1}, Decimal{-22755, 3}},
}

var MultiplyTests = []struct {
	first  Decimal
	second Decimal
	result Decimal
}{
	{Decimal{555, 1}, Decimal{222, 1}, Decimal{123210, 2}},
	{Decimal{555, 1}, Decimal{222, 2}, Decimal{123210, 3}},
	{Decimal{555, 1}, Decimal{-222, 1}, Decimal{-123210, 2}},
	{Decimal{-55, 1}, Decimal{-22, 1}, Decimal{1210, 2}},
	{Decimal{4, 2}, Decimal{100, 0}, Decimal{400, 2}},
	{Decimal{-555, 3}, Decimal{222, 1}, Decimal{-123210, 4}},
}

func TestAdd(t *testing.T) {
	for _, tt := range AddTests {
		actual := Add(&tt.first, &tt.second)
		if !reflect.DeepEqual(actual, &tt.result) {
			t.Errorf("Add(%v, %v): expected %v, actual %v", &tt.first, &tt.second, &tt.result, actual)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, tt := range SubstractTests {
		actual := Subtract(&tt.first, &tt.second)
		if !reflect.DeepEqual(actual, &tt.result) {
			t.Errorf("Subtract(%v, %v): expected %v, actual %v", &tt.first, &tt.second, &tt.result, actual)
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, tt := range MultiplyTests {
		actual := Multiply(&tt.first, &tt.second)
		if !reflect.DeepEqual(actual, &tt.result) {
			t.Errorf("Multiply(%v, %v): expected %v, actual %v", &tt.first, &tt.second, &tt.result, actual)
		}
	}
}
