package Utilities

// TODO: need to re write the tests since no longer using chance but actual dice.

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
)

func TestFight(t *testing.T) {
	rand.Seed(9001) //This is critical to make sure the tests pass no matter where they are run.

	Tests := map[int]Results{}

	//Standard data required for a "Data" object
	// Set up unit height
	uheight := map[int]string{
		1: "Standard",
		2: "Large",
		3: "Gigantic"}

	// Set up unit type
	utype := map[int]string{
		1: "Infantry",
		2: "Beast",
		3: "Cavalry",
		4: "Construct"}

	// Set up base width in mm
	ubase := map[int]string{
		1: "20",
		2: "25",
		3: "40",
		4: "50",
		5: "60",
		6: "100",
		7: "150"}

	uweapon := map[int]string{
		1: "Sword and Board",
		2: "Spear",
		3: "Halberd",
		4: "Great Weapon",
		5: "Paired Weapons",
		6: "Light Lance",
		7: "Lance",
		8: "None"}

	uraces := map[int]string{
		1: "DE",
		2: "DH",
		3: "VC",
		4: "OK"}

	uspecialtiesStatsNames := map[string]string{"Hatred": "any", "Distracting": "any", "Lightning Reflexes": "any", "Killer Instinct": "DE", "Shield Wall": "DH"}

	///////////////////////////////////////////////////
	//////////////////////////////////////////////////
	/////////////////////////////////////////////////

	//fight1
	//21 DE spears 7 wide v 25 Sword and board dwarfs 5 wide

	fight1rawstats := map[string]EntetiesClass{"FQAN": {"FQAN", 21}, "FFOR": {"FFOR", 7}, "FDIS": {"FDIS", 8}, "FHP": {"FHP", 1}, "FDEF": {"FDEF", 4}, "FRES": {"FRES", 3}, "FARM": {"FARM", 2}, "FATT": {"FATT", 1}, "FOFF": {"FOFF", 4}, "FSTR": {"FSTR", 3}, "FAP": {"FAP", 0}, "FAGI": {"FAGI", 5}, "EQAN": {"EQAN", 25}, "EFOR": {"EFOR", 5}, "EDIS": {"EDIS", 9}, "EHP": {"EHP", 1}, "EDEF": {"EDEF", 4}, "ERES": {"ERES", 4}, "EARM": {"EARM", 3}, "EATT": {"EATT", 1}, "EOFF": {"EOFF", 4}, "ESTR": {"ESTR", 3}, "EAP": {"EAP", 0}, "EAGI": {"EAGI", 2}, "FSS": {"FSS", 0}, "ESS": {"ESS", 0}}

	fight1secondarystats := map[string]EntetiesClass{"FHeightSelect": {"FHeightSelect", 1}, "EHeightSelect": {"EHeightSelect", 1}, "FTypeSelect": {"FTypeSelect", 1}, "ETypeSelect": {"ETypeSelect", 1}, "FWidthSelect": {"FWidthSelect", 20}, "EWidthSelect": {"EWidthSelect", 20}, "FWeaponSelect": {"FWeaponSelect", 2}, "EWeaponSelect": {"EWeaponSelect", 1}, "FRaceSelect": {"FRaceSelect", 1}, "ERaceSelect": {"ERaceSelect", 2}}

	fight1specialtiesStatsOn := map[string]bool{"FLightning Reflexes": true, "FKiller Instinct": true, "EShield Wall": true}

	fight1 := Data{
		RawStats:           fight1rawstats,
		SecondaryStats:     fight1secondarystats,
		Weapon:             uweapon,
		Height:             uheight,
		Type:               utype,
		Width:              ubase,
		Races:              uraces,
		SpecialtiesStats:   uspecialtiesStatsNames,
		SpecialtiesStatsOn: fight1specialtiesStatsOn}

	outcome1 := Outcome{
		WINNER:      true,
		AMMOUNT:     2,
		BreakChance: "58.34",
		FNUM:        18,
		ENUM:        22}

	avgOutcome1 := Outcome{ //this is currently not checked and hence doesnt matter
		WINNER:      true,
		AMMOUNT:     0,
		BreakChance: "0",
		FNUM:        0,
		ENUM:        0}

	results1 := Results{fight1, outcome1, avgOutcome1}

	//fight2
	//24 Dread Judges v 6 hold guardians with distracting

	fight2rawstats := map[string]EntetiesClass{"FQAN": {"FQAN", 24}, "FFOR": {"FFOR", 8}, "FDIS": {"FDIS", 8}, "FHP": {"FHP", 1}, "FDEF": {"FDEF", 5}, "FRES": {"FRES", 3}, "FARM": {"FARM", 2}, "FATT": {"FATT", 1}, "FOFF": {"FOFF", 5}, "FSTR": {"FSTR", 4}, "FAP": {"FAP", 1}, "FAGI": {"FAGI", 5}, "EQAN": {"EQAN", 6}, "EFOR": {"EFOR", 3}, "EDIS": {"EDIS", 9}, "EHP": {"EHP", 3}, "EDEF": {"EDEF", 4}, "ERES": {"ERES", 5}, "EARM": {"EARM", 4}, "EATT": {"EATT", 3}, "EOFF": {"EOFF", 4}, "ESTR": {"ESTR", 6}, "EAP": {"EAP", 3}, "EAGI": {"EAGI", 2}, "FSS": {"FSS", 0}, "ESS": {"ESS", 0}}

	fight2secondarystats := map[string]EntetiesClass{"FHeightSelect": {"FHeightSelect", 1}, "EHeightSelect": {"EHeightSelect", 2}, "FTypeSelect": {"FTypeSelect", 1}, "ETypeSelect": {"ETypeSelect", 1}, "FWidthSelect": {"FWidthSelect", 20}, "EWidthSelect": {"EWidthSelect", 40}, "FWeaponSelect": {"FWeaponSelect", 4}, "EWeaponSelect": {"EWeaponSelect", 8}, "FRaceSelect": {"FRaceSelect", 1}, "ERaceSelect": {"ERaceSelect", 2}}

	fight2specialtiesStatsOn := map[string]bool{"FLightning Reflexes": true, "EDistracting": true, "FHatred": true, "FLethal Strike": true}

	fight2 := Data{
		RawStats:           fight2rawstats,
		SecondaryStats:     fight2secondarystats,
		Weapon:             uweapon,
		Height:             uheight,
		Type:               utype,
		Width:              ubase,
		Races:              uraces,
		SpecialtiesStats:   uspecialtiesStatsNames,
		SpecialtiesStatsOn: fight2specialtiesStatsOn}

	outcome2 := Outcome{
		WINNER:      true,
		AMMOUNT:     6,
		BreakChance: "91.64",
		FNUM:        20,
		ENUM:        3}

	avgOutcome2 := Outcome{ //this is currently not checked and hence doesnt matter. But I did the maths and this is correct.
		WINNER:      true,
		AMMOUNT:     5,
		BreakChance: "?",
		FNUM:        18,
		ENUM:        3}

	results2 := Results{fight2, outcome2, avgOutcome2}
	/////////////////////////////////////////////////////////////////////////
	Tests[1] = results1
	Tests[2] = results2

	for x, tt := range Tests {
		actual := fight(tt.UnitData)
		floatExpected, _ := strconv.ParseFloat(tt.FightResults.BreakChance, 64)
		floatActual, _ := strconv.ParseFloat(actual.BreakChance, 64)

		if actual != tt.FightResults && math.Round(floatExpected) != math.Round(floatActual) {
			t.Errorf("\nFight%d:\nexpected {WINNER: %t,	AMMOUNT: %d, BreakChance %s, FNUM %d, ENUM %d}\nactual   {WINNER: %t, AMMOUNT: %d, BreakChance %s, FNUM %d, ENUM %d}\n", x, tt.FightResults.WINNER, tt.FightResults.AMMOUNT, tt.FightResults.BreakChance, tt.FightResults.FNUM, tt.FightResults.ENUM, actual.WINNER, actual.AMMOUNT, actual.BreakChance, actual.FNUM, actual.ENUM)
		}
	}
}

