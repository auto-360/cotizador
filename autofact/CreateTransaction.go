package autofact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CaracteristicasVehiculo struct {
	Kilometraje int    `json:"kilometraje"`
	IDRegion    int    `json:"id_region"`
	Color       string `json:"color"`
	Patente     string `json:"patente"`
}
type Cliente struct {
	Solicitante          string `json:"solicitante"`
	Nombre               string `json:"nombre"`
	Rut                  string `json:"rut"`
	Telefono             int    `json:"telefono"`
	Email                string `json:"email"`
	MarcaIntencionCompra string `json:"marca_intencion_compra"`
}

type TxRequest struct {
	FechaTasacion           string                  `json:"fecha_tasacion"`
	IDVersion               int                     `json:"id_version"`
	CaracteristicasVehiculo CaracteristicasVehiculo `json:"caracteristicas_vehiculo"`
	Cliente                 Cliente                 `json:"cliente"`
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

func CreateTransaction(tr *TxRequest) TxResponse {

	url := URL + "/v1/tasaciones/"
	method := "POST"

	payload, _ := json.Marshal(tr)

	fmt.Println(string(payload))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return TxResponse{}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+TOKEN)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return TxResponse{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return TxResponse{}
	}

	var txResponse TxResponse
	err = json.Unmarshal(body, &txResponse)
	if err != nil {
		fmt.Println(err)
	}

	return txResponse
}
