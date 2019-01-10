package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/duxbuse/LogisticsMaps"
)


func diceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving DiceRoller Page\n")
	logisticsmaps.RenderDiceRoller(w, r, "dice")
}

//Dummy page to use for testing
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Hello World")

	fmt.Fprintf(w, "Serving Hello World Page\n For other pages please go to /dice or /clasher")
}

func clasherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Clasher Page\n")
	logisticsmaps.RenderClasher(w, r, "clasher")
}

func main() {
	port := 9000
	http.HandleFunc("/dice/", diceHandler)
	http.HandleFunc("/clasher/", clasherHandler)
	http.HandleFunc("/", handler)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./../public/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./../public/css"))))

	fmt.Printf("Listening on Port: %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
