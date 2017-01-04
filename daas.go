// - GET `/roll` generates a random roll from a D6 die.
// - GET `/roll?die=D<N>` (where N is a positive integer)
// - Both calls return JSON including the die type, and the random roll from the die.
// - Invalid parameters return a 400 and suitable JSON encoded error message.

package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Result struct {
	Roll int
	Type string
}

func diceServer(w http.ResponseWriter, req *http.Request) {
	diceType := req.FormValue("die")

	var numberLimit int
	var err error
	var json []byte

	if strings.HasPrefix(diceType, "D") {
		typeNumber := strings.Split(diceType, "D")[1]
		if numberLimit, err = strconv.Atoi(typeNumber); err == nil {
			json, err = generateJsonResult(numberLimit, diceType)
		}
	} else if diceType == "" {
		json, err = generateJsonResult(6, "D6")
	} else {
		err = errors.New("Error: Invalid die params")
	}

	if err == nil {
		w.Write(json)
	} else {
		http.Error(w, "Error: Invalid die params", 400)
	}
}

func generateJsonResult(limit int, diceType string) ([]byte, error) {
	roll := rand.Intn(int(limit)-1) + 1
	result := Result{roll, diceType}
	return json.Marshal(result)
}

func main() {
	http.HandleFunc("/roll", diceServer)
	http.ListenAndServe(":8080", nil)
}
