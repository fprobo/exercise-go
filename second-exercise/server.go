package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func square(w http.ResponseWriter, req *http.Request) {
	numberStr := req.URL.Path[1:]

	if numberStr != "" {
		number, err := strconv.ParseFloat(numberStr, 64)

		if err != nil {
			fmt.Fprint(w, "Convert error!!")
		}

		fmt.Fprint(w, "Result: ", int(math.Pow(number, 2)))

	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/square", square)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
	// log.Println("Listening...")
}
