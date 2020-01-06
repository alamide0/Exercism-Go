package armstrong

const (
	initSize = 9
)

func IsNumber(n int) bool {
	nums := make([]int, initSize)
	count, tmp := 0, n
	for {
		if count >= initSize{
			nums = append(nums, n % 10)
		} else {
			nums[count] = n % 10
		}
		n /= 10
		count++
		if n == 0 {
			break
		}
	}

	nums = nums[:count]

	sum := 0
	for _, v := range nums {
		sum += pow(v, count)
	}

	return sum == tmp
}

func pow(n, t int) int {
	res := 1
	for i := 0; i < t; i++ {
		res *= n 
	}

	return res
}