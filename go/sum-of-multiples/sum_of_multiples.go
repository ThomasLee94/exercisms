package summultiples

func SumMultiples(lim int, dividors []int) int {
	var output int

	for i := 0; i < lim; i++ { // iterate over numbers in range of limit
		for _, d := range dividors {

			// case: divide by 0
			if d == 0 {
				break
			}
			// case: i divides evenly by dividors then add it to the sum
			if i%d == 0 {
				output += i
				break
			}
		}
	}
	return output
}
