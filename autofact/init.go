package autofact

import (
	"os"
	"time"
)

var URL, TOKEN string

func init() {
	URL = os.Getenv("AUTOFACT_URL")
	USER := os.Getenv("AUTOFACT_USER")
	PASS := os.Getenv("AUTOFACT_PASS")

	go func() {
		for {

			TOKEN = getToken(USER, PASS)
			time.Sleep(time.Minute * 5)
		}

	}()

}
