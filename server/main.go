package main

import (
	"net/http"

	function "example.com/cloudfunction"
)

func main() {
	http.HandleFunc("/GetModel", function.GetModel)
	http.HandleFunc("/CreateTransaction", function.CreateTransaction)

	http.ListenAndServe(":8090", nil)
}
