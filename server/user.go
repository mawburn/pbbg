package main

import (
	//   "crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// 	"strings"
	// 	"github.com/go-redis/redis/v7"
)

type AuthPayload struct {
	Code string `json:"code"`
}

func authUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ap AuthPayload
	err := decoder.Decode(&ap)

	if err != nil {
		panic(err)
	}

	fmt.Println(ap.Code)

	response, postErr := http.PostForm(envVars.OauthUrl, url.Values{
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

	var out interface{}

	json.Unmarshal(body, &out)

	fmt.Println("%+v", out)
}
