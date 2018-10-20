package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/amnay-mo/fizzbuzz-api/api"
	"github.com/amnay-mo/fizzbuzz-api/utils"
)

func getHTTPPPort() int {
	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 {
		log.Fatalf("Bad APP_PORT: %v", portStr)
	}
	return port
}

func main() {
	http.HandleFunc("/fizzbuzz", api.HandleFizzBuzz)
	appPort := getHTTPPPort()
	appAddr := fmt.Sprintf(":%d", appPort)
	log.Printf("listening on %s", appAddr)
	err := http.ListenAndServe(appAddr, utils.LoggerMiddleware{Next: http.DefaultServeMux})
	if err != nil {
		log.Fatalf("http server exited with error: %v", err)
	}
}
