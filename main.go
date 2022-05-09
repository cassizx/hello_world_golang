package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

const DefaultPort = "9090";

func getServerPort() (string) {
	port := os.Getenv("SERVER_PORT");
	if port != "" {
		return port;
	}
	return DefaultPort;
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Ping from %s", r.RemoteAddr)
	fmt.Fprintln(w, "Hello from Golang app")
	fmt.Fprintln(w, "Instance " + os.Getenv("HOSTNAME"))
	fmt.Fprintln(w, "Request URL " + r.URL.Path + " from " + r.RemoteAddr)
}

func main ()  {
	log.Println("Running server on 0.0.0.0:" + getServerPort())
	http.HandleFunc("/", echoHandler)
	http.ListenAndServe(":" + getServerPort(), nil)
}
