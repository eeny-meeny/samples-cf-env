package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type appConfig struct {
	AppName          string       `json:"application_name"`
	AppUris          []string     `json:"application_uris"`
	Limits           limitsConfig `json:"limits"`
	AppSpaceName     string       `json:"space_name"`
	AppInstanceIndex string
}

type limitsConfig struct {
	Disk int `json:"disk"`
	FDs  int `json:"fds"`
	Mem  int `json:"mem"`
}

func main() {
	indexHTML, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		log.Fatalf("error reading index.html template file: %v", err)
	}

	t, err := template.New("index.html").Parse(string(indexHTML))
	if err != nil {
		log.Fatalf("error parsing index.html template: %v", err)
	}

	srv := &server{t: t}
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "7777"
	}

	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatalf("invalid port number: %v", err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", port), srv)
}

type server struct {
	t *template.Template
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		s.serveTemplate(w, req)
		return
	}
	http.FileServer(http.Dir("./static")).ServeHTTP(w, req)
}

func (s *server) serveTemplate(w http.ResponseWriter, req *http.Request) {
	cfg, err := appConfigFromEnv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.t.ExecuteTemplate(w, "index.html", cfg)
}

func appConfigFromEnv() (*appConfig, error) {
	vcapApp := os.Getenv("VCAP_APPLICATION")
	var cfg appConfig
	if err := json.Unmarshal([]byte(vcapApp), &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON from VCAP_APPLICATION env: %v", err)
	}

	cfg.AppInstanceIndex = os.Getenv("INSTANCE_INDEX")
	return &cfg, nil
}
