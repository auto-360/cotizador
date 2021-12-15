package autofact

import "os"

var URL, TOKEN string

func init() {
	URL = os.Getenv("AUTOFACT_URL")
	USER := os.Getenv("AUTOFACT_USER")
	PASS := os.Getenv("AUTOFACT_PASS")

	TOKEN = getToken(USER, PASS)
}
