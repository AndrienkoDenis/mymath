package mymath

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	// Тестирование функции Sqrt
	cases := []struct {
		in, want float64
	}{
		{4, math.Sqrt(4)},
		{9, math.Sqrt(9)},
		{16, math.Sqrt(16)},
	}

	for _, c := range cases {
		got := Sqrt(c.in)
		if got != c.want {
			t.Errorf("Sqrt(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCeil(t *testing.T) {
	// Тестирование функции Ceil
	cases := []struct {
		in, want float64
	}{
		{3.2, math.Ceil(3.2)},
		{5.7, math.Ceil(5.7)},
		{7.0, math.Ceil(7.0)},
	}

	for _, c := range cases {
		got := Ceil(c.in)
		if got != c.want {
			t.Errorf("Ceil(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFloor(t *testing.T) {
	// Тестирование функции Floor
	cases := []struct {
		in, want float64
	}{
		{3.2, math.Floor(3.2)},
		{5.7, math.Floor(5.7)},
		{7.0, math.Floor(7.0)},
	}

	for _, c := range cases {
		got := Floor(c.in)
		if got != c.want {
			t.Errorf("Floor(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPow(t *testing.T) {
	// Тестирование функции Pow
	cases := []struct {
		x, y, want float64
	}{
		{2, 3, math.Pow(2, 3)},
		{3, 2, math.Pow(3, 2)},
		{4, 0.5, math.Pow(4, 0.5)},
	}

	for _, c := range cases {
		got := Pow(c.x, c.y)
		if got != c.want {
			t.Errorf("Pow(%v, %v) == %v, want %v", c.x, c.y, got, c.want)
		}
	}
}

func TestMax(t *testing.T) {
	// Тестирование функции Max
	cases := []struct {
		x, y, want float64
	}{
		{2, 3, math.Max(2, 3)},
		{5, 2, math.Max(5, 2)},
		{7, 7, math.Max(7, 7)},
	}

	for _, c := range cases {
		got := Max(c.x, c.y)
		if got != c.want {
			t.Errorf("Max(%v, %v) == %v, want %v", c.x, c.y, got, c.want)
		}
	}
}

func TestMin(t *testing.T) {
	// Тестирование функции Min
	cases := []struct {
		x, y, want float64
	}{
		{2, 3, math.Min(2, 3)},
		{5, 2, math.Min(5, 2)},
		{7, 7, math.Min(7, 7)},
	}

	for _, c := range cases {
		got := Min(c.x, c.y)
		if got != c.want {
			t.Errorf("Min(%v, %v) == %v, want %v", c.x, c.y, got, c.want)
		}
	}
}
