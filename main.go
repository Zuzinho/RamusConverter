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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server on '/'")
		ts, err := template.ParseFiles(tmpl)
		if err != nil {
			throwOnErr(err, &w)
			return
		}

		err = ts.Execute(w, "Здесь будет текст с описанием модели")
		if err != nil {
			throwOnErr(err, &w)
			return
		}
	})

	mux.HandleFunc("/convert/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server on '/convert/'")
		ts, err := template.ParseFiles(tmpl)
		if err != nil {
			throwOnErr(err, &w)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			throwOnErr(err, &w)
			return
		}

		if ok, err := checker.IsCorrectFormat(fileHeader); !ok {
			throwOnErr(err, &w)
			return
		}

		rd, err := converter.Convert(file)
		logOnErr(err)

		buf, err := io.ReadAll(rd)
		logOnErr(err)

		err = ts.Execute(w, string(buf))
		if err != nil {
			throwOnErr(err, &w)
			return
		}
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
	http.Error(*w, err.Error(), 500)
	log.Println(err)
}
