package main

import (
	"log"
	"net/http"
	"ramus/env"
	"ramus/errhandler"
	"ramus/handler"
	"ramus/stringcreater/list"
	"ramus/stringcreater/table"
)

func main() {
	port := env.Port

	log.Printf("Opened server on port %s", port)

	mux := http.NewServeMux()

	mux.HandleFunc(string(handler.Default), func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server on '%s'", handler.Default)
		handler.HandleDefault(w)
	})

	mux.HandleFunc(string(handler.ConvertList), func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server on '%s'", handler.ConvertList)
		handler.HandleConverter(w, r, list.CreateList)
	})

	mux.HandleFunc(string(handler.ConvertTable), func(w http.ResponseWriter, r *http.Request) {
		log.Printf("server on '%s'", handler.ConvertTable)
		practiceNum := r.FormValue("practice")
		tableFunc, err := table.TablesByPractice(practiceNum)
		if err != nil {
			errhandler.ThrowOnErr(err, &w)
			return
		}

		handler.HandleConverter(w, r, tableFunc)
	})

	err := http.ListenAndServe(":"+port, mux)
	errhandler.LogOnErr(err)
}
