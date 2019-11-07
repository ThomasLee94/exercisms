package diffsquares

func SumOfSquares(num int) int {
	output := 0
	for i := 0; i <= num; i++ {
		output += i * i
	}

	return output
}

func SquareOfSum(num int) int {
	output := 0
	for i := 0; i <= num; i++ {
		output += i
	}

	return output * output
}

func Difference(num int) int {
	num_1 := SumOfSquares(num)
	num_2 := SquareOfSum(num)

	return num_2 - num_1
}
