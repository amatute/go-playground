package problems

import "testing"

func TestRunningSum(t *testing.T) {

	tests := []struct {
		in []int
		want []int
	}{
		{
			[]int{1,2,3,4},
			[]int{1,3,6,10},
		},
		{
			[]int{1,1,1,1,1},
			[]int{1,2,3,4,5},
		},
		{
			[]int{3,1,2,10,1},
			[]int{3,4,6,16,17},
		},
	}

	assertOutput := func(want, output []int) {
		for i := 0; i < len(want); i++ {
			if want[i] != output[i] {
				t.Errorf("at index %d want %d got %d", i, want[i], output[i])
			}
		}
	}

	for _, tt := range tests {

		t.Run("It should return running sum of the array", func(t *testing.T) {
			ans := RunningSum(tt.in)
			assertOutput(tt.want, ans)
	
		})
		
	}

}

