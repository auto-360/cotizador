// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/mail"
	"os"
	"time"

	"example.com/cloudfunction/autofact"
	mailSender "example.com/cloudfunction/mail"
	"example.com/cloudfunction/pilot"
	"example.com/cloudfunction/utils"
)

var email string = os.Getenv("EMAIL")
var domain string = os.Getenv("DOMAIN")
var apiKey string = os.Getenv("apiKey")

type Client struct {
	Patente     string `json:"patente"`
	Region      int    `json:"region"`
	Kilometraje int    `json:"kilometraje"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Anio        int    `json:"ano"`
	Version     string `json:"version"`
	VersionID   int    `json:"versionID"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	Phone       int    `json:"telefono"`
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetModel(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (r).Method == "OPTIONS" {
		return
	}
	patente := r.URL.Query().Get("patente")
	if len(patente) != 6 {
		return
	}
	value := autofact.GetModels(patente)
	w.Write(*value)
}

//check if string is a valid email
func checkemail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (r).Method == "OPTIONS" {
		return
	}
	tx := autofact.TxRequest{}

	client := Client{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &client)

	tx.CaracteristicasVehiculo.Patente = client.Patente
	tx.CaracteristicasVehiculo.Color = client.Color
	tx.CaracteristicasVehiculo.IDRegion = client.Region
	tx.CaracteristicasVehiculo.Kilometraje = client.Kilometraje

	tx.IDVersion = client.VersionID
	t := time.Now()
	tx.FechaTasacion = t.Format("2006-01-02")
	if checkemail(client.Email) {
		tx.Cliente.Email = client.Email
	}
	tx.Cliente.Nombre = fmt.Sprintf("%s %s", client.Name, client.LastName)
	tx.Cliente.Rut = "" // client.Rut
	tx.Cliente.Solicitante = "Portal"
	if (client.Phone) >= int(math.Pow(10, 7)) {
		tx.Cliente.Telefono = client.Phone
	}
	rx := autofact.CreateTransaction(&tx)
	data := utils.Data{}
	data.Fullname = tx.Cliente.Nombre
	data.Name = client.Name
	data.LastName = client.LastName
	data.Brand = client.Marca
	data.Version = client.Version
	data.TxResponse = &rx
	data.Modelo = client.Modelo
	data.Color = client.Color
	data.Anio = client.Anio
	data.SetHash()

	b, _ := json.Marshal(data)
	w.Write(b)

	mailSender.Send(&data)

}

func CreateAssistance(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (r).Method == "OPTIONS" {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	data := utils.Data{}
	json.Unmarshal(body, &data)

	if !data.Validate() {
		return
	}

	d := pilot.Data{}

	d.PilotFirstname = data.Name
	d.PilotLastname = data.LastName
	d.PilotCellphone = string(data.TxResponse.Cliente.Telefono)
	d.PilotEmail = data.TxResponse.Cliente.Email
	d.PilotContactTypeId = "2"
	d.PilotBusinessTypeId = "2"

	d.PilotCarBrand = data.Brand
	d.PilotCarModel = data.Modelo
	if r.URL.Query().Get("modo") == "Venta" {

		d.PilotNotes = fmt.Sprintf("Vehiculo %s %s %s %s Banda %d - %d interesado en venta directa", data.Brand, data.Modelo,
			data.Version, data.Color,
			data.TxResponse.Indicadores.BandaMin, data.TxResponse.Indicadores.BandaMax)
		pilot.Send(&d)
	} else if r.URL.Query().Get("modo") == "Consignacion" {
		d.PilotNotes = fmt.Sprintf("Vehiculo %s %s %s %s Banda %d - %d interesado en consignacion", data.Brand, data.Modelo,
			data.Version, data.Color,
			data.TxResponse.Indicadores.BandaMin, data.TxResponse.Indicadores.BandaMax)
		pilot.Send(&d)
	}

	u := utils.URL{URL: data.GetMSG()}
	b, _ := json.Marshal(u)
	w.Write(b)
}
