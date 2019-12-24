package main

import (
	"encoding/json"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func main() {
	s := `
    type Query {
      hello: String!
    }
  `

	schema := graphql.MustParseSchema(s, &query{})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    jsonMap, err := json.Marshal(getGameMap())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMap)
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":4242", nil))
}
