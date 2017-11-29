package main

import (
	"digitalocean-client/config"
	"digitalocean-client/handler"
	"flag"
	"fmt"
	"log"
	"net/http"
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

	http.Handle("/api", handler.New())

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("%v", err)
	}

}
