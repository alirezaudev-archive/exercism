package dndcharacter

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = rand.Intn(6) + 1
	}

	sort.Ints(a)
	return a[1] + a[2] + a[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{}
	values := reflect.ValueOf(&character).Elem()
	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)
		if field.CanSet() {
			if i == 6 {
				constitution := values.Field(2).Int()
				field.SetInt(10 + int64(Modifier(int(constitution))))
			} else {
				field.SetInt(int64(Ability()))
			}
		}
	}
	return character
}
