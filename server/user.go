package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/segmentio/ksuid"
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
	Email      string  `json:"email"`
	PlayerName *string `json:"playerName"`
	UserName   string  `json:"userName"`
	LastLogin  int32   `json:"lastLogin`
	UserId     string  `json:"userId"`
}

func authUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ap AuthPayload
	err := decoder.Decode(&ap)

	if err != nil {
		fmt.Println("decoder")
		panic(err)
	}

	response, postErr := http.PostForm(envVars.OauthUrl+"/token", url.Values{
		"code":          {ap.Code},
		"grant_type":    {"authorization_code"},
		"client_id":     {envVars.ClientId},
		"client_secret": {envVars.ClientSecret},
		"redirect_uri":  {envVars.OauthRedirect},
	})

	if response.StatusCode == 400 {
		fmt.Println("Code Used")
		return
	}

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

	keySet, jwkErr := jwk.Fetch(envVars.CognitoUrl)

	if jwkErr != nil {
		fmt.Println("jwkErr")
	}

	token, jwtErr := jwt.Parse(out.IdToken, func(token *jwt.Token) (interface{}, error) {
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

	if jwtErr != nil {
		fmt.Println("JWT Error", jwtErr, "\n", out)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		fmt.Println("not ok")
		fmt.Println(ok)
	}

	var user UserInfo

	user = UserInfo{
		Email:     claims["email"].(string),
		UserName:  claims["cognito:username"].(string),
		LastLogin: int32(time.Now().Unix()),
		UserId:    ksuid.New().String(),
	}

	sessionId := ksuid.New().String()

	jsonUser, _ := json.Marshal(user)

	rErr := dbConns.Redis.Set(sessionId, jsonUser, 720*time.Hour).Err()

	if rErr != nil {
		fmt.Println("redis")
		fmt.Println(rErr)
	}

	expire := time.Now().AddDate(0, 0, 30)
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Expires:  expire,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)

	w.Write(json.RawMessage(`{"ok": true}`))
}
