package bookstore

import (
	"sort"
)

const (
	MAX = int(^uint(0) >> 1)
)

func Cost(basket []int) int {
	booksMap := [5]int{}

	for _, v := range basket {
		booksMap[v-1] += 1
	}

	n := getLen(booksMap)

	if n == 1 {
		return len(basket) * getMulti(n)
	}

	if n == len(basket) {
		return n * getMulti(n)
	}

	price := MAX

	sort.Sort(sort.Reverse(sort.IntSlice(booksMap[:])))
	////////////////////////////first method/////////////////////////////////////
	//find as long sub segments as possible
	for i := len(basket) / n; i <= n; i++ {

		p := calPrice(booksMap, i)

		if p < price {
			price = p
		}
	}
	/////////////////////////////first method////////////////////////////////////

	////////////////////////////second method////////////////////////////////////
	//the second method is slower than the first, but it is rigorous
	//step 1: find all the combinations
	//step 2: remove illegal
	//step 3: cacluate price
	// _, parts := Part(len(basket), getMax(booksMap), n)
	// for i := 0; i < len(parts); i++ {
	// 	if isIllegal(booksMap, parts[i]) {
	// 		continue
	// 	}
	// 	temp := 0
	// 	for j := 0; j < len(parts[0]); j++ {
	// 		temp += parts[i][j] * getMulti(parts[i][j])
	// 	}

	// 	if temp < price {
	// 		price = temp
	// 	}
	// }
	/////////////////////////////second method////////////////////////////////////

	return price
}

func calPrice(bMap [5]int, partSize int) (res int) {

	for {
		isEmpty := true
		ks := 0
		for i, v := range bMap {
			if v > 0 {
				bMap[i] -= 1
				ks++
				isEmpty = false
				if ks == partSize {
					break
				}
			}
		}

		res += ks * getMulti(ks)
		if isEmpty {
			break
		}
	}
	return
}

func getLen(bMap [5]int) (res int) {
	for _, v := range bMap {
		if v > 0 {
			res++
		}
	}
	return
}

func getMulti(n int) (res int) {
	switch n {
	case 1:
		res = 800
	case 2:
		res = 760
	case 3:
		res = 720
	case 4:
		res = 640
	case 5:
		res = 600
	}
	return
}

// func getMax(bMap [5]int) (res int) {
// 	for _, v := range bMap {
// 		if v > res {
// 			res = v
// 		}
// 	}
// 	return
// }

// func isIllegal(bMap [5]int, parts []int) bool {
// 	count := 0
// 	for {
// 		isEmpty := true
// 		ks := 0
// 		for i, v := range bMap {
// 			if v > 0 {
// 				bMap[i] -= 1
// 				ks++
// 				isEmpty = false
// 				if ks == parts[count] {
// 					break
// 				}
// 			}
// 		}

// 		if isEmpty {
// 			break
// 		}

// 		if ks != parts[count] {
// 			return true
// 		}
// 		count++
// 	}
// 	return false
// }

// func Part(ns, part, max int) (res int, data [][]int) {

// 	if part == 1 && ns > 0 {
// 		return 1, [][]int {{ns}}
// 	}

// 	if part == 1 && ns <= 0 {
// 		return 0, nil
// 	}
// 	start := 0

// 	if ns % part == 0 {
// 		start = ns / part
// 	}else {
// 		start = ns / part + 1
// 	}

// 	for i := start; i <= max; i++ {
// 		if n, d := Part(ns-i, part-1, i); d != nil {
// 			for j := 0; j < len(d); j++ {
// 				data = append(data, append([]int{i}, d[j]...))
// 			}
// 			res += n
// 		}
// 	}

// 	return
// }
