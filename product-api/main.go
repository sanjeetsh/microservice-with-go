package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	//create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	//create a new server
	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigtern ir interupt or gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
