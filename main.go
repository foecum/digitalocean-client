package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/foecum/digitalocean-client/config"
	"github.com/foecum/digitalocean-client/handler"
)

func main() {
	//Make sure the file exists
	configType := flag.String("configtype", "file", "Type of configuration e.g environ, file")
	configPath := flag.String("configpath", "", "path of the config file")

	flag.Parse()

	cfg := config.Config{}

	switch *configType {
	case "file":
		err := cfg.ReadConfig(*configPath)
		if err != nil {
			log.Fatal(err)
		}
	case "environ":
		err := cfg.UseCustomEnvConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	h := handler.Handler{}
	h.New(cfg.AccessToken)
	// Create a server
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: h.Mux,
	}

	// Check for a closing signal
	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	// Start server
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}

}
