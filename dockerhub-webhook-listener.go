package main

import (
	"flag"
	"log"

	"./server"

	"gopkg.in/gcfg.v1"
	"os"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting server on %s", config.ListenAddr)
	server.NewServer(config)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

func getConfig() (*server.Config, error) {
	var listenAddr = flag.String("listen", "localhost:8080", "<address>:<port> to listen on")
	var configFile = flag.String("config-file", "", "Location of handler config file")
	flag.Parse()

	config := &server.Config{}

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
