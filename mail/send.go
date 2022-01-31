package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"time"

	"example.com/cloudfunction/autofact"
	"github.com/mailgun/mailgun-go/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Data struct {
	Fullname   string
	Brand      string
	Version    string
	Color      string
	Modelo     string
	TxResponse *autofact.TxResponse
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

func Send(data *Data) {

	tmpl := template.Must(template.ParseFiles("templates/email.html"))

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Println(err)
		return
	}

	body := tpl.String()

	mg := mailgun.NewMailgun(domain, apiKey)

	sender := "no-responder@auto360.cl"
	subject := "Compra Vehiculo auto360"

	recipient := "malba@mmae.cl"

	message := mg.NewMessage(sender, subject, "Send", recipient)
	message.SetHtml(body)

	//message.AddCC(email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, id, _ := mg.Send(ctx, message)

	// if err != nil {
	// 	//log.Fatal(err)
	// }

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

}
