package allergies

const (
	eggs = uint(1 << iota)
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var keys = []uint{
	eggs,
	peanuts,
	shellfish,
	strawberries,
	tomatoes,
	chocolate,
	pollen,
	cats,
}

var values = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) (res []string) {

	for i, v := range keys {

		if v > score {
			break
		}

		if v&score != 0 {
			res = append(res, values[i])
		}
	}

	return
}

func AllergicTo(score uint, substance string) bool {

	for i, v := range values {
		if v == substance {
			return score&keys[i] != 0
		}
	}

	return false
}
