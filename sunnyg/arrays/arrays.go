package arrays

// Sums a slice of ints.
func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

// Sums an arbitrary number of slices of ints.
func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return
}

// Sums the tails of an arbitrary number of slices of ints.
func SumAllTails(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
