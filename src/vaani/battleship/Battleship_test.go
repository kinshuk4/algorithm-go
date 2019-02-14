package battleship_test

import (
	"testing"
	"fmt"
	"vaani/codility/battleship"
)

func TestBinaryGap(t *testing.T) {
	actual := battleship.Solution(4, "1B 2C,2D 4D", "2B 2D 3D 4D 4A")
	fmt.Println(actual)
	//actual := lesson01.BinaryGap(1041)
	//expected := 5
	//if actual != expected {
	//	t.Errorf("BinaryGap failed, expected %d, got %d", expected, actual)
	//}
	//
	//actual = lesson01.BinaryGap(8)
	//expected = 0
	//if actual != expected {
	//	t.Errorf("BinaryGap failed, expected %d, got %d", expected, actual)
	//}
	//
	//actual = lesson01.BinaryGap(9)
	//expected = 2
	//if actual != expected {
	//	t.Errorf("BinaryGap failed, expected %d, got %d", expected, actual)
	//}
}

