package problems

import "testing"

func TestRomanToInt(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("It should return the int version of a roman version number", func(t *testing.T) {
		got := RomanToInt("I")
		want := 1
		assertCorrectMessage(t, got, want)

		got = RomanToInt("VIII")
		want = 8
		assertCorrectMessage(t, got, want)

		got = RomanToInt("IX")
		want = 9
		assertCorrectMessage(t, got, want)

		got = RomanToInt("LVIII")
		want = 58
		assertCorrectMessage(t, got, want)

		got = RomanToInt("MCMXCIV")
		want = 1994
		assertCorrectMessage(t, got, want)

	})
}

func TestTwoSum(t *testing.T) {

	assertOutput := func(want, output []int) {
		for i := 0; i < len(want); i++ {
			if want[i] != output[i] {
				t.Errorf("index %d is diferent than %d", want[i], output[i])
			}
		}
	}

	t.Run("It should returns the indexes of the two numbers such when they are added is equal to target", func(t *testing.T) {

		nums := []int{2,7,11,15}
		target := 9
		want := []int{0,1}
		output := TwoSum(nums, target)
		assertOutput(want, output)

		nums = []int{3,2,4}
		target = 6
		want = []int{1,2}
		output = TwoSum(nums, target)
		assertOutput(want, output)

	})
	
}
