package main

import "testing"

func TestDecorruption(t *testing.T) {
	data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161

	calculated := DeCorrupt(data)

	if calculated != expected {
		t.Errorf("expected %d, got %d", expected, calculated)
	}
}
