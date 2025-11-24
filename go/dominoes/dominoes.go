package dominoes

type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return input, true
	}

	if len(input) == 1 {
		return input, input[0][0] == input[0][1]
	}

	return buildChain(input, []Domino{})
}

func buildChain(remaining []Domino, chain []Domino) ([]Domino, bool) {
	if len(remaining) == 0 {
		return chain, chain[0][0] == chain[len(chain)-1][1]
	}

	var required int
	if len(chain) == 0 {
		required = -1
	} else {
		required = chain[len(chain)-1][1]
	}

	for i := 0; i < len(remaining); i++ {
		domino := remaining[i]
		if len(chain) == 0 || domino[0] == required {
			newChain := append(append([]Domino{}, chain...), domino)
			remained := removeDomino(remaining, i)
			if result, ok := buildChain(remained, newChain); ok {
				return result, true
			}
		}

		if len(chain) == 0 || domino[1] == required {
			reverse := Domino{domino[1], domino[0]}
			newChain := append(append([]Domino{}, chain...), reverse)
			remained := removeDomino(remaining, i)
			if result, ok := buildChain(remained, newChain); ok {
				return result, true
			}
		}
	}

	return nil, false
}

func removeDomino(dominoes []Domino, index int) []Domino {
	result := make([]Domino, 0, len(dominoes)-1)
	result = append(result, dominoes[:index]...)
	result = append(result, dominoes[index+1:]...)
	return result
}
