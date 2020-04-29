package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Cloud Native Buildpacks!")
	})
	http.ListenAndServe(":8080", nil)
}
