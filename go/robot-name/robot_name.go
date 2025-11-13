package robotname

import (
	"errors"
	"fmt"
)

type Robot struct {
	name string
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var increment = 0

const maximum = 26 * 26 * 1000

func generateName() (string, error) {
	if increment >= maximum {
		return "", errors.New("namespace is exhausted")
	}

	num := increment % 1000
	remaining := increment / 1000
	secondLetter := remaining % 26
	firstLetter := remaining / 26

	name := string(letters[firstLetter]) +
		string(letters[secondLetter]) +
		fmt.Sprintf("%03d", num)

	increment++
	return name, nil
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateName()
		if err != nil {
			return name, err
		}

		r.name = name
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
