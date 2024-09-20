package tunnel

import (
	"fmt"
	"math/rand"
)

// RandTunnel működése:
// Vissza kell tudja adni, hogy mi van előre jobbra és balra
// Kell tudjon választani, hogy az adott irányba mi van.
// Lehet hogy újabb Tunnel van és lehet, hogy Room van.
// Front a right és a left csak és kizárólag Room vagy Tunnel értéket vehet fel.
// És minden hívásnál random dőljön el, hogy melyik irányba mit adunk vissza.

func RandTunnel() (front, right, left string) {

	/*frontRand := rand.Intn(2)
	if frontRand == 0 {
		front = "room"

	} else {
		front = "tunnel"
	}
	rightRand := rand.Intn(2)
	if rightRand == 0 {
		right = "room"

	} else {
		right = "tunnel"
	}
	leftRand := rand.Intn(2)
	if leftRand == 0 {
		left = "room"

	} else {
		left = "tunnel"
	}
	fmt.Println(front, right, left)
	return front, right, left*/

	randPG := []string{"room", "tunnel"}

	front = randPG[rand.Intn(2)]
	right = randPG[rand.Intn(2)]
	left = randPG[rand.Intn(2)]

	fmt.Println(front, right, left)
	return front, right, left
}
