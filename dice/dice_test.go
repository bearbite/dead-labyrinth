package dice

import "testing"

func TestRoll(t *testing.T) {

	dice1, dice2 := Roll()

	if dice1 < 1 || dice1 > 6 {
		t.Error("A dobokocka 1 elvárt értéke 1 és 6 között van, de a kapott érték:", dice1)
	}

	if dice2 < 1 || dice2 > 6 {
		t.Error("A dobokocka 2 elvárt értéke 1 és 6 között van, de a kapott érték:", dice2)
	}

}
