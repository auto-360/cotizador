package autofact

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TxRequest struct {
	FechaTasacion           string `json:"fecha_tasacion"`
	IDVersion               int    `json:"id_version"`
	CaracteristicasVehiculo struct {
		Kilometraje int    `json:"kilometraje"`
		IDRegion    int    `json:"id_region"`
		Color       string `json:"color"`
		Patente     string `json:"patente"`
	} `json:"caracteristicas_vehiculo"`
	AjustesAdicionales []struct {
		Descripcion string `json:"descripcion"`
		Monto       int    `json:"monto"`
	} `json:"ajustes_adicionales"`
	Cliente struct {
		Solicitante          string `json:"solicitante"`
		Nombre               string `json:"nombre"`
		Rut                  string `json:"rut"`
		Telefono             int    `json:"telefono"`
		Email                string `json:"email"`
		MarcaIntencionCompra string `json:"marca_intencion_compra"`
	} `json:"cliente"`
	Imagenes []interface{} `json:"imagenes"`
}

type TxResponse struct {
	ID                      int    `json:"id"`
	IDVersion               int    `json:"id_version"`
	FechaConsulta           string `json:"fecha_consulta"`
	FechaResultados         string `json:"fecha_resultados"`
	CaracteristicasVehiculo struct {
		Kilometraje int    `json:"kilometraje"`
		IDRegion    int    `json:"id_region"`
		Color       string `json:"color"`
		Patente     string `json:"patente"`
	} `json:"caracteristicas_vehiculo"`
	Imagenes           []interface{} `json:"imagenes"`
	PrecioAutopress    int           `json:"precio_autopress"`
	AjustesAdicionales []interface{} `json:"ajustes_adicionales"`
	Indicadores        struct {
		BandaMin              int     `json:"banda_min"`
		BandaMax              int     `json:"banda_max"`
		KmReferencia          int     `json:"km_referencia"`
		ValorResidual         float64 `json:"valor_residual"`
		Presencia             string  `json:"presencia"`
		CantidadPublicaciones int     `json:"cantidad_publicaciones"`
	} `json:"indicadores"`
	PrecioRetomaInicial int `json:"precio_retoma_inicial"`
	Cliente             struct {
		Solicitante          string `json:"solicitante"`
		Nombre               string `json:"nombre"`
		Rut                  string `json:"rut"`
		Telefono             int    `json:"telefono"`
		Email                string `json:"email"`
		MarcaIntencionCompra string `json:"marca_intencion_compra"`
	} `json:"cliente"`
}

func CreateTransaction() {

	url := URL + "/v1/tasaciones/"
	method := "POST"

	payload := strings.NewReader(`{
  "fecha_tasacion": "2021-12-14",
  "id_version": 61417,
  "caracteristicas_vehiculo": {
    "kilometraje": 0,
    "id_region": 13,
    "color": "azul",
    "patente": "TD4301"
  },
  "ajustes_adicionales": [],
  "cliente": {
    "solicitante": "string",
    "nombre": "Jorge Cáceres",
    "rut": "23858481-7",
    "telefono": 999999999,
    "email": "cliente@gmail.com",
    "marca_intencion_compra": "si"
  },
  "imagenes": []
}
`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+TOKEN)

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
