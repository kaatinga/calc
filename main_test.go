package calc

import "testing"

func TestCalc(t *testing.T) {

	tests := []struct {
		name       string
		expression string
		wantResult int64
		wantErr    bool
	}{
		{"sum1", "1+2", 3, false},
		{"sum2", "1 + 2", 3, false},
		{"sum3", "1 +2", 3, false},
		{"sum4", "1+ 2", 3, false},
		{"sum5", "11+22", 33, false},
		{"sum6", "11 + 22", 33, false},
		{"multiply1", "11*22", 242, false},
		{"multiply2", "1*2", 2, false},
		{"multiply3", "1 *2", 2, false},
		{"multiply4", "1* 2", 2, false},
		{"divide by zero", "1/0", 0, true},
		{"incorrect expression1", "1", 0, true},
		{"incorrect expression2", "1+", 0, true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Calc(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Calc() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
