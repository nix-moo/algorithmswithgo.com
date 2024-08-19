func Solution(U int, weight []int) int {

	turns := 0
	c1 := 0
	c2 := 1
	for c2 < len(weight) {
		w := weight[c1] + weight[c2]

		if w <= U {
			c1 = c2
			c2++
		} else {
			turns++
			if weight[c1] >= weight[c2] {
				c1 = c2
				c2++
			} else {
				c2++
			}
		}
	}
	return turns
}

/*
Given an array of weights and a maximum, find the minimum number of ints to remove in order for all the weights to be under the maximum when taken 2 at.
*/