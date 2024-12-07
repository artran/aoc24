package main

import "testing"

type TestInput struct {
	data     string
	expected int
}

func TestDecorruption(t *testing.T) {
	test1 := TestInput{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161}
	test2 := TestInput{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48}
	tests := []TestInput{test1, test2}

	for idx, test := range tests {

		calculated := DeCorrupt(test.data)

		if calculated != test.expected {
			t.Errorf("Test %d: expected %d, got %d", idx, test.expected, calculated)
		}
	}
}
