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
	TOKEN = getToken(USER, PASS)

	go func() {
		for {
			time.Sleep(time.Minute * 5)
			TOKEN = getToken(USER, PASS)
		}

	}()

}
