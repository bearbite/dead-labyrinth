package dice

import "math/rand"

// Roll visszaad két véletlenszerű számot 1 és 6 között
func Roll() (int, int) {
	roll1 := rand.Intn(6) + 1
	roll2 := rand.Intn(6) + 1
	return roll1, roll2

}
