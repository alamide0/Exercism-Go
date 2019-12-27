package say

var (
	t1 = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	t2 = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	t3 = []string{"hundred", "thousand", "million", "billion"}
)

func Say(input int64) (string, bool) {

	if input == 0 {
		return t1[0], true
	}

	if input < 0 || input > 999999999999 {
		return "", false
	}

	//////////////////////////////////////First Step split input to int array///////////////////////////////
	nums, size := make([]int64, 14), 0

	for {
		nums[size] = input % 10
		input /= 10
		size++
		if input == 0 {
			break
		}
	}

	nums = nums[:size]
	for i, end := 0, size/2; i < end; i++ { //reverse
		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
	}
	//////////////////////////////////////First Step split input to int array///////////////////////////////

	//////////////////////////////////////Second Step get each part english text///////////////////////////////
	index, res := 0, ""

	for i := len(nums) - 3; i >= -2; i -= 3 {

		/**
		i == -2 the part only have one number
		i == -1 the part only have two number
		i == 0  the part have three number
		**/
		cur := ""

		if i >= 0 && nums[i] > 0 {
			cur = cur + t1[nums[i]] + " " + t3[0]
		}

		if i+1 >= 0 {
			if nums[i+1] > 1 {
				cur = format(cur, t2[nums[i+1]])
				if nums[i+2] > 0 {
					cur = cur + "-" + t1[nums[i+2]]
				}
			} else if nums[i+1] == 1 {
				cur = format(cur, t1[nums[i+1]*10+nums[i+2]])
			} else {
				if nums[i+2] > 0 {
					cur = format(cur, t1[nums[i+2]])
				}
			}
		} else if i+2 >= 0 {
			if nums[i+2] > 0 {
				cur = format(cur, t1[nums[i+2]])
			}
		}

		if index > 0 {
			if cur != "" {
				if res != "" {
					res = cur + " " + t3[index] + " " + res
				} else {
					res = cur + " " + t3[index]
				}
			}
		} else {
			res = cur
		}

		index += 1
	}
	//////////////////////////////////////Second Step get each part english text///////////////////////////////

	return res, true
}

func format(cur string, str string) string {
	if cur == "" {
		return str
	}

	return cur + " " + str
}
