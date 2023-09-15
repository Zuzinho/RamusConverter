package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"ramus/checker"
	"ramus/converter"
)

func main() {
	port := "3000"

	log.Println("Opened server on port 3000")

	mux := http.NewServeMux()

	tmpl := "./templates/index.tmpl"

	mux.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server on '/load'")
		ts, err := template.ParseFiles(tmpl)
		throwOnErr(err, &w)

		err = ts.Execute(w, "Здесь будет текст с описанием модели")
		throwOnErr(err, &w)
	})

	mux.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server on '/convert'")
		ts, err := template.ParseFiles(tmpl)
		throwOnErr(err, &w)

		file, fileHeader, err := r.FormFile("file")
		throwOnErr(err, &w)

		if ok, err := checker.IsCorrectFormat(fileHeader); !ok {
			throwOnErr(err, &w)
		}

		rd, err := converter.Convert(file)
		logOnErr(err)

		buf, err := io.ReadAll(rd)
		logOnErr(err)

		err = ts.Execute(w, string(buf))
		throwOnErr(err, &w)
	})

	err := http.ListenAndServe(":"+port, mux)
	logOnErr(err)
}

func logOnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func throwOnErr(err error, w *http.ResponseWriter) {
	if err != nil {
		http.Error(*w, "Internal Server Error", 500)
		log.Println(err)
	}
}
