package binarysearch

func SearchInts(slice []int, key int) int {

	start, end := 0, len(slice)-1

	for start <= end {
		p := (start + end) / 2

		if slice[p] == key {
			return p
		} else if slice[p] > key {
			end = p - 1
		} else if slice[p] < key {
			start = p + 1
		}
	}

	return -1
}
