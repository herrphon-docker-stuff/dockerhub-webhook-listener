package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"./api"
	"./handler"
)

type Server struct {
	httpClient *http.Client
	config     *api.Config
	handlers   HandlerRegistry
}

func New (config *api.Config) *Server {
	return &Server{&http.Client{}, config, HandlerRegistry{}}
}

func (s *Server) RegisterHandler() {
	s.handlers.Add((&handler.Logger{}).Call)

	// handlers.Add((&listener.Mailgun{config.Mailgun}).Call)
}

func (s *Server) Serve() error {
	if len(s.config.Apikeys.Key) == 0 {
		log.Print("Warning: The server is about to start without any authentication.  Anyone can trigger handlers to fire off")
		log.Print("To enable authentication, you must add an `apikeys` section with at least 1 `key`")
	}

	http.HandleFunc("/", s.reqHandler)
	if s.config.Tls.Key != "" && s.config.Tls.Cert != "" {
		log.Print("Starting with SSL")
		return http.ListenAndServeTLS(s.config.ListenAddr, s.config.Tls.Cert, s.config.Tls.Key, Log(http.DefaultServeMux))
	}
	log.Print("Warning: Server is starting without SSL, you should not pass any credentials using this configuration")
	log.Print("To use SSL, you must provide a config file with a [tls] section, and provide locations to a `key` file and a `cert` file")
	return http.ListenAndServe(s.config.ListenAddr, Log(http.DefaultServeMux))
}

// Send callback request
func (s *Server) sendCallback(callbackUrl string, msg *api.CallbackMessage) {
	log.Printf("Send callback to %s", callbackUrl)

	jsonStr, err := json.Marshal(msg)
	if err != nil {
		log.Print("Failed to marshal callback message")
		log.Print(err)
		return
	}

	req, err := http.NewRequest("POST", callbackUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Print("Failed to make callback request")
		log.Print(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = s.httpClient.Do(req)

	if err != nil {
		log.Print("Failed to request callback")
		log.Print(err)
		return
	}

	log.Print("Succeeded to request callback")
}

func (s *Server) reqHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var imgConfig api.HubMessage

	err := decoder.Decode(&imgConfig)
	if err != nil {
		http.Error(w, "Could not decode json", 500)
		log.Print(err)
		return
	}

	if s.authenticateRequest(r) {
		go s.handleMsg(imgConfig)
		return
	}

	http.Error(w, "Not Authorized", 401)
	s.sendCallback(imgConfig.Callback_url, &api.CallbackMessage{
		State:       "failure",
		Description: "Not authorized",
	})
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.RemoteAddr, r.Method)
		handler.ServeHTTP(w, r)
	})
}

func (s *Server) authenticateRequest(r *http.Request) bool {
	key := r.URL.Query().Get("apikey")
	for _, cfgKey := range s.config.Apikeys.Key {
		if key == cfgKey {
			return true
		}
		continue
	}
	return (len(s.config.Apikeys.Key) == 0) || false
}

func (s *Server) handleMsg(msg api.HubMessage) {
	s.handlers.Call(msg)

	/*
		//TODO: fix callback later
		sendCallback(img.Callback_url, &CallbackMessage{
			State: "success",
			Description: "Hook successfully received",
		})
	*/
}
