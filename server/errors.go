package main

import (
	"encoding/json"
	"net/http"
)

// Err400 - writes 400 error
func Err400(w http.ResponseWriter, errors []string) {
	writeErr(w, errors, http.StatusBadRequest)
}

// Err404 - writes 404 error
func Err404(w http.ResponseWriter) {
	writeErr(w, []string{"Not Found"}, http.StatusNotFound)
}

// Err500 - - writes 500 error
func Err500(w http.ResponseWriter, errors []string) {
	writeErr(w, errors, http.StatusInternalServerError)
}

func writeErr(w http.ResponseWriter, errors []string, code int) {
	errArr, _ := json.Marshal(errors)
	w.WriteHeader(code)
	w.Write(json.RawMessage(`{"errors": ` + string(errArr) + `}`))
}
