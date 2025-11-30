package bookstore

var groupPrice = map[int]int{
	0: 800,
	1: 1520,
	2: 2160,
	3: 2560,
	4: 3000,
}

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	var counts [5]int
	for _, b := range books {
		counts[b-1]++
	}

	memo := make(map[[5]int]int)
	return bestCost(counts, memo)
}

func bestCost(counts [5]int, memo map[[5]int]int) int {
	if v, ok := memo[counts]; ok {
		return v
	}

	if isEmpty(counts) {
		memo[counts] = 0
		return 0
	}

	best := 1 << 60

	for mask := 1; mask < (1 << 5); mask++ {
		size := 0
		next := counts
		valid := true

		for i := 0; i < 5; i++ {
			if mask&(1<<i) != 0 {
				if next[i] == 0 {
					valid = false
					break
				}
				next[i]--
				size++
			}
		}

		if !valid {
			continue
		}

		priceGroup := groupPrice[size-1]
		total := priceGroup + bestCost(next, memo)

		if total < best {
			best = total
		}
	}

	memo[counts] = best
	return best
}

func isEmpty(counts [5]int) bool {
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
