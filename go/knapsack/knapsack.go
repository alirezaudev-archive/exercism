package knapsack

type Item struct {
	Weight, Value int
}

func Knapsack(maximumWeight int, items []Item) int {
	return solve(maximumWeight, items, 0)
}

func solve(remaining int, items []Item, idx int) int {
	if idx >= len(items) || remaining <= 0 {
		return 0
	}

	bestWithout := solve(remaining, items, idx+1)

	item := items[idx]
	bestWith := 0
	if item.Weight <= remaining {
		bestWith = item.Value + solve(remaining-item.Weight, items, idx+1)
	}

	if bestWith > bestWithout {
		return bestWith
	}
	return bestWithout
}
