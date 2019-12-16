package series

func All(n int, series string) (res []string) {
	if n > len(series){
		return nil
	}

	for i := 0; i < len(series) - n + 1; i++ {
		res = append(res, series[i: n+i])
	}

	return
}

func UnsafeFirst(n int, series string) (res string){
	return All(n, series)[0]
}

func First(n int, series string)(res string, ok bool){
	if n > len(series) {
		return res, false
	}

	return All(n, series)[0], true
}
