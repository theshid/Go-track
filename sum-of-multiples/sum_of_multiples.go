package summultiples

func SumMultiples(limit int, divisors ...int) int {
	if limit < 2 || len(divisors) == 0 {
		return 0
	}
	var sum int
	multiples := make(map[int]struct{})
	for _, v := range divisors {
		if v==0 {
			continue
		}
		for j := v; j < limit; j += v {
			if _, ok := multiples[j]; !ok {
				multiples[j] = struct{}{}
				sum += j
			}
		}
	}
	return sum
}
