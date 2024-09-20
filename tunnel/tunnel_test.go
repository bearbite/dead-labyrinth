package tunnel

import "testing"

func TestRandTunnel(t *testing.T) {

	front, right, left := RandTunnel()

	if front != "room" && front != "tunnel" {
		t.Error("a front vart erteke room vagy tunnel. Kapott erteke:", front)
	}
	if right != "room" && right != "tunnel" {
		t.Error("az right vart erteke room vagy tunnel. Kapott erteke:", right)
	}
	if left != "room" && left != "tunnel" {
		t.Error("az left vart erteke room vagy tunnel. Kapott erteke:", left)
	}
}
