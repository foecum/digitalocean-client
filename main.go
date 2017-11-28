package main

import (
	"digitalocean-client/auth"
	"digitalocean-client/config"
	"flag"
	"log"
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

	client := auth.NewClient(cfg.AccessToken)
	if client == nil {
		log.Fatal("Client not created")
	}
}
