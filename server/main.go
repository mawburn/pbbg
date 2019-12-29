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
	ClientId string
	ClientSecret string
}

var envVars EnvVars

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}
	
	envVars.RedisPassword = os.Getenv("REDIS_PASS")
	envVars.ClientId = os.Getenv("OAUTH_CLIENT_ID")
	envVars.ClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")

	initActions()
	r, c := initApi()
	defer c.Close()

	generateSectors(c)

	log.Fatal(http.ListenAndServe(":4242", r))
}

func ticker() {

}

func initApi() (*chi.Mux, *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "5!Ne[^5+k2HneNH", // no password set
		DB:       0,  // use default DB
	})

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
    middleware.WithValue("redis", client),
	)

	r.Get("/map", getGameMap)
  r.Post("/login", authUser)

	r.Post("/move", playerMove)

	return r, client
}

func contentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
