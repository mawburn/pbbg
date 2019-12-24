package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.Handle("/map", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    header := w.Header()
    header.Add("Access-Control-Allow-Origin", "*")
    header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
    header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
  
    if r.Method == "OPTIONS" {
      w.WriteHeader(http.StatusOK)
      return
    }
    
		jsonMap, err := json.Marshal(gameMap())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMap)
	}))

	log.Fatal(http.ListenAndServe(":4242", nil))
}