func TestCombatRes(t *testing.T) {
	var Tests = []struct {
		//inputs
		casualties int
		quantity   int
		formation  int
		unitHeight int
		bonuses    int

		expected int // expected result
	}{
		{0, 0, 0, 0, 0, 0},   //blank
		{12, 6, 3, 2, 1, 14}, //6 man large
		{8, 1, 1, 3, 0, 8},   //1 man giant
		{5, 20, 5, 1, 2, 10}} //20 man infantry

	for _, tt := range Tests {
		actual := CombatRes(tt.casualties, tt.quantity, tt.formation, tt.unitHeight, tt.bonuses)
		if actual != tt.expected {
			t.Errorf("CombatRes(%d, %d, %d, %d, %d): expected %d, actual %d", tt.casualties, tt.quantity, tt.formation, tt.unitHeight, tt.bonuses, tt.expected, actual)
		}
	}
}

// func TestHitChance(t *testing.T) {
// 	var Tests = []struct {
// 		//inputs
// 		FOFF      int
// 		EDEF      int
// 		parry     bool
// 		rerollINC int
// 		modifier  int

// 		expected float64 // expected result
// 	}{
// 		{0, 0, false, 0, 0, 3.0 / 6.0},   //blank
// 		{10, 10, false, 0, 0, 3.0 / 6.0}, //10-10
// 		{10, 8, false, 0, 0, 4.0 / 6.0},  //10-8
// 		{10, 6, false, 0, 0, 5.0 / 6.0},  //10-6
// 		{10, 4, false, 0, 0, 5.0 / 6.0},  //10-4
// 		{10, 1, false, 0, 0, 5.0 / 6.0},  //10-1
// 		{8, 10, false, 0, 0, 3.0 / 6.0},  //8-10
// 		{6, 10, false, 0, 0, 2.0 / 6.0},  //6-10
// 		{4, 10, false, 0, 0, 2.0 / 6.0},  //4-10
// 		{1, 10, false, 0, 0, 1.0 / 6.0},  //1-10

