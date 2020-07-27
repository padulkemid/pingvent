package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/padulkemid/pingpos/auth"
	"github.com/padulkemid/pingpos/config"
	"github.com/padulkemid/pingpos/graph"
	"github.com/padulkemid/pingpos/graph/generated"
	"github.com/rs/cors"
)

const defaultPort = "4000"

func main() {
	config.Connection()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	routerOptions := cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
    // dev : http://localhost:8000
		AllowedOrigins:   []string{"http://localhost:8000", "https://tranquil-atoll-83602.herokuapp.com"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}

	newRouter := cors.New(routerOptions).Handler

	// apply middleware
	router.Use(auth.Middleware())
	router.Use(newRouter)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	websocketTransport := &transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// add domains here
				// field for tommorrow
				return r.Host == "herokuapp.com"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	// Add the transport request
	srv.AddTransport(websocketTransport)

  // disable this line in prod
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}

	log.Printf("server is started on http://localhost:%s", port)
}
