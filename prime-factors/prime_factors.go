package prime

func Factors(input int64) []int64{

	res := make([]int64, 0)
	t := input
	
	for i := int64(2); i <= input; {

		if t%i == 0 {
			t = t / i
			res = append(res, i)
			if t == 1 {
				break
			}
		} else {
			i++
		}

	}

	return res
}
