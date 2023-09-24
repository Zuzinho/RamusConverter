package handler

import (
	"html/template"
	"net/http"
	"ramus/checker"
	"ramus/converter"
	"ramus/converter/ramustypes/box"
	"ramus/env"
	"ramus/errhandler"
	"strings"
)

type createBuilderFunc func(mainBox *box.Box) *strings.Builder

func HandleConverter(w http.ResponseWriter, r *http.Request, convert createBuilderFunc) {
	ts, err := template.ParseFiles(env.Templates)
	if err != nil {
		errhandler.ThrowOnErr(err, &w)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		errhandler.ThrowOnErr(err, &w)
		return
	}

	if ok, err := checker.IsCorrectFormat(fileHeader); !ok {
		errhandler.ThrowOnErr(err, &w)
		return
	}

	mainBox, err := converter.Convert(file)
	errhandler.LogOnErr(err)

	builder := convert(mainBox)

	err = ts.Execute(w, builder.String())
	if err != nil {
		errhandler.ThrowOnErr(err, &w)
		return
	}
}

func HandleDefault(w http.ResponseWriter) {
	ts, err := template.ParseFiles(env.Templates)
	if err != nil {
		errhandler.ThrowOnErr(err, &w)
		return
	}

	err = ts.Execute(w, "Здесь будет текст с описанием модели")
	if err != nil {
		errhandler.ThrowOnErr(err, &w)
		return
	}
}
