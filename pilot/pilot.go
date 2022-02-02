package pilot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Data struct {
	Action              string
	Appkey              string
	PilotFirstname      string
	PilotLastname       string
	PilotCellphone      string
	PilotEmail          string
	PilotContactTypeId  string
	PilotBusinessTypeId string
	PilotSuboriginId    string
	PilotNotes          string
	PilotCarBrand       string
	PilotCarModel       string
}

func (d *Data) SetKey() {
	d.Appkey = "1E1D2D94-F8F3-4E2F-BE02-D15B5DDD9B2A"
	d.PilotSuboriginId = "EVQOWN8FV19NJNMDC"
	d.Action = "create"
}
func (d *Data) Encode() string {
	d.SetKey()
	data := url.Values{}
	data.Set("action", d.Action)
	data.Set("appkey", d.Appkey)
	data.Set("pilot_firstname", d.PilotFirstname)
	data.Set("pilot_lastname", d.PilotLastname)
	data.Set("pilot_cellphone", d.PilotCellphone)
	data.Set("pilot_email", d.PilotEmail)
	data.Set("pilot_contact_type_id", d.PilotContactTypeId)
	data.Set("pilot_business_type_id", d.PilotBusinessTypeId)
	data.Set("pilot_suborigin_id", d.PilotSuboriginId)
	data.Set("pilot_notes", d.PilotNotes)
	data.Set("pilot_car_brand", d.PilotCarBrand)
	data.Set("pilot_car_model", d.PilotCarModel)
	return data.Encode()
}

func Send(data *Data) {

	url := "https://api.pilotsolution.net/webhooks/welcome.php"
	method := "POST"

	str := data.Encode()

	payload := strings.NewReader(str)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
