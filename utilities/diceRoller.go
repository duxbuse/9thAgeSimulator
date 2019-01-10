package Utilities

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
)

type DicePage struct {
	Result    string
	Threshold int
	Forward   bool
	Reroll    bool
	Min       int
	Max       int
}

func RenderDiceRoller(w http.ResponseWriter, r *http.Request, tmpl string) {

	fmt.Print("Serving Dice Page\n")
	threshold, _ := strconv.Atoi(r.FormValue("threshold"))
	// Ensure threshold is bounded by 2d6
	if threshold < 2 {
		threshold = 2
	} else if threshold > 12 {
		threshold = 12
	}

	forward, _ := strconv.ParseBool(r.FormValue("direction"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	max, _ := strconv.Atoi(r.FormValue("max"))
	reroll, _ := strconv.ParseBool(r.FormValue("reroll"))

	result := ChanceOfSuccess(threshold, forward, reroll, min, max)
	resultString := strconv.FormatFloat(result*100, 'f', 2, 64)

	p := &DicePage{Result: resultString, Threshold: threshold, Forward: forward, Min: min, Max: max, Reroll: reroll}
	renderDiceTemplate(w, "dice", p)

}

func renderDiceTemplate(w http.ResponseWriter, tmpl string, p *DicePage) {
	t, err := template.ParseFiles("./../public/views/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
ChanceOfSuccess is a function for determining the chance of beating a certain value when rolling two dice. The second paramter is to determine if the user whishes to get under or over the threshold value. It also handles the concept of minimizing or maximising. Maximising is the act of rolling an additional dice and then discarding the lowest value for the total result. Minimising is the reverse. The last option is to allow rerolls in the evnt of a failed case.
*/
func ChanceOfSuccess(threshold int, over bool, reroll bool, min int, max int) float64 {
	iterations := 1000000
	successes := 0

	for i := 0; i < iterations; i++ {
		resultsCount := 2 + min + max
		sum := generateresults(resultsCount, max, min)

		if over { //get higher than the value
			if sum >= threshold {
				successes++
			} else if reroll {
				// reroll as it failed
				sum = generateresults(resultsCount, max, min)
				if sum >= threshold {
					successes++
				}
			}
		} else { //get under the value
			if sum <= threshold {
				successes++
			} else if reroll {
				// reroll as it failed
				sum = generateresults(resultsCount, max, min)
				if sum <= threshold {
					successes++
				}
			}
		}
	}
	chanceOfSuccess := float64(successes) / float64(iterations)
	return chanceOfSuccess
}

func generateresults(resultsCount int, max int, min int) int {
	results := make([]int, resultsCount)

	//Roll all the dice
	for index := range results {
		results[index] = RollDice()
	}

	//Sort the dice to make it easy to remove min and max values
	sort.Ints(results)

	//sum results of dice you dont discard
	sum := results[max] + results[len(results)-1-min]

	return sum
}

/*
RollDice is a function representing rolling a single dice with 6 sides. returning a random interger between 1 and 6.
*/
func RollDice() int {
	return rand.Intn(6) + 1
}
