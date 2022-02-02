package mail

import "os"

var apiKey string = os.Getenv("apiKey")
var domain string = os.Getenv("DOMAIN")
