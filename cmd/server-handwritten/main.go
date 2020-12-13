package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/siuyin/dflt"
)

func main() {
	// 10_OMIT
	fmt.Println("Hand-written sum server")

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf("%v\n", r.FormValue("num1"))
		a, err := strconv.Atoi(r.FormValue("num1"))
		if err != nil {
			http.Error(w, fmt.Sprintf("could not parse integer: %s", r.FormValue("num1")), 500)
			return
		}
		b, err := strconv.Atoi(r.FormValue("num2"))
		if err != nil {
			http.Error(w, fmt.Sprintf("could not parse integer: %s", r.FormValue("num2")), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"sum":%d}`, a+b)
	})

	port := dflt.EnvString("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
	// 20_OMIT
}
