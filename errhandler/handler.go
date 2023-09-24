package errhandler

import (
	"log"
	"net/http"
)

func LogOnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ThrowOnErr(err error, w *http.ResponseWriter) {
	http.Error(*w, err.Error(), 500)
	log.Println(err)
}
