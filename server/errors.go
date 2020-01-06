package main

import (
	"encoding/json"
	"net/http"
)

func Err400(w http.ResponseWriter, errors []string) {
	writeErr(w, errors, http.StatusBadRequest)
}

func Err404(w http.ResponseWriter) {
	writeErr(w, []string{"Not Found"}, http.StatusNotFound)
}

func Err500(w http.ResponseWriter, errors []string) {
	writeErr(w, errors, http.StatusInternalServerError)
}

func writeErr(w http.ResponseWriter, errors []string, code int) {
	errArr, _ := json.Marshal(errors)
	w.WriteHeader(code)
	w.Write(json.RawMessage(`{"errors": ` + string(errArr) + `}`))
}
