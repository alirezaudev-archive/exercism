package yacht

func Score(dice []int, category string) int {
	counts := map[int]int{}
	sum := 0
	for _, n := range dice {
		counts[n]++
		sum += n
	}

	numbers := map[string]int{
		"ones":   1,
		"twos":   2,
		"threes": 3,
		"fours":  4,
		"fives":  5,
		"sixes":  6,
	}
	if number, ok := numbers[category]; ok {
		return counts[number] * number
	}

	switch category {
	case "full house":
		if len(counts) == 2 && (counts[dice[0]] == 3 || counts[dice[0]] == 2) {
			return sum
		}
	case "four of a kind":
		if len(counts) <= 2 {
			if counts[dice[0]] == 1 {
				return sum - dice[0]
			} else if counts[dice[0]] >= 4 {
				return dice[0] * 4
			}
		}
	case "little straight":
		if _, ok := counts[6]; !ok && len(counts) == 5 {
			return 30
		}
	case "big straight":
		if _, ok := counts[1]; !ok && len(counts) == 5 {
			return 30
		}
	case "choice":
		return sum
	case "yacht":
		if len(counts) == 1 {
			return 50
		}
	}

	return 0
}
