package autofact

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetModels(patente string) *[]byte {

	url := URL + "/v1/versiones/?patente=" + patente
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+TOKEN)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return &([]byte{})
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return &([]byte{})
	}
	return &body
}
