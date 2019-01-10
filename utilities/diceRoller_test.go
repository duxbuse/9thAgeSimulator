package logisticsmaps

import "testing"
import "math"

func TestRollDice(t *testing.T) {
	// Dice modeled are d6's so can only roll number 1-6
	for i := 0; i < 10000; i++ {
		value := RollDice()
		if value < 1 || value > 6 {
			t.Errorf("Rolled a dice with a value not [1-6]")
		}
	}
}

func Test7s(t *testing.T) {
	// Rolling 2 dice with no modifiers will sum to 7, 58% of the time
	expected := 58.00
	recieved := math.Floor(ChanceOfSuccess(7, true, false, 0, 0) * 100)
	if recieved != expected {
		t.Errorf("Calculated wrong Value, got: %.2f, want: %.2f.", recieved, expected)
	}
}

func TestChanceOfSuccess(t *testing.T) {
	var Tests = []struct {
		//inputs
		threshold int
		forward   bool
		reroll    bool
		min       int
		max       int

		expected float64 // expected result
	}{
		{2, true, false, 0, 0, 10000}, //Charge=2
		{2, true, true, 0, 0, 10000},  //Charge=2, reroll
		{2, true, false, 0, 1, 10000}, //Charge=2, maximised 1
		{2, true, true, 0, 1, 10000},  //Charge=2, maximised 1, reroll
		{2, true, false, 0, 2, 10000}, //Charge=2, maximised 2
		{2, true, true, 0, 2, 10000},  //Charge=2, maximised 2, reroll
		{2, true, false, 1, 0, 10000}, //Charge=2, minimised 1
		{2, true, true, 1, 0, 10000},  //Charge=2, minimised 1, reroll
		{2, true, false, 2, 0, 10000}, //Charge=2, minimised 2
		{2, true, true, 2, 0, 10000},  //Charge=2, minimised 2, reroll

		{3, true, false, 0, 0, 9722}, //Charge=3
		{3, true, true, 0, 0, 9994},  //Charge=3, reroll
		{3, true, false, 0, 1, 9954}, //Charge=3, maximised 1
		{3, true, true, 0, 1, 9999},  //Charge=3, maximised 1, reroll
		{3, true, false, 0, 2, 9992}, //Charge=3, maximised 2
		{3, true, true, 0, 2, 9999},  //Charge=3, maximised 2, reroll
		{3, true, false, 1, 0, 9259}, //Charge=3, minimised 1
		{3, true, true, 1, 0, 9945},  //Charge=3, minimised 1, reroll
		{3, true, false, 2, 0, 8681}, //Charge=3, minimised 2
		{3, true, true, 2, 0, 9826},  //Charge=3, minimised 2, reroll

		{7, true, false, 0, 0, 5833}, //Charge=7
		{7, true, true, 0, 0, 8263},  //Charge=7, reroll
		{7, true, false, 0, 1, 8046}, //Charge=7, maximised 1
		{7, true, true, 0, 1, 9622},  //Charge=7, maximised 1, reroll
		{7, true, false, 0, 2, 9097}, //Charge=7, maximised 2
		{7, true, true, 0, 2, 9918},  //Charge=7, maximised 2, reroll
		{7, true, false, 1, 0, 3194}, //Charge=7, minimised 1
		{7, true, true, 1, 0, 5367},  //Charge=7, minimised 1, reroll
		{7, true, false, 2, 0, 1736}, //Charge=7, minimised 2
		{7, true, true, 2, 0, 3171},  //Charge=7, minimised 2, reroll

		{9, true, false, 0, 0, 2778}, //Charge=9
		{9, true, true, 0, 0, 4784},  //Charge=9, reroll
		{9, true, false, 0, 1, 5231}, //Charge=9, maximised 1
		{9, true, true, 0, 1, 7725},  //Charge=9, maximised 1, reroll
		{9, true, false, 0, 2, 6944}, //Charge=9, maximised 2
		{9, true, true, 0, 2, 9066},  //Charge=9, maximised 2, reroll
		{9, true, false, 1, 0, 1065}, //Charge=9, minimised 1
		{9, true, true, 1, 0, 2016},  //Charge=9, minimised 1, reroll
		{9, true, false, 2, 0, 401},  //Charge=9, minimised 2
		{9, true, true, 2, 0, 786},   //Charge=9, minimised 2, reroll

		{13, true, false, 0, 0, 0}} //blank case

	for _, tt := range Tests {
		//test to the nearest percent
		actual := math.Round(ChanceOfSuccess(tt.threshold, tt.forward, tt.reroll, tt.min, tt.max) * 100)
		if actual != math.Round(tt.expected/100) {
			t.Errorf("\nChanceOfSuccess(%d, %t, %t, %d, %d):\nexpected %.0f\nactual   %.0f", tt.threshold, tt.forward, tt.reroll, tt.min, tt.max, math.Round(tt.expected/100), actual)
		}
	}
}
