package main

import (
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

type EnvVars struct {
	RedisPassword string
	OauthUrl      string
	OauthRedirect string
	ClientId      string
	ClientSecret  string
	CognitoUrl    string
}

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
	envVars.OauthUrl = os.Getenv("OAUTH_URL")
	envVars.OauthRedirect = os.Getenv("OAUTH_REDIRECT")
	envVars.ClientId = os.Getenv("OAUTH_CLIENT_ID")
	envVars.ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")
	envVars.CognitoUrl = os.Getenv("COGNITO_KEY_URL")

	initActions()
	r := initApi()

	generateSectors(dbConns.Redis)

	log.Fatal(http.ListenAndServe(":4242", r))

	defer dbConns.Redis.Close()
}

func ticker() {

}

func initApi() *chi.Mux {
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
	)

	r.Get("/map", getGameMap)
	r.Post("/login", authUser)
	r.Post("/move", playerMove)

	return r
}

func contentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
