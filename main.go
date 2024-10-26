package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/manikant8/GoServer-rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// take the envt var from env file & pull them in current envt
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	// initializing a new instance of a router using the chi package
	router := chi.NewRouter()

	// To request server from browser
	// CORS (Cross-Origin Resource Sharing)
	// CORS (Cross-Origin Resource Sharing) is a security feature implemented by web browsers that allows
	// or restricts web applications running at one origin (domain) to make requests to a different origin (domain).
	// This is typically referred to as cross-origin requests.
	// the router will handle routing requests to the appropriate endpoints.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1Router.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Get("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	// srv: This is a variable that stores a pointer to the newly created HTTP server.
	// &http.Server{...}: This creates a new instance of the http.Server struct. The & symbol means it’s a pointer to this new server.
	srv := &http.Server{
		// Handler: router: The Handler field is specifying which handler the server should use to process incoming HTTP requests.
		Handler: router,

		// The Addr field specifies the address on which the server should listen for incoming requests.
		Addr: ":" + portString,
	}

	log.Printf("Server Strarting on port %v", portString)

	// ListenAndServe() is a method of the http.Server struct that starts the server and listens for incoming HTTP requests on the address specified by srv.Addr
	// If the server starts successfully, it will keep running and listen for HTTP requests. The server will only return an error (which will be stored in err) if something goes wrong, like:
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