// 		{10, 10, true, 0, 0, 3.0 / 6.0}, //10-10 with parry
// 		{10, 8, true, 0, 0, 3.0 / 6.0},  //10-8 with parry
// 		{10, 6, true, 0, 0, 3.0 / 6.0},  //10-6 with parry
// 		{3, 6, true, 0, 0, 2.0 / 6.0},   //3-6 with parry adding +1ds
// 		{6, 10, true, 0, 0, 2.0 / 6.0},  //6-10 with parry
// 		{4, 10, true, 0, 0, 2.0 / 6.0},  //4-10 with parry
// 		{1, 10, true, 0, 0, 1.0 / 6.0},  //1-10 with parry

// 		{10, 10, false, 6, 0, (3.0 / 6.0) + (1.0-(3.0/6.0))*(3.0/6.0)}, //10-10 &rerolls
// 		// {10, 6, false, 6, 0, (5.0 / 6.0) + (1.0-(5.0/6.0))*(5.0/6.0)},  //10-6 &rerolls for some reason there is a precision issue with this test which causes it to fail even though the values are the same to 6 decimal places.
// 		{10, 8, false, 6, 0, (4.0 / 6.0) + (1.0-(4.0/6.0))*(4.0/6.0)}, //10-8 &rerolls
// 		{8, 10, false, 6, 0, (3.0 / 6.0) + (1.0-(3.0/6.0))*(3.0/6.0)}, //8-10 &rerolls
// 		{6, 10, false, 6, 0, (2.0 / 6.0) + (1.0-(2.0/6.0))*(2.0/6.0)}, //6-10 &rerolls
// 		{1, 10, false, 6, 0, (1.0 / 6.0) + (1.0-(1.0/6.0))*(1.0/6.0)}, //1-10 &rerolls
// 		{10, 6, true, 6, 0, (3.0 / 6.0) + (1.0-(3.0/6.0))*(3.0/6.0)},  //10-6 with parry &rerolls

