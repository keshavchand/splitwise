package calculator

import "testing"

func TestCalc(t *testing.T) {
	tests := []struct {
		input  string
		output float64
	}{
		{"2 3 +", 5},
		{"2 3 -", -1},
		{"2 3 *", 6},
		{"2 3 /", 2.0 / 3.0},
		{"2 3 + 4 *", 20},
		{"2 3 4 + *", 14},
		{"2 3 4 + * 5 -", 9},
		{"2 3 4 + * 5 - 2 /", 4.5},
	}

	for _, test := range tests {
		result, err := CalculatePostFix(test.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != test.output {
			t.Errorf("For Expression %s: Expected %v, got %v", test.input, test.output, result)
		}
	}
}
