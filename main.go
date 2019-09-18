package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/amnay-mo/fizzbuzz-api/fizzbuzz"
	"github.com/amnay-mo/fizzbuzz-api/stats"
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

func getRedisAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	return addr
}

func main() {
	store := stats.NewRedisStatsStore(getRedisAddr())
	http.Handle("/fizzbuzz", stats.Middleware{
		Store: store,
		Next:  http.HandlerFunc(fizzbuzz.HandleFizzBuzz),
	})
	http.Handle("/fizzbuzz/stats", &stats.HTTPHandler{Store: store})
	appPort := getHTTPPPort()
	appAddr := fmt.Sprintf(":%d", appPort)
	log.Printf("listening on %s", appAddr)
	err := http.ListenAndServe(
		appAddr,
		utils.LoggerMiddleware{
			Next: http.DefaultServeMux,
		},
	)
	if err != nil {
		log.Fatalf("http server exited with error: %v", err)
	}
}