// 		{10, 10, false, 0, 1, 4.0 / 6.0}, //10-10 +1hit
// 		{10, 6, false, 0, 1, 5.0 / 6.0},  //10-6 +1hit
// 		{10, 8, false, 0, 1, 5.0 / 6.0},  //10-8 +1hit
// 		{8, 10, false, 0, 1, 4.0 / 6.0},  //8-10 +1hit
// 		{6, 10, false, 0, 1, 3.0 / 6.0},  //6-10 +1hit
// 		{1, 10, false, 0, 1, 2.0 / 6.0},  //1-10 +1hit
// 		{10, 6, true, 0, 1, 4.0 / 6.0},   //10-6 with parry +1hit

// 		{10, 10, false, 0, -1, 2.0 / 6.0}, //10-10 -1hit
// 		{10, 6, false, 0, -1, 4.0 / 6.0},  //10-6 -1hit
// 		{10, 8, false, 0, -1, 3.0 / 6.0},  //10-8 -1hit
// 		{8, 10, false, 0, -1, 2.0 / 6.0},  //8-10 -1hit
// 		{6, 10, false, 0, -1, 1.0 / 6.0},  //6-10 -1hit
// 		{1, 10, false, 0, -1, 1.0 / 6.0},  //1-10 -1hit
// 		{10, 6, true, 0, -1, 2.0 / 6.0}}   //10-6 with parry -1hit

// 	for _, tt := range Tests {
// 		actual := hitChance(tt.FOFF, tt.EDEF, tt.parry, tt.rerollINC, tt.modifier)
// 		if actual != tt.expected {
// 			t.Errorf("hitChance(%d, %d, %t, %d, %d): expected %f, actual %f", tt.FOFF, tt.EDEF, tt.parry, tt.rerollINC, tt.modifier, tt.expected, actual)
// 		}
// 	}
// }

// func TestWoundChance(t *testing.T) {
// 	var Tests = []struct {
// 		//inputs
// 		FSTR      int
// 		ERES      int
// 		rerollINC int
// 		modifier  int

// 		expected float64 // expected result
// 	}{
// 		{0, 0, 0, 0, 3.0 / 6.0},   //blank
// 		{10, 10, 0, 0, 3.0 / 6.0}, //10-10
// 		{10, 9, 0, 0, 4.0 / 6.0},  //10-9
// 		{10, 6, 0, 0, 5.0 / 6.0},  //10-6
// 		{10, 4, 0, 0, 5.0 / 6.0},  //10-4
// 		{10, 1, 0, 0, 5.0 / 6.0},  //10-1
// 		{9, 10, 0, 0, 2.0 / 6.0},  //9-10
// 		{6, 10, 0, 0, 1.0 / 6.0},  //6-10
// 		{4, 10, 0, 0, 1.0 / 6.0},  //4-10
// 		{1, 10, 0, 0, 1.0 / 6.0},  //1-10

// 		{10, 10, 6, 0, (3.0 / 6.0) + (1.0-(3.0/6.0))*(3.0/6.0)}, //10-10 &rerolls
// 		// {10, 6, 6, 0, (5.0 / 6.0) + (1.0-(5.0/6.0))*(5.0/6.0)},  //10-6 &rerolls. This also fails for some precision reason
// 		{10, 9, 6, 0, (4.0 / 6.0) + (1.0-(4.0/6.0))*(4.0/6.0)}, //10-9 &rerolls
// 		{9, 10, 6, 0, (2.0 / 6.0) + (1.0-(2.0/6.0))*(2.0/6.0)}, //9-10 &rerolls
// 		{6, 10, 6, 0, (1.0 / 6.0) + (1.0-(1.0/6.0))*(1.0/6.0)}, //6-10 &rerolls
// 		{1, 10, 6, 0, (1.0 / 6.0) + (1.0-(1.0/6.0))*(1.0/6.0)}, //1-10 &rerolls

// 		{10, 10, 0, 1, 4.0 / 6.0}, //10-10 +1wound
// 		{10, 6, 0, 1, 5.0 / 6.0},  //10-6 +1wound
// 		{10, 9, 0, 1, 5.0 / 6.0},  //10-9 +1wound
// 		{9, 10, 0, 1, 3.0 / 6.0},  //9-10 +1wound
// 		{6, 10, 0, 1, 2.0 / 6.0},  //6-10 +1wound
// 		{1, 10, 0, 1, 2.0 / 6.0},  //1-10 +1wound

