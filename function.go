// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"example.com/cloudfunction/autofact"
)

var i int

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i = r1.Intn(100)
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func GetModel(w http.ResponseWriter, r *http.Request) {
	patente := r.URL.Query().Get("patente")
	if len(patente) != 6 {
		return
	}
	value := autofact.GetModels(patente)
	w.Write(*value)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	IP := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IP = IP + ipnet.IP.String() + "\n"
			}
		}
	}

	fmt.Fprintf(w, "Hello ,%d , %s!", i, IP)

}
