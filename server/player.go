package main

import (
	"encoding/json"
	"net/http"
  "fmt"
  "strings"
  "crypto/rand"
  
  "github.com/go-redis/redis/v7"
)

type PlayerMove struct {
	Direction string `json:"direction"`
}

type Player struct {
	Id string
}

func playerMove(w http.ResponseWriter, r *http.Request, c *redis.Client) {
  reqToken := r.Header.Get("Authorization")
  splitToken := strings.Split(reqToken, "Bearer ")
  reqToken = splitToken[1]
  
  fmt.Println(reqToken)
  
	decoder := json.NewDecoder(r.Body)
	var m PlayerMove
	err := decoder.Decode(&m)

	if err != nil {
		panic(err)
	}

	w.Write(json.RawMessage(`{"precomputed": true}`))
}

func getToken(r *http.Request, c *redis.Client) string {
  reqToken := r.Header.Get("Authorization")
  splitToken := strings.Split(reqToken, "Bearer ")
  reqToken = splitToken[1]
  
  if false {
    return "X"
  }
  
  b := make([]byte, 16)
	rand.Read(b)
	fmt.Sprintf("%x", b)
  
  return string(b)
}
