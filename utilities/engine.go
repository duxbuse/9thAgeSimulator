package logisticsmaps

import (
	"math"
	"strconv"
)

func fight(data Data) Outcome {
	// TODO: Charging friend/enemyAGI++
	friendAGI := data.RawStats["FAGI"].Value
	friendLR := data.SpecialtiesStatsOn["FLightning Reflexes"]
	enemyAGI := data.RawStats["EAGI"].Value
	enemyLR := data.SpecialtiesStatsOn["FLightning Reflexes"]
	secondHitMod := 0
	firstHitMod := 0

	if data.SecondaryStats["FWeaponSelect"].Value == 4 { //greatweapon
		if !friendLR {
			friendAGI = 1
		}

	}
	if data.SecondaryStats["EWeaponSelect"].Value == 4 { //great weapon
		if !enemyLR {
			enemyAGI = 1
		}

	}
	//who fights first
	beforeOrder := fightOrder(friendAGI, enemyAGI)
	order := beforeOrder
	notOrder := order
	if beforeOrder == 'S' {
		order = 'F'
	}
	if order == 'F' {
		notOrder = 'E'
	} else { //order == 'E'
		notOrder = 'F'
	}

	//first fights first
	firstHeightSelection := data.SecondaryStats[string(order)+"HeightSelect"].Value
	firstBaseWidthSelection := data.SecondaryStats[string(order)+"WidthSelect"].Value
	firstBaseWidth, _ := strconv.Atoi(data.Width[firstBaseWidthSelection])
	firstFOR := data.RawStats[string(order)+"FOR"].Value
	firstQAN := data.RawStats[string(order)+"QAN"].Value
	firstHP := data.RawStats[string(order)+"HP"].Value
	firstATT := data.RawStats[string(order)+"ATT"].Value
	firstOFF := data.RawStats[string(order)+"OFF"].Value
	firstDEF := data.RawStats[string(order)+"DEF"].Value
	firstSTR := data.RawStats[string(order)+"STR"].Value
	firstRES := data.RawStats[string(order)+"RES"].Value
	firstARM := data.RawStats[string(order)+"ARM"].Value
	firstDIS := data.RawStats[string(order)+"DIS"].Value
	firstAP := data.RawStats[string(order)+"AP"].Value
	firstSS := data.RawStats[string(order)+"SS"].Value
	firstParry := false
	firstBSB := false
	firstHitReroll := 0
	firstWoundReroll := 0
	firstFIAR := 0
	firstHatred := data.SpecialtiesStatsOn[string(order)+"Hatred"]
	firstDistracting := data.SpecialtiesStatsOn[string(order)+"Distracting"]
	firstLightningReflexes := data.SpecialtiesStatsOn[string(order)+"Lightning Reflexes"]
	firstKillerInstinct := data.SpecialtiesStatsOn[string(order)+"Killer Instinct"]
	firstShieldWall := data.SpecialtiesStatsOn[string(order)+"Shield Wall"]

	//Make changes for firsts weapons
	switch data.SecondaryStats[string(order)+"WeaponSelect"].Value {
	case 1: //Sword and Board
		firstParry = true
	case 2: //Spear
		firstFIAR++
		firstAP++
	case 3: //Halberd
		firstSTR++
		firstAP++
	case 4: //Greatweapon
		firstSTR += 2
		firstAP += 2
		firstLightningReflexes = false
		//AGI allready handled
	case 5: //Paired Weapons
		firstOFF++
		firstATT++
	case 6: //Light Lance TODO: when we have charing this needs to be conditional
		firstSTR++
		firstAP++
	case 7: //Lance
		firstSTR += 2
		firstAP += 2
	case 8: //none
	}
	// Make changes for specialties
	if firstShieldWall && firstSS < 2 {
		//handle charging so its not always a 5++
		firstSS = 2
	}
	if firstLightningReflexes {
		firstHitMod++
	}
	if firstDistracting {
		secondHitMod--
	}
	if firstHatred {
		firstHitReroll = 6 //reroll upto all values
	}
	if firstKillerInstinct {
		firstWoundReroll = 1 //only reroll 1's
	}

	secondHeightSelection := data.SecondaryStats[string(notOrder)+"HeightSelect"].Value
	secondBaseWidthSelection := data.SecondaryStats[string(notOrder)+"WidthSelect"].Value
	secondBaseWidth, _ := strconv.Atoi(data.Width[secondBaseWidthSelection])
	secondFOR := data.RawStats[string(notOrder)+"FOR"].Value
	secondQAN := data.RawStats[string(notOrder)+"QAN"].Value
	secondHP := data.RawStats[string(notOrder)+"HP"].Value
	secondATT := data.RawStats[string(notOrder)+"ATT"].Value
	secondOFF := data.RawStats[string(notOrder)+"OFF"].Value
	secondDEF := data.RawStats[string(notOrder)+"DEF"].Value
	secondSTR := data.RawStats[string(notOrder)+"STR"].Value
	secondRES := data.RawStats[string(notOrder)+"RES"].Value
	secondARM := data.RawStats[string(notOrder)+"ARM"].Value
	secondDIS := data.RawStats[string(notOrder)+"DIS"].Value
	secondAP := data.RawStats[string(notOrder)+"AP"].Value
	secondSS := data.RawStats[string(notOrder)+"SS"].Value

	secondParry := false
	secondBSB := false
	secondHitReroll := 0
	secondWoundReroll := 0
	secondFIAR := 0
	secondHatred := data.SpecialtiesStatsOn[string(notOrder)+"Hatred"]
	secondDistracting := data.SpecialtiesStatsOn[string(notOrder)+"Distracting"]
	secondLightningReflexes := data.SpecialtiesStatsOn[string(notOrder)+"Lightning Reflexes"]
	secondKillerInstinct := data.SpecialtiesStatsOn[string(notOrder)+"Killer Instinct"]
	secondShieldWall := data.SpecialtiesStatsOn[string(notOrder)+"Shield Wall"]
	//Make changes for seconds weapons
	switch data.SecondaryStats[string(notOrder)+"WeaponSelect"].Value {
	case 1: //Sword and Board
		secondParry = true
	case 2: //Spear
		secondFIAR++
		secondAP++
		//TODO: if being charged increase ap
	case 3: //Halberd
		secondSTR++
		secondAP++
	case 4: //Greatweapon
		secondSTR += 2
		secondAP += 2
		secondLightningReflexes = false
		//AGI allready handled
	case 5: //Paired Weapons
		secondOFF++
		secondATT++
	case 6: //Light Lance
		secondSTR++
		secondAP++
	case 7: //Lance
		secondSTR += 2
		secondAP += 2
	case 8: //none
	}
	// Make changes for specailties
	if secondShieldWall && secondSS < 2 {
		//handle charging so its not always a 5++
		secondSS = 2
	}
	if secondLightningReflexes {
		secondHitMod++
	}
	if secondDistracting {
		firstHitMod--
	}
	if secondHatred {
		secondHitReroll = 6 //reroll upto all values
	}
	if secondKillerInstinct {
		secondWoundReroll = 1 //only reroll 1's
	}
	/////////////////////////////////////////////////////////////////
	// Calculate the results now all the stats are avaliable

	// Whos fighting
	firstCombatants, secondCombatants := numOfCombatants(firstFOR, firstQAN, firstBaseWidth, secondFOR, secondQAN, secondBaseWidth)

	firstAttacks, firstBonusHits := numOfAttacks(firstCombatants, firstATT, firstQAN, firstFOR, firstHeightSelection, secondHeightSelection, firstFIAR)

	firstHitChance := hitChance(firstOFF, secondDEF, secondParry, firstHitReroll, firstHitMod)

	firstWoundChance := woundChance(firstSTR, secondRES, 0, firstWoundReroll) //TODO: bring in modifiers properly
	firstArmourFailChance := armourFailChance(firstAP, secondARM)

	firstSpecialFailChance := armourFailChance(0, secondSS) //ap is always 0 for special saves

	firstCasualties := (firstAttacks*firstHitChance + firstBonusHits) * firstWoundChance * firstArmourFailChance * firstSpecialFailChance

	// Take off the casualties now if not simultaneous combat
	if !(beforeOrder == 'S') {
		secondQAN = secondQAN - int(math.Floor(firstCasualties/float64(secondHP)))

	}

	secondAttacks, secondBonusHits := numOfAttacks(secondCombatants, secondATT, secondQAN, secondFOR, secondHeightSelection, firstHeightSelection, secondFIAR)

	if secondLightningReflexes {
		secondHitMod++
	}
	secondHitChance := hitChance(secondOFF, firstDEF, firstParry, secondHitReroll, secondHitMod)
	secondWoundChance := woundChance(secondSTR, firstRES, secondWoundReroll, 0) //TODO: bring in modifiers and rerolls properly
	secondArmourFailChance := armourFailChance(secondAP, firstARM)
	secondSpecialFailChance := armourFailChance(0, firstSS)

	secondCasualties := (secondAttacks*secondHitChance + secondBonusHits) * secondWoundChance * secondArmourFailChance * secondSpecialFailChance

	// Take off the casualties
	firstQAN = firstQAN - int(math.Floor(secondCasualties/float64(firstHP)))
	if beforeOrder == 'S' { //and the inital casualties for simultaneous combat.
		secondQAN = secondQAN - int(math.Floor(firstCasualties/float64(secondHP)))
	}

	firstCombatRes := CombatRes(firstCasualties, firstQAN, firstFOR, firstHeightSelection, 0) //TODO: bonuses need work like having a banner or charging.
	secondCombatRes := CombatRes(secondCasualties, secondQAN, secondFOR, secondHeightSelection, 0)

	combatResSum := firstCombatRes - secondCombatRes
	firstRanks := ranks(firstQAN, firstFOR, firstHeightSelection)
	secondRanks := ranks(secondQAN, secondFOR, secondHeightSelection)
	breakchance := 0.0 //draw so no breaking
	firstWon := false
	if combatResSum > 0 {
		firstWon = true
		threshold := secondDIS - combatResSum //combatresSum is positive
		if secondRanks > firstRanks {         //steadfast
			threshold = secondDIS
		}
		breakchance = 1.0 - ChanceOfSuccess(threshold, false, secondBSB, 0, 0)
	} else if combatResSum < 0 {
		firstWon = false
		threshold := firstDIS + combatResSum //combatresSum is negative
		if firstRanks > secondRanks {        //steadfast
			threshold = firstDIS
		}
		breakchance = 1.0 - ChanceOfSuccess(threshold, false, firstBSB, 0, 0)
	}
	breakchanceString := "N/A"
	if breakchance != 0.0 {
		breakchanceString = strconv.FormatFloat(breakchance*100, 'f', 2, 64)
	}
	if order == 'F' {
		return Outcome{firstWon, Abs(combatResSum), breakchanceString, firstQAN, secondQAN}
	}
	//Reverse the units remaining since the enemy went first
	return Outcome{firstWon, Abs(combatResSum), breakchanceString, secondQAN, firstQAN}

}

func CombatRes(casualties float64, quantity int, formation int, unitHeight int, bonuses int) int {

	rankbonus := float64(ranks(quantity, formation, unitHeight) - 1) //first rank doesnt give a bonus
	if rankbonus < 0.0 {
		rankbonus = 0.0 //ensure you cant go negative
	}
	return int(math.Floor(math.Min(rankbonus, 3) + casualties + float64(bonuses)))
}
func hitChance(FOFF int, EDEF int, parry bool, rerollINC int, modifier int) float64 {
	// rerollINC represents the values up to reroll out of 6 to reroll. EG rerollINC =  1 only rerolls values of 1, rerollINC = 6 rerolls all values.
	//modifier can be +1 or -1 to represent hitting easier.
	//TODO: potentialy use parry as an int so that things like can never be hit on better than a x+ can be quasi parry. Since parry is essentially cant be hit on better than a 4+
	diff := FOFF - EDEF
	if parry && diff < 0 {
		diff-- //the enemy gets an extra point of ds with shield if normally higher
	}
	hit := 0.0
	if diff >= 4 {
		hit = 5.0
	} else if diff > 0 {
		hit = 4.0
	} else if diff >= -3 {
		hit = 3.0
	} else if diff >= -7 {
		hit = 2.0
	} else {
		hit = 1.0
	}
	if parry && hit > 3.0 {
		hit = 3.0
	}
	hit = math.Min(math.Max(hit+float64(modifier), 1.0), 5.0) //hit value out of 6 that will hit

	chance := hit / 6.0
	failedchance := (6.0 - hit) / 6.0
	rerollpercent := float64(rerollINC) / 6.0
	percentToReroll := math.Min(failedchance, rerollpercent)

	total := chance + percentToReroll*chance
	return total
}
func woundChance(FSTR int, ERES int, rerollINC int, modifier int) float64 {
	diff := FSTR - ERES
	wound := 0.0
	if diff >= 2 {
		wound = 5.0
	} else if diff >= 1 {
		wound = 4.0
	} else if diff == 0 {
		wound = 3.0
	} else if diff >= -1 {
		wound = 2.0
	} else {
		wound = 1.0
	}
	wound = math.Min(math.Max(wound+float64(modifier), 1.0), 5.0) //wound value out of 6 that will wound

	chance := wound / 6.0
	failedchance := (6.0 - wound) / 6.0
	rerollpercent := float64(rerollINC) / 6.0
	percentToReroll := math.Min(failedchance, rerollpercent)

	total := chance + percentToReroll*chance
	return total
}
func armourFailChance(FAP int, EARM int) float64 { //TODO: rerolls both failed and successfull
	chance := EARM - FAP
	if chance > 5 {
		chance = 5
	} else if chance < 0 {
		chance = 0
	}
	return (6 - float64(chance)) / 6
}
func fightOrder(FAGI int, EAGI int) rune {
	if EAGI > FAGI {
		return 'E' //Enemy first
	} else if EAGI < FAGI {
		return 'F' //Friend first
	}
	return 'S'
}
func numOfCombatants(AFOR int, AQAN int, AbaseW int, BFOR int, BQAN int, BbaseW int) (int, int) {
	numA := math.Min(float64(AFOR), float64(AQAN))
	numB := math.Min(float64(BFOR), float64(BQAN))
	unitAWidth := numA * float64(AbaseW)
	unitBWidth := numB * float64(BbaseW)

	if unitAWidth < unitBWidth {
		// A fights with all
		fightA := int(numA)
		// B fights with as few as can fit into A's width plus upto 1 on each end
		fightB := int(math.Min(math.Floor(unitAWidth/float64(BbaseW))+2.0, float64(BQAN)))
		return fightA, fightB
	} else if unitBWidth < unitAWidth {
		// B fights with all
		fightB := int(numB)
		// A fights with as few as can fit into B's width plus upto 1 on each end
		fightA := int(math.Min(math.Floor(unitBWidth/float64(AbaseW))+2.0, float64(AQAN)))
		return fightA, fightB
	}
	//ere body fights cause they are the same widths
	return int(numA), int(numB)
}
func ranks(quantity int, formation int, height int) int {
	ranks := (float64(quantity) / float64(formation))
	backRank := math.Mod(float64(quantity), float64(formation))
	fullranks := 0.0

	switch height {
	case 1: //"infantry"
		if formation >= 5 {
			fullranks = math.Floor(ranks)
		}
		if backRank >= 5 {
			fullranks++
		}
	case 2: //large
		if formation >= 3 {
			fullranks = math.Floor(ranks)
		}
		if backRank >= 3 {
			fullranks++
		}
	case 3: //gigantic
		if formation >= 1 {
			fullranks = math.Floor(ranks)
		}
		if backRank >= 1 {
			fullranks++
		}
	}
	return int(fullranks)
}

/*
numOfAttacks returns the number of attacks the unit will get to make, as well as how many bonus hits like stomps as the second return value.
*/
func numOfAttacks(combatants int, attacks int, quantity int, formation int, firstHeight int, secondHeight int, fightExtraRank int) (float64, float64) {
	//height 1 = standard
	//height 2 = large
	//height 3 = gigantic

	fightingRanks := 2
	if formation >= 8 {
		fightingRanks = 3
	}
	fightingRanks = fightingRanks + fightExtraRank

	frontRowAttacks := float64(attacks * combatants)

	maxSupportingAttacks := 1
	bonusHits := 0.0 //no bonus for regular height
	if firstHeight == 2 {
		maxSupportingAttacks = 3
		if secondHeight == 1 { //can only stomp standard
			bonusHits = float64(combatants) //1 stomp attack each
		}
	} else if firstHeight == 3 {
		maxSupportingAttacks = 5
		if secondHeight == 1 { //can only stomp standard
			bonusHits = float64(combatants) * 3.5 //should be a d6 but meh.
		}
	}
	//TODO: bonus hits +=(d6+1) for impact hits. Need to know if chariots etc. Reliant on if charging

	// min of max number of supporting attacks for guys engaged or every remaining model supporting
	supportingAttacks := math.Min(float64(combatants*(fightingRanks-1)*maxSupportingAttacks), float64((quantity-formation)*maxSupportingAttacks))

	return frontRowAttacks + supportingAttacks, bonusHits
}

/*
Abs returns the absoloute of 2 ints
*/
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
