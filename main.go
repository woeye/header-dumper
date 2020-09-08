package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	httpPort int
)

const (
	defaultHTTPort = 8000
)

func init() {
	flag.IntVar(&httpPort, "http_port", defaultHTTPort, "Specifies the port to listen on")
}

func main() {
	// Env var overrides flag
	if envPort, found := os.LookupEnv("HTTP_PORT"); found {
		var err error
		httpPort, err = strconv.Atoi(envPort)
		if err != nil {
			log.Fatalf("could not parse env variable HTTP_PORT: %v\n", err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Headers ->")
		for k, v := range r.Header {
			log.Printf("\t%s: %v\n", k, v)
		}

		log.Println("Cookies ->")
		for _, c := range r.Cookies() {
			log.Printf("\t[%s - %s] %s: %v\n", c.Domain, c.Path, c.Name, c.Value)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	log.Printf("Listening on port: %d\n", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil))

}
