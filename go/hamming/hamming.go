package hamming

import(
	"errors"
)

func Distance(a, b string) (int, error) {
	
	// COUNT THE NUM OF CHARACTERS THAT DONT MATCH
		counter := 0
		// COMPARE THE TWO STRANDS
		if len(a) != len(b){
			err := errors.New("length of the two DNA strands do not match")
			return -1,err
		}

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				counter ++
			}
		}

		return counter, nil
}
