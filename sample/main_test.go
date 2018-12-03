package main

import "testing"

func TestMul(t *testing.T) {
	total := Mul(5, 4)
	if total != 20 {
		t.Errorf("Multiplication Error : %d, want: %d.", total, 20)
	}
}

func TestMin(t *testing.T) {
	min := Min(9, 1)
	if min != 1 {
		t.Errorf("Multiplication Error : %d, want: %d.", min, 1)
	}
}
