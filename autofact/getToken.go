package autofact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	User   string `json:"usuario"`
	Passwd string `json:"contrasena"`
}

type token struct {
	Token string `json:"token"`
}

func getToken(username, passwd string) string {

	url := URL + "/v1/auth/"
	method := "POST"

	u := user{username, passwd}
	payload, _ := json.Marshal(u)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	t := token{}

	fmt.Println(string(body))
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
	}

	return t.Token
}
