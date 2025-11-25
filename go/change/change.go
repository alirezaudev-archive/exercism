package change

import (
	"errors"
	"sort"
)

func Change(coins []int, target int) ([]int, error) {
	if target < 0 {
		return nil, errors.New("negative target")
	}
	if target == 0 {
		return []int{}, nil
	}

	sort.Ints(coins)

	dp := make([]int, target+1)
	parent := make([]int, target+1)

	const INF = 1 << 30
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	for _, c := range coins {
		for amt := c; amt <= target; amt++ {
			if dp[amt-c]+1 < dp[amt] {
				dp[amt] = dp[amt-c] + 1
				parent[amt] = c
			}
		}
	}

	if dp[target] == INF {
		return nil, errors.New("no solution")
	}

	var result []int
	cur := target
	for cur > 0 {
		coin := parent[cur]
		result = append(result, coin)
		cur -= coin
	}

	sort.Ints(result)

	return result, nil
}
