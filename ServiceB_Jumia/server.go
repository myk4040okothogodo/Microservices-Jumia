package main

import (
  "context"
  "fmt"
	"log"
	"net/http"
	"os"
  "os/signal"
  "time"
   database "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/pkg/db/migrations/mysql"
   "github.com/go-chi/chi/v5"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/graph"
  "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/graph/generated"

)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  serviceASvc := os.Getenv("SERVICE_A_SERVICE")
  if serviceASvc == "" {
      log.Printf("Failed to load environment variable: %s", "SERVICE_A_SERVICE", setting default),
      serviceASvc = "localhost:9092"
  }

  


  //svc, err := apiHolder.NewServiceHolder()
  //if err != nil {
   //     log.Fatalf("failed to create grpc api holder: %s", err)
  //}

  conn, err := grpc.Dial("localhost:9092")
  if err != nil {
      panic(err)
  }

  defer conn.Close()

  router := chi.NewRouter() 
  database.InitDB()
  database.Migrate()

	server := handler.NewDefaultServer(ServiceB_Jumia.NewExecutableSchema(generated.Config{Resolvers:  &graph.Resolver{}}))
  //server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.NewResolver(svc)}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)


  srv :=  &http.Server{
        Addr:    fmt.Sprintf(":%s", port),
        WriteTimeout : time.Second * 10,
        ReadTimeout : time.Second * 10,
        Handler : router,
  }


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
 
  go func() {

  log.Fatal(srv.ListenAndServe())
  }()

  sig := make(chan os.Signal, 1)
  signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

  //Block untill cancel signal is received
  <-sig
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  defer cancel()
  log.Print("shutting down graph_api server")

  if err := srv.Shutdown(ctx); err != nil {
      log.Print(err)
  }

  <-ctx.Done()
  os.Exit(0)
