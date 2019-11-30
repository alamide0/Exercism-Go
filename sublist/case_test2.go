package sublist

var testCases2 = []struct {
	description string
	listOne     []int
	listTwo     []int
	expected    Relation
}{
	{
		description: "sublist at end",
		listOne:     []int{3, 4, 5},
		listTwo:     []int{0, 1, 2, 3, 4, 5},
		expected:    "sublist",
	},
}