package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func cognitoCallback(ctx context.Context, d *aegis.HandlerDependencies, req *aegis.APIGatewayProxyRequest, res *aegis.APIGatewayProxyResponse, params url.Values) error {
	// 	// Exchange code for token
	// 	tokens, err := d.Services.Cognito.GetTokens(req.QueryStringParameters["code"], []string{})
	// 	if err != nil {
	// 		log.Println("Couldn't get access token", err)
	// 		res.JSONError(500, err)
	// 	} else {
	// 		// verify the token
	// 		_, err := d.Services.Cognito.ParseAndVerifyJWT(tokens.IDToken)
	// 		if err == nil {
	// 			host := req.GetHeader("Host")
	// 			stage := req.RequestContext.Stage
	// 			res.SetHeader("Set-Cookie", "access_token="+tokens.AccessToken+"; Domain="+host+"; Secure; HttpOnly")
	// 			res.Redirect(301, "https://"+host+"/"+stage+"/protected")
	// 		} else {
	// 			res.JSONError(401, errors.New("unauthorized, invalid token"))
	// 		}
	// 	}
	// 	return nil
}
