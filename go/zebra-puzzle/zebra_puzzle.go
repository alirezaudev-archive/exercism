package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

const (
	// nationalities
	englishman = 0
	spaniard   = 1
	ukrainian  = 2
	norwegian  = 3
	japanese   = 4

	// colors
	red    = 0
	green  = 1
	ivory  = 2
	yellow = 3
	blue   = 4

	// drinks
	water       = 0
	coffee      = 1
	tea         = 2
	milk        = 3
	orangeJuice = 4

	// pets
	dog    = 0
	snails = 1
	fox    = 2
	horse  = 3
	zebra  = 4

	// smokes
	oldGold       = 0
	kools         = 1
	chesterfields = 2
	luckyStrike   = 3
	parliaments   = 4
)

var nationalities = map[int]string{
	englishman: "Englishman",
	spaniard:   "Spaniard",
	ukrainian:  "Ukrainian",
	norwegian:  "Norwegian",
	japanese:   "Japanese",
}

func SolvePuzzle() Solution {
	indices := []int{0, 1, 2, 3, 4}

	for _, nationality := range permutations(indices) {
		// The Norwegian lives in the first house.
		if nationality[0] != norwegian {
			continue
		}

		// The Englishman lives in the red house.
		for _, color := range permutations(indices) {
			if pos(nationality, englishman) != pos(color, red) {
				continue
			}
			// The green house is immediately to the right of the ivory house.
			if pos(color, green) != pos(color, ivory)+1 {
				continue
			}
			// The Norwegian lives next to the blue house.
			if !nextTo(pos(color, blue), pos(nationality, norwegian)) {
				continue
			}

			for _, drink := range permutations(indices) {
				// Milk is drunk in the middle house.
				if pos(drink, milk) != 2 {
					continue
				}
				// The Ukrainian drinks tea.
				if pos(nationality, ukrainian) != pos(drink, tea) {
					continue
				}
				// Coffee is drunk in the green house.
				if pos(drink, coffee) != pos(color, green) {
					continue
				}

				for _, pet := range permutations(indices) {
					// The Spaniard owns the dog.
					if pos(nationality, spaniard) != pos(pet, dog) {
						continue
					}

					for _, smoke := range permutations(indices) {
						// The Old Gold smoker owns snails.
						if pos(smoke, oldGold) != pos(pet, snails) {
							continue
						}
						// Kools are smoked in the yellow house
						if pos(smoke, kools) != pos(color, yellow) {
							continue
						}
						// The man who smokes Chesterfields lives in the house next to the man with the fox.
						if !nextTo(pos(smoke, chesterfields), pos(pet, fox)) {
							continue
						}
						// Kools are smoked in the house next to the house where the horse is kept.
						if !nextTo(pos(smoke, kools), pos(pet, horse)) {
							continue
						}
						// The Lucky Strike smoker drinks orange juice.
						if pos(smoke, luckyStrike) != pos(drink, orangeJuice) {
							continue
						}
						// The Japanese smokes Parliaments.
						if pos(nationality, japanese) != pos(smoke, parliaments) {
							continue
						}

						return Solution{
							DrinksWater: nationalities[nationality[pos(drink, water)]],
							OwnsZebra:   nationalities[nationality[pos(pet, zebra)]],
						}
					}
				}
			}
		}
	}

	return Solution{}
}

func permutations(arr []int) [][]int {
	var result [][]int
	var generate func(int)
	a := make([]int, len(arr))
	copy(a, arr)

	generate = func(n int) {
		if n == 1 {
			temp := make([]int, len(a))
			copy(temp, a)
			result = append(result, temp)
			return
		}
		for i := 0; i < n; i++ {
			generate(n - 1)
			if n%2 == 1 {
				a[0], a[n-1] = a[n-1], a[0]
			} else {
				a[i], a[n-1] = a[n-1], a[i]
			}
		}
	}
	generate(len(a))
	return result
}

func pos(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func nextTo(a, b int) bool {
	return a == b+1 || a == b-1
}
