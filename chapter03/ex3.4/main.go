package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jijiwhywhy/goplSolutions/chapter03/ex3.4/surface"
)

/*
Exercise 3.4: Following the approach of the Lissajous example in Section 1.7,
construct a web server that computes surfaces and writes SVG data to the client.
The server must set the Content-Type header like this: w.Header().Set("Content-Type", "image/svg+xml")
(This step was not required in the Lissajous example because the server uses standard heuristics to recognize
common formats like PNG from the ﬁrst 512 bytes of the response and generates the proper header.)
Allow the client to specify values like height, width, and color as HTTP request parameters.
*/

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

/*
func parseFirstIntOrDefault(array []string, defaultValue int) int {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.Atoi(array[0])
	if err != nil {
		return defaultValue
	}
	return value
}


func parseFirstFloat64OrDefault(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}
	return value
}


func parseFirstColorOrDefault(array []string, defaultValue color.Color) color.Color {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := colors.ColorFromString(array[0])
	if err != nil {
		return defaultValue
	}
	return value
}*/
