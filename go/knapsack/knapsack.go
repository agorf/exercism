package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
// Source: https://en.wikipedia.org/wiki/Knapsack_problem#0-1_knapsack_problem
func Knapsack(maximumWeight int, items []Item) int {
	n := len(items)
	weights := make([]int, n)
	values := make([]int, n)
	m := make([][]int, n+1)
	for i, item := range items {
		weights[i] = item.Weight
		values[i] = item.Value
	}
	for i := range n + 1 {
		m[i] = make([]int, maximumWeight+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= maximumWeight; j++ {
			if weights[i-1] > j {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = max(m[i-1][j], m[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return m[n][maximumWeight]
}
