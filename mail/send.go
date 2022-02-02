package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"time"

	_ "embed"

	"example.com/cloudfunction/utils"
	"github.com/mailgun/mailgun-go/v4"
)

//go:embed email.html
var email []byte

//go:embed email_more_info.html
var emailMore []byte

func Send(data *utils.Data) {

	//tmpl := template.Must(template.ParseFiles("email.html"))
	tmpl, _ := template.New("todos").Parse(string(email))

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Println(err)
		return
	}

	body := tpl.String()

	mg := mailgun.NewMailgun(domain, apiKey)

	sender := "no-responder@auto360.cl"
	subject := "Compra Vehiculo auto360"

	recipient := data.TxResponse.Cliente.Email

	message := mg.NewMessage(sender, subject, "Send", recipient)
	message.SetHtml(body)

	//message.AddCC(email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, id, _ := mg.Send(ctx, message)

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	//	tmpl = template.Must(template.ParseFiles("email_more_info.html"))
	tmpl, _ = template.New("todos").Parse(string(emailMore))

	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Println(err)
		return
	}

	sender = "no-responder@auto360.cl"
	subject = fmt.Sprintf("Cotizacion Vehiculo %s %s %s",
		data.TxResponse.CaracteristicasVehiculo.Patente,
		data.Brand,
		data.Modelo)

	recipient = data.TxResponse.Cliente.Email

	message = mg.NewMessage(sender, subject, "Send", recipient)
	message.SetHtml(body)

	//message.AddCC(email)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	resp, id, _ = mg.Send(ctx, message)

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

}
