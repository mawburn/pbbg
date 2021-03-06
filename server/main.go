package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

// EnvVars - holds enviroment variables to be pulled in
type EnvVars struct {
	RedisPassword string
	OauthURL      string
	OauthRedirect string
	ClientID      string
	ClientSecret  string
	CognitoURL    string
}

// DBConns - holds connections to databases
type DBConns struct {
	Redis *redis.Client
}

var envVars EnvVars
var dbConns DBConns

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVars.RedisPassword = os.Getenv("REDIS_PASS")
	envVars.OauthURL = os.Getenv("OAUTH_URL")
	envVars.OauthRedirect = os.Getenv("OAUTH_REDIRECT")
	envVars.ClientID = os.Getenv("OAUTH_CLIENT_ID")
	envVars.ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	envVars.CognitoURL = os.Getenv("COGNITO_KEY_URL")

	initActions()
	r := initAPI()

	generateSectors()

	go func() {
		runTicker()
	}()

	log.Fatal(http.ListenAndServe(":4242", r))

	defer dbConns.Redis.Close()
}

func ticker() {

}

func initAPI() *chi.Mux {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: envVars.RedisPassword,
		DB:       0,
	})

	dbConns.Redis = client

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*.universeone.win"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(
		cors.Handler,
		render.SetContentType(render.ContentTypeJSON),
		contentMiddleware,
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
		middleware.RealIP,
		authMiddleware,
	)

	// Auth
	r.Post("/login", authUser)

	// Map
	r.Get("/map", getGalaxyMap)

	// Player Actions
	r.Get("/currentSector", playerSector)
	r.Post("/move", playerMove)

	return r
}

func contentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/code" || path == "/login" {
			next.ServeHTTP(w, r)
		} else {
			cookie, cookieErr := r.Cookie("session_id")

			if cookieErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(json.RawMessage(`{"errors": ["401 - Unauthorized"]}`))
				return
			}

			redisVal, redisErr := dbConns.Redis.Get("session:"+cookie.Value).Result()

			if redisErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(json.RawMessage(`{"errors": ["401 - Unauthorized"]}`))
				return
			}

			var jVal UserInfo
			jErr := json.Unmarshal([]byte(redisVal), &jVal)

			if jErr != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(json.RawMessage(`{"errors": ["401 - Unauthorized"]}`))
				return
			}

			ctx := context.WithValue(r.Context(), "userId", jVal.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
