package atlas

import (
	"fmt"
	"github.com/bearbite/dead-labyrinth/dice"
	"github.com/bearbite/dead-labyrinth/monster"
	"github.com/bearbite/dead-labyrinth/tunnel"
)

func Atlas() {

	dice1, dice2 := dice.Roll()

	// dobokocka értékének összege
	randomSzam := dice1 + dice2

	fmt.Println("Ennyit dobott a user:", randomSzam)

	type Tunnel struct {
		Front string
		Right string
		Left  string
	}

	var tunnels []Tunnel

	// az I-hez először  hozzárendelünk egy  értéket.
	// A ciklus egészen addig fut, ameddign az i kisebb mint a randomSzam értéke.
	// és az i++ minden ciklusban egyel növeli az i értékét.
	for i := 0; i < randomSzam; i++ {

		// a front, right, left változók értékei a tunnel.RandTunnel() értékei lesznek.
		tunnelInterface := tunnel.New(1, 1)

		// a tunnels tömbhöz hozzáadjuk a front, right, left értékeit.
		tunnels = append(tunnels, Tunnel{
			Front: tunnelInterface.GetFront(),
			Right: tunnelInterface.GetRight(),
			Left:  tunnelInterface.GetLeft(),
		})

		/* tunnels[i] = Tunnel{
			Front: front,
			Right: right,
			Left:  left,
		} */

	}

	fmt.Println("A sliceba ezek kerültek", tunnels)

	for _, value := range tunnels {
		// add new string to the value
		value.Front = "Hello"
		value.Right = "Hello"
		value.Left = "Hello"
	}

	for _, value := range tunnels {
		fmt.Println(value)
	}

	monster1 := monster.New()
	monster1.SetName("Morgot")
	monster1.SetHealth(100)

	monster2 := monster.New()
	monster2.SetName("Gollum")
	monster2.SetHealth(50)

	dice1, dice2 = dice.Roll()
	newDiceValue := dice1 + dice2

	fmt.Println("A második szörny neve:", monster2.GetName())
	fmt.Println("A második szörny élete:", monster2.GetHealth())

	monster2.SetHealth(monster2.GetHealth() - newDiceValue)

	fmt.Println("A második szörny neve:", monster2.GetName())
	fmt.Println("A második szörny élete:", monster2.GetHealth())

}

/* func Minta() {

	enStringjeim := []string{"szia", "mia"}
	for i := 0; i < 10; i++ {
		enStringjeim = append(enStringjeim, "Hello")
	}

	fmt.Println(enStringjeim)

}
*/

/*
   [*]
   [*]
[*][-][*][*]
   [r][*][*][*][*]

*/
