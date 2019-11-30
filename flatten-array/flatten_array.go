package flatten

func Flatten(value interface{}) []interface{}{
	result := []interface{}{}

	if value == nil {
		return result
	}

	if v, ok := value.([]interface{}); ok {
		for _, item := range v {
			if item == nil {
				continue
			}
			result = append(result, Flatten(item)...)
		}
	} else {
		result = append(result, value)
	}

	return result
}