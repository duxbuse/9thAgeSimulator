package logisticsmaps

import (
	"html/template"
	"net/http"
	"strconv"
)

type EntetiesClass struct {
	Name  string
	Value int
}

/*
IsEnemy calculates which side it is on.
*/
func (ec EntetiesClass) IsEnemy(name string) bool {
	if string([]rune(name)[0]) == "E" {
		return true
	}
	return false
}

type Data struct {
	RawStats       map[string]EntetiesClass
	SecondaryStats map[string]EntetiesClass

	Weapon             map[int]string
	Height             map[int]string
	Type               map[int]string
	Width              map[int]string
	Races              map[int]string
	SpecialtiesStats   map[string]string
	SpecialtiesStatsOn map[string]bool
}
type Outcome struct {
	WINNER      bool
	AMMOUNT     int
	BreakChance string
	FNUM        int
	ENUM        int
}

type Results struct {
	UnitData     Data
	FightResults Outcome
}

/*
RenderClasher renders the page.
*/
func RenderClasher(w http.ResponseWriter, r *http.Request, tmpl string) {

	// Set up Unit Stats
	urawstats := map[string]EntetiesClass{}
	usecondarystats := map[string]EntetiesClass{}

	//number values
	rawstatsnames := []string{"FQAN", "FFOR", "FDIS", "FHP", "FDEF", "FRES", "FARM", "FATT", "FOFF", "FSTR", "FAP", "FAGI", "EQAN", "EFOR", "EDIS", "EHP", "EDEF", "ERES", "EARM", "EATT", "EOFF", "ESTR", "EAP", "EAGI", "FSS", "ESS"}
	for _, element := range rawstatsnames {
		value, _ := strconv.Atoi(r.FormValue(element))

		urawstats[element] = EntetiesClass{Name: element[1:], Value: value}
	}

	secondarystatsnames := []string{"FHeightSelect", "EHeightSelect", "FTypeSelect", "ETypeSelect", "FWidthSelect", "EWidthSelect", "FWeaponSelect", "EWeaponSelect", "FRaceSelect", "ERaceSelect"} //dropdown values
	for _, element := range secondarystatsnames {
		value, _ := strconv.Atoi(r.FormValue(element))

		usecondarystats[element] = EntetiesClass{Name: element[1:], Value: value}
	}
	//drop down lists
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

	races := map[int]string{
		1: "DE",
		2: "DH",
		3: "VC",
		4: "OK"}

	specialtiesStatsNames := map[string]string{"Hatred": "any", "Distracting": "any", "Lightning Reflexes": "any", "Killer Instinct": "DE", "Shield Wall": "DH"}

	specialtiesStatsOn := map[string]bool{}
	// TODO: Make the specialties impact the reuslts
	for k := range specialtiesStatsNames {
		specialtiesStatsOn["F"+k] = len(r.FormValue("F"+k)) > 0
		specialtiesStatsOn["E"+k] = len(r.FormValue("E"+k)) > 0
	}
	// Save all unit data as one object
	data := Data{RawStats: urawstats, SecondaryStats: usecondarystats, Weapon: uweapon, Height: uheight, Type: utype, Width: ubase, Races: races, SpecialtiesStats: specialtiesStatsNames, SpecialtiesStatsOn: specialtiesStatsOn}

	outcome := fight(data)

	payload := Results{UnitData: data, FightResults: outcome}

	// Begin templating
	t, err := template.ParseFiles("./../public/views/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
