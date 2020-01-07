package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/segmentio/ksuid"
)

type authPayload struct {
	Code string `json:"code"`
}

type oAuth struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// UserInfo - information about the user
type UserInfo struct {
	Email      string  `json:"email"`
	PlayerName *string `json:"playerName"`
	LastLogin  int32   `json:"lastLogin`
	UserID     string  `json:"userId"`
}

var keySet *jwk.Set

const startingSector string = "0011"
const startingSystem string = "S001"
const userIDSalt string = "Matt's PBBG!! and Aspen helped :)"

func authUser(w http.ResponseWriter, r *http.Request) {
	oAuth, err := getAuthToken(r.Body)

	if err != nil {
		Err500(w, []string{err.Error()})
		return
	}

	token, err := parseJwtToken(oAuth.IDToken)

	if err != nil {
		Err500(w, []string{`JWT Error - ` + err.Error()})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		Err500(w, []string{"Token claims invalid"})
		return
	}

	cognitoUserID := claims["cognito:username"].(string)

	h := sha1.New()
	h.Write([]byte(cognitoUserID + userIDSalt))

	bs := fmt.Sprintf("%x", h.Sum(nil))

	userIDRune := []rune(bs)
	userID := string(userIDRune[0:20])

	userVal, err := dbConns.Redis.Get("user:"+userID).Result()

	var user UserInfo

	if err != nil {
		user = UserInfo{
			Email:     claims["email"].(string),
			LastLogin: int32(time.Now().Unix()),
			UserID:    userID,
		}

		// need to trigger player name ask
	} else {
		err := json.Unmarshal([]byte(userVal), &user)

		user.LastLogin = int32(time.Now().Unix())

		if err != nil {
			Err500(w, []string{"Error unmarshalling user store"})
			return
		}
	}

	sessionID := ksuid.New().String()

	jsonUser, err := json.Marshal(user)

	if err != nil {
		Err500(w, []string{"Error unmarhsalling user"})
		return
	}

	rErr := dbConns.Redis.Set("session:"+sessionID, jsonUser, 720*time.Hour).Err()
	rErr2 := dbConns.Redis.Set("user:"+userID, jsonUser, 0).Err()

	if rErr != nil {
		Err500(w, []string{"Error setting session"})
		return
	} else if rErr2 != nil {
		Err500(w, []string{"Error updating/adding user"})
		return
	}

	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  time.Now().AddDate(0, 0, 30),
		HttpOnly: true,
		Secure:   true,
	}

	playerExists := dbConns.Redis.Exists("player:" + userID).Val()

	if playerExists == 0 {
		pData, _ := json.Marshal(Player{CurSectorId: startingSector, CurSystemId: startingSystem})
		err := dbConns.Redis.Set("player:"+userID, pData, 0).Err()

		if err != nil {
			fmt.Println(err)
			Err500(w, []string{"Error adding player"})
			return
		}
	}

	http.SetCookie(w, &cookie)
	w.Write(json.RawMessage(`{"id": "` + userID + `", "email": "` + user.Email + `"}`))
}

func getAuthToken(bodyIo io.ReadCloser) (oAuth, error) {
	decoder := json.NewDecoder(bodyIo)
	var ap authPayload
	err := decoder.Decode(&ap)

	if err != nil {
		return oAuth{}, fmt.Errorf("Error decoding body")
	}

	response, err := http.PostForm(envVars.OauthURL+"/token", url.Values{
		"code":          {ap.Code},
		"grant_type":    {"authorization_code"},
		"client_id":     {envVars.ClientID},
		"client_secret": {envVars.ClientSecret},
		"redirect_uri":  {envVars.OauthRedirect},
	})

	if err != nil {
		return oAuth{}, err
	} else if response.StatusCode == 400 {
		return oAuth{}, fmt.Errorf("Auth Code Used")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return oAuth{}, fmt.Errorf("Error reading OAuth body")
	}

	var out oAuth

	err = json.Unmarshal(body, &out)

	return out, err
}

func parseJwtToken(tokenStr string) (*jwt.Token, error) {
	if keySet == nil {
		tmpKeySet, err := jwk.Fetch(envVars.CognitoURL)

		if err != nil {
			return nil, fmt.Errorf("Unable to retrieve cognito key")
		}

		keySet = tmpKeySet
	}

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("kid header not found")
		}
		keys := keySet.LookupKeyID(kid)
		if len(keys) == 0 {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		return keys[0].Materialize()
	})
}