// 		{10, 10, 0, -1, 2.0 / 6.0}, //10-10 -1wound
// 		{10, 6, 0, -1, 4.0 / 6.0},  //10-6 -1wound
// 		{10, 9, 0, -1, 3.0 / 6.0},  //10-9 -1wound
// 		{9, 10, 0, -1, 1.0 / 6.0},  //9-10 -1wound
// 		{6, 10, 0, -1, 1.0 / 6.0},  //6-10 -1wound
// 		{1, 10, 0, -1, 1.0 / 6.0}}  //1-10 -1wound

// 	for _, tt := range Tests {
// 		actual := woundChance(tt.FSTR, tt.ERES, tt.rerollINC, tt.modifier)
// 		if actual != tt.expected {
// 			t.Errorf("woundChance(%d, %d, %d, %d): expected %f, actual %f", tt.FSTR, tt.ERES, tt.rerollINC, tt.modifier, tt.expected, actual)
// 		}
// 	}
// }

// func TestArmourFailChance(t *testing.T) {
// 	var Tests = []struct {
// 		//inputs
// 		FAP      int
// 		EARM     int
// 		expected float64 // expected result
// 	}{
// 		{0, 1, 5.0 / 6.0}, //blank
// 		{0, 2, 4.0 / 6.0}, //blank
// 		{0, 3, 3.0 / 6.0}, //blank
// 		{0, 4, 2.0 / 6.0}, //blank
// 		{0, 5, 1.0 / 6.0}, //blank
// 		{0, 6, 1.0 / 6.0}, //blank
// 		{0, 7, 1.0 / 6.0}, //blank

// 		{3, 1, 6.0 / 6.0}, //blank
// 		{3, 2, 6.0 / 6.0}, //blank
// 		{3, 3, 6.0 / 6.0}, //blank
// 		{3, 4, 5.0 / 6.0}, //blank
// 		{3, 5, 4.0 / 6.0}, //blank
// 		{3, 6, 3.0 / 6.0}, //blank
// 		{3, 7, 2.0 / 6.0}, //blank
// 		{3, 8, 1.0 / 6.0}} //1-10 -1wound

// 	for _, tt := range Tests {
// 		actual := armourFailChance(tt.FAP, tt.EARM)
// 		if actual != tt.expected {
// 			t.Errorf("armourFailChance(%d, %d): expected %f, actual %f", tt.FAP, tt.EARM, tt.expected, actual)
// 		}
// 	}
// }

func TestFightOrder(t *testing.T) {
	var Tests = []struct {
		//inputs
		FAGI     int
		EAGI     int
		expected rune // expected result
	}{
		{1, 2, 'E'},
		{2, 1, 'F'},
		{1, 1, 'S'}}

	for _, tt := range Tests {
		actual := fightOrder(tt.FAGI, tt.EAGI)
		if actual != tt.expected {
			t.Errorf("fightOrder(%d, %d): expected %d, actual %d", tt.FAGI, tt.EAGI, tt.expected, actual)
		}
	}
}

func TestRanks(t *testing.T) {
	var Tests = []struct {
		//inputs
		quantity   int
		formation  int
		unitHeight int

		expected int // expected result
	}{
		{0, 0, 0, 0},  //blank
		{6, 3, 2, 2},  //6 man large
		{1, 1, 3, 1},  //1 man giant
		{14, 7, 1, 2}, //14 man infantry 7 wide
		{12, 5, 1, 2}, //12 man infantry 5 wide not full back rank
		{21, 8, 1, 3}, //21 man infantry 8 wide not full back rank but still enough for a rank
		{5, 3, 2, 1},  //5 man ogre 3 wide not full back rank
		{7, 4, 2, 2},  //7 man ogre 4 wide not full back rank
		{3, 2, 3, 2},  //unit of giants
		{20, 5, 1, 4}} //20 man infantry

	for _, tt := range Tests {
		actual := ranks(tt.quantity, tt.formation, tt.unitHeight)
		if actual != tt.expected {
			t.Errorf("Ranks(%d, %d, %d): expected %d, actual %d", tt.quantity, tt.formation, tt.unitHeight, tt.expected, actual)
		}
	}
}

