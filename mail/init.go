package mail

import "os"

var apiKey string = os.Getenv("apiKey")
var domain string = os.Getenv("DOMAIN")

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
