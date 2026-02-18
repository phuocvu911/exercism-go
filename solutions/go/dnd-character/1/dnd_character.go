package dndcharacter
import "math/rand"
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
	mid:= score-10
    if mid<0{
        mid-=1
    }
    return mid/2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	d1:= rand.Intn(6) + 1
	d2:= rand.Intn(6) + 1
	d3:= rand.Intn(6) + 1
	//d4:= rand.IntN(6) + 1
    return d1+d2+d3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    con := Ability()
    hit:= Modifier(con) + 10
	return Character{Ability(),Ability(),con,Ability(),Ability(),Ability(),hit}
}