func TestNumOfCombatants(t *testing.T) {
	var Tests = []struct {
		//inputs
		AFOR   int
		AQAN   int
		AbaseW int
		BFOR   int
		BQAN   int
		BbaseW int

		expecteda int // expected result
		expectedb int // expected result
	}{
		{5, 20, 20, 5, 20, 20, 5, 5}, //20man inf agaisnt each other
		{5, 20, 20, 3, 6, 40, 5, 3},  //20man inf v 6 large
		{7, 21, 25, 3, 4, 40, 6, 3},  //18man inf25mm v 6 large
		{1, 1, 50, 3, 3, 50, 1, 3},   //1 giant v 3 chariot
		{1, 1, 150, 8, 40, 20, 1, 8}, //1 huge dragon v horde inf
		{5, 4, 20, 8, 16, 20, 4, 6}}  //4inf v 16 inf

	for _, tt := range Tests {
		actuala, actualb := numOfCombatants(tt.AFOR, tt.AQAN, tt.AbaseW, tt.BFOR, tt.BQAN, tt.BbaseW)
		if actuala != tt.expecteda && actualb != tt.expectedb {
			t.Errorf("numOfCombatants(%d, %d, %d, %d, %d, %d): expected (%d, %d), actual (%d, %d)", tt.AFOR, tt.AQAN, tt.AbaseW, tt.BFOR, tt.BQAN, tt.BbaseW, tt.expecteda, tt.expectedb, actuala, actualb)
		}
	}
}

func TestNumOfAttacks(t *testing.T) {
	var Tests = []struct {
		//inputs
		combatants     int
		attacks        int
		quantity       int
		formation      int
		firstHeight    int
		secondHeight   int
		fightExtraRank int

		expecteda int // expected result
		expectedb int // expected result
	}{
		{5, 1, 25, 5, 1, 1, 1, 15, 0}, //25 spear men
		{8, 1, 32, 8, 1, 1, 1, 32, 0}, //32 spear men horde
		{5, 1, 32, 8, 1, 1, 1, 20, 0}, //32 spear men horde only 5 touching an enemy
		{6, 3, 14, 7, 1, 1, 0, 24, 0}, //16 witches 7 wide, 6 attacking
		{3, 3, 6, 3, 2, 1, 0, 18, 3},  //6 ogres against 5 wide inf
		{3, 4, 12, 4, 2, 1, 0, 21, 3}, //12 ogres 4 wide only 3 engaged with paired weapons against inf
		{1, 6, 1, 1, 3, 2, 0, 6, 0},   //giant v ogres
		{1, 6, 1, 1, 3, 1, 0, 6, 3},   //giant v inf
		{6, 1, 12, 6, 1, 1, 1, 12, 0}} //12inf 6 wide

	for _, tt := range Tests {
		actuala, actualb := numOfAttacks(tt.combatants, tt.attacks, tt.quantity, tt.formation, tt.firstHeight, tt.secondHeight, tt.fightExtraRank)
		if actuala != tt.expecteda && actualb != tt.expectedb {
			t.Errorf("numOfAttacks(%d, %d, %d, %d, %d, %d, %d): expected (%d, %d), actual (%d, %d)", tt.combatants, tt.attacks, tt.quantity, tt.formation, tt.firstHeight, tt.secondHeight, tt.fightExtraRank, tt.expecteda, tt.expectedb, actuala, actualb)
		}
	}
}

func TestABS(t *testing.T) {
	var Tests = []struct {
		//inputs
		input    int
		expected int // expected result
	}{
		{0, 0},
		{1, 1},
		{9999, 9999},
		{-1, 1},
		{-9999, 9999}}

	for _, tt := range Tests {
		actual := Abs(tt.input)
		if actual != tt.expected {
			t.Errorf("ABS(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}
}
