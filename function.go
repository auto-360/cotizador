// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"example.com/cloudfunction/autofact"
)

type Client struct {
	IdVersion            int    `json:"idVersion"`
	Patente              string `json:"patente"`
	Color                string `json:"color"`
	IDRegion             int    `json:"idRegion"`
	Kilometraje          int    `json:"kilometraje"`
	Nombre               string `json:"nombre"`
	Apellido             string `json:"apellido"`
	Rut                  string `json:"rut"`
	Telefono             int    `json:"telefono"`
	Email                string `json:"email"`
	MarcaIntencionCompra string `json:"marcaIntencionCompra"`
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func GetModel(w http.ResponseWriter, r *http.Request) {
	patente := r.URL.Query().Get("patente")
	if len(patente) != 6 {
		return
	}
	value := autofact.GetModels(patente)
	w.Write(*value)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

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
	tx.CaracteristicasVehiculo.IDRegion = client.IDRegion
	tx.CaracteristicasVehiculo.Kilometraje = client.Kilometraje

	tx.IDVersion = client.IdVersion
	t := time.Now()
	tx.FechaTasacion = t.Format("2006-01-02")
	tx.Cliente.Email = client.Email
	tx.Cliente.Nombre = fmt.Sprintf("%s %s", client.Nombre, client.Apellido)
	tx.Cliente.Rut = client.Rut
	tx.Cliente.Solicitante = "string"
	tx.Cliente.Telefono = client.Telefono
	tx.Cliente.MarcaIntencionCompra = client.MarcaIntencionCompra

	rx := autofact.CreateTransaction(&tx)
	b, _ := json.Marshal(rx)
	w.Write(b)

}

func CreateAssistance(w http.ResponseWriter, r *http.Request) {

}
