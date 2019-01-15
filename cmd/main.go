package main

import (
	"runtime/pprof"
	"os"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/duxbuse/utilities"
)

func diceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving DiceRoller Page\n")
	utilities.RenderDiceRoller(w, r, "dice")
}

//Dummy page to use for testing
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Hello World\n")

	fmt.Fprintf(w, "For other pages please go to /dice or /clasher")
}

func clasherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Clasher Page\n")
	utilities.RenderClasher(w, r, "clasher")
	
}
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {
	//Profiling
	flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }

	// Webserver
	port := 9000
	http.HandleFunc("/dice/", diceHandler)
	http.HandleFunc("/clasher/", clasherHandler)
	http.HandleFunc("/", handler)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./../public/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./../public/css"))))

	fmt.Printf("Listening on Port: %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
