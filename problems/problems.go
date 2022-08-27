package problems

// 13. Roman to Integer
func RomanToInt(s string) int {
	number := []string{}
	for _, c := range s {
		number = append(number, string(c))
	}

	r := 0
	t := 0

	for i := 0; i < len(number); i++ {
		letterValue := getValue(number[i])

		if i == 0 {
			t = letterValue
			continue
		}

		if letterValue < t {
			r += t
			t = letterValue
			continue
		}

		if letterValue == t {
			t += letterValue
		}

		if letterValue > t {
			t = letterValue - t
		}
	}

	return r + t
}

func getValue(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

// 1. Two Sum
func TwoSum(nums []int, target int) []int {
	output := []int{}
	numsMap := map[int]int{}

	for i, n := range nums {
		numsMap[n] = i
	}

	for i, n := range nums {
		s := target - n
		if v, ok := numsMap[s]; ok {
			if i != v {
				output = append(output, i, v)
				return output
			}
		}
	}

	return output
}

// 724. Find Pivot Index
func PivotIndex(nums []int) int {
	leftSum := 0
	sum := 0
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		if leftSum == sum - leftSum - nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}