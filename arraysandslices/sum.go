package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for i, numbers := range numbers {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(numbers[1:])
		}
	}

	return sums
}
