package dndcharacter

import (
	"math"
	"math/rand"
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

func ThrowDice(sides int) int {
	return rand.Intn(sides) + 1
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum, minThrow := 0, 7
	for i := 1; i <= 4; i++ {
		throw := ThrowDice(6)
		sum += throw
		if throw < minThrow {
			minThrow = throw
		}
	}
	return sum - minThrow
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
