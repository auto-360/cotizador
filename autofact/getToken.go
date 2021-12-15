package autofact

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type user struct {
	User   string
	Passwd string
}

type token struct {
	Token string `json:"token"`
}

func getToken(user, passwd string) string {

	url := "https://api-integracion.autopress.cl/v1/auth/"
	method := "POST"

	payload := strings.NewReader(`{
  "usuario": "auto360_rest",
  "contrasena": "saBN8cjCtxxFrh567R33zLWm"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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

	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
	}

	return t.Token
}
