package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"

	"example.com/cloudfunction/autofact"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Data struct {
	Fullname   string
	Name       string
	LastName   string
	Brand      string
	Version    string
	Color      string
	Modelo     string
	Anio       int
	TxResponse *autofact.TxResponse
	Hash       string
}

type URL struct {
	URL string
}

func (o *Data) generateHash() string {
	str := fmt.Sprintf("%s%s%s%s%s%d", o.Fullname, o.Brand, o.Version, o.Color, o.Modelo, o.Anio)
	str = fmt.Sprintf("%s%d%d", str, o.TxResponse.Indicadores.BandaMin, o.TxResponse.Indicadores.BandaMax)
	hasher := sha1.New()
	hasher.Write([]byte(str))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha

}
func (o *Data) SetHash() {
	o.Hash = o.generateHash()
}

func (o *Data) Validate() bool {
	return o.Hash == o.generateHash()
}

func (d *Data) GetRegion() string {
	return regiones[d.TxResponse.CaracteristicasVehiculo.IDRegion-1]
}

func (d *Data) GetKilometraje() string {
	p := message.NewPrinter(language.Spanish)
	return p.Sprintf("%d", d.TxResponse.CaracteristicasVehiculo.Kilometraje)
}

func (d *Data) GetPrice() string {
	p := message.NewPrinter(language.Spanish)
	return p.Sprintf("$%d $%d", d.TxResponse.Indicadores.BandaMin, d.TxResponse.Indicadores.BandaMax)
}

func (d *Data) GetMSG() string {

	phoneNumber := "56964185854"
	msg := fmt.Sprintf("Hola quisiera vender mi vehiculo %s %s %s", d.TxResponse.CaracteristicasVehiculo.Patente, d.Brand, d.Modelo)
	msg = url.QueryEscape(msg)
	msg = fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", phoneNumber, msg)
	return msg
}

func number(i int) string {
	nmbr := fmt.Sprintf("%d", i)
	out := "569"
	dif := 11 - len(nmbr)

	if dif > 0 {
		out = out[:dif]
	} else {
		out = ""
	}
	return out + nmbr
}

func (d *Data) GetMSGClient() string {

	phoneNumber := number(d.TxResponse.Cliente.Telefono)
	msg := fmt.Sprintf("Hola estas interesado en vender tu vehiculo %s %s %s", d.TxResponse.CaracteristicasVehiculo.Patente, d.Brand, d.Modelo)
	msg = url.QueryEscape(msg)
	msg = fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", phoneNumber, msg)
	return msg
}

var regiones = []string{
	"Región de Tarapacá",
	"Región de Antofagasta",
	"Región de Atacama",
	"Región de Coquimbo",
	"Región de Valparaíso",
	"Región del Libertador General Bernardo O’Higgins",
	"Región del Maule",
	"Región del Biobío",
	"Región de La Araucanía",
	"Región de Los Lagos",
	"Región Metropolitana de Santiago",
	"Región de Aysén del General Carlos Ibáñez del Campo",
	"Región de Magallanes y la Antártica Chilena",
	"Región de Los Ríos",
	"Región de Arica y Parinacota",
	"Región del Ñuble",
}
