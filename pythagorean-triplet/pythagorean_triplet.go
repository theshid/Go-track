package pythagorean

import "math"



// Triplet ...
type Triplet [3]int

// Range ...
func Range(min, max int) (triplet []Triplet) {
	for i := min; i <= max; i++ {
		for j := i + 1; j < max; j++ {
			sqr := i*i + j*j
			sqrt := int(math.Sqrt(float64(sqr)))
			if sqr == sqrt*sqrt && sqrt <= max {
				triplet = append(triplet, Triplet{i, j, sqrt})
			}
		}
	}
	return
}

// Sum ...
func Sum(p int) (res []Triplet) {
	triplets := Range(1, p/2)
	for i := 0; i < len(triplets); i++ {
		if (triplets[i][0] + triplets[i][2] + triplets[i][1]) == p {
			res = append(res, triplets[i])
		}
	}
	return
}
