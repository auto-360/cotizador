// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"

	"example.com/cloudfunction/autofact"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func GetModel(w http.ResponseWriter, r *http.Request) {
	value := autofact.GetModels("bbbb24")
	w.Write(*value)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola TX!")
}
