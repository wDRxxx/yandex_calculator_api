package calculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	c := NewCalculatorService()

	tt := []struct {
		input    string
		expected float64
		err      bool
	}{
		{
			input:    "2+2",
			expected: 4,
			err:      false,
		},
		{
			input:    "2+2*2",
			expected: 6,
			err:      false,
		},
		{
			input:    "2/2",
			expected: 1,
			err:      false,
		},
		{
			input:    "9*18/5",
			expected: 32.4,
			err:      false,
		},
		{
			input:    "zxczxczxc",
			expected: 0,
			err:      true,
		},
	}

	for _, tc := range tt {
		res, err := c.Calculate(tc.input)
		if tc.err && err == nil || !tc.err && err != nil {
			t.Errorf("input: %s, expected %f, got %s", tc.input, tc.expected, err)
		}
		if res != tc.expected {
			t.Errorf("input: %s, expected %f, got %f", tc.input, tc.expected, res)
		}
	}
}
