package bowling

import "errors"

type Game struct {
	rolls    []int
	isSecond bool
}

func NewGame() *Game {
	return &Game{rolls: make([]int, 0, 21)}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return errors.New("invalid pin count")
	}

	if _, err := g.Score(); err == nil {
		return errors.New("game is over")
	}

	pinsDown := 0
	if g.isSecond {
		pinsDown = g.rolls[len(g.rolls)-1]
	}

	if pins+pinsDown > 10 {
		return errors.New("pin count exceeds pins on the lane")
	}

	g.rolls = append(g.rolls, pins)
	g.isSecond = !g.isSecond

	if pins == 10 {
		g.isSecond = false
	}

	return nil
}

func (g *Game) Score() (int, error) {
	total := 0
	rollIdx := 0

	for frame := 0; frame < 10; frame++ {
		if rollIdx+1 >= len(g.rolls) {
			return 0, errors.New("game not finished")
		}

		first := g.rolls[rollIdx]
		second := g.rolls[rollIdx+1]

		if first == 10 {
			if rollIdx+2 >= len(g.rolls) {
				return 0, errors.New("game not finished")
			}
			total += 10 + second + g.rolls[rollIdx+2]
			rollIdx++
		} else if first+second == 10 {
			if rollIdx+2 >= len(g.rolls) {
				return 0, errors.New("game not finished")
			}
			total += 10 + g.rolls[rollIdx+2]
			rollIdx += 2
		} else {
			total += first + second
			rollIdx += 2
		}
	}

	return total, nil
}
