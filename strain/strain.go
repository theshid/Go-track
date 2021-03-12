package strain

const testVersion = 1

// Ints defines a collection of int values
type Ints []int

// Lists defines a collection of arrays of ints
type Lists [][]int

// Strings defines a collection of strings
type Strings []string

func (i Ints) Keep(Strain func(int) bool) Ints {
	var result Ints
	for _, v := range i {
		if Strain(v) {
			result = append(result, v)
		}
	}
	return result
}

func (i Ints) Discard(Strain func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !Strain(n) })
}

func (l Lists) Keep(Strain func([]int) bool) Lists {
	var result Lists
	for _, v := range l {

		if Strain(v) {
			result = append(result, v)
		}

	}

	return result
}

func (s Strings) Keep(Strain func(string) bool) Strings {
	var result Strings

	for _, v := range s {
		if Strain(v) {
			result = append(result, v)
		}

	}
	return result
}
