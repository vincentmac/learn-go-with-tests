package main

func Sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, a := range numbersToSum {
		if len(a) == 0 {
			sums = append(sums, 0)
			continue
		}
		tails := a[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}
