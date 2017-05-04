package main

import (
	"flag"
	"log"

	"./server"
	"./server/api"

	"gopkg.in/gcfg.v1"
	"os"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on %s", config.ListenAddr)
	s := server.New(config)
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}

func getConfig() (*api.Config, error) {
	var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")
	var configFile = flag.String("config-file", "", "Location of handler config file")
	flag.Parse()

	config := &api.Config{}

	if *configFile != "" {
		err := gcfg.ReadFileInto(config, *configFile)
		if err != nil {
			return nil, err
		}
	}

	os.Getenv("FOO")
	config.ListenAddr = *listenAddr

	return config, nil
}
