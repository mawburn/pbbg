package main

import (
	//   "crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"net/http"
	"net/url"
	// 	"strings"
	// 	"github.com/go-redis/redis/v7"
)

type AuthPayload struct {
	Code string `json:"code"`
}

type OAuth struct {
	IdToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type UserInfo struct {
	*OAuth
	Email      string   `json:"email"`
	PlayerName string   `json:"playerName"`
	UserNames  []string `json:"userNames"`
	LastLogin  int64    `json:"lastLogin`
	SessionId  string   `json:"sessionId"`
}

func authUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ap AuthPayload
	err := decoder.Decode(&ap)

	if err != nil {
		panic(err)
	}

	fmt.Println(ap.Code)

	response, postErr := http.PostForm(envVars.OauthUrl+"/token", url.Values{
		"code":          {ap.Code},
		"grant_type":    {"authorization_code"},
		"client_id":     {envVars.ClientId},
		"client_secret": {envVars.ClientSecret},
		"redirect_uri":  {envVars.OauthRedirect},
	})

	if postErr != nil {
		fmt.Println("post error", envVars.OauthUrl)
	}

	defer response.Body.Close()
	body, ioErr := ioutil.ReadAll(response.Body)

	if ioErr != nil {
		fmt.Println("io Error")
	}

	var out OAuth

	json.Unmarshal(body, &out)

	keySet, err := jwk.Fetch(envVars.CognitoUrl)

	token, err := jwt.Parse(out.IdToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			fmt.Println("kid header not found")
			return nil, nil
		}
		keys := keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		return keys[0].Materialize()
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"], claims["cognito:username"], "\n")
	} else {
		fmt.Println(err)
	}
}
