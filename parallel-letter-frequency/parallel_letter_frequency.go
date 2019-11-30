package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strs []string) FreqMap {
	ch := make(chan FreqMap, len(strs))

	result := FreqMap{}
	
	for _, s := range strs {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range strs {
		for k, v := range <- ch{
			result[k] += v
		}
	}

	return result
}
