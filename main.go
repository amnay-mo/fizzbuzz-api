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

type config struct {
	AppPort   int
	RedisAddr string
	MaxLimit  int
}

func getConfig() *config {
	return &config{
		AppPort:   getHTTPPPort(),
		RedisAddr: getRedisAddr(),
		MaxLimit:  getMaxLimit(),
	}
}

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
func getMaxLimit() int {
	maxLimitStr := os.Getenv("MAX_LIMIT")
	if maxLimitStr == "" {
		maxLimitStr = "0"
	}
	maxLimit, err := strconv.Atoi(maxLimitStr)
	if err != nil || maxLimit < 0 {
		log.Fatalf("Bad MAX_LIMIT: %v", maxLimitStr)
	}
	return maxLimit
}

func getRedisAddr() string {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	return addr
}

func main() {
	configuration := getConfig()
	store := stats.NewRedisStatsStore(configuration.RedisAddr)
	http.Handle("/fizzbuzz", stats.Middleware{
		Store: store,
		Next: fizzbuzz.GetHandlerFunc(
			fizzbuzz.WithMaxLimit(fizzbuzz.Sequence, configuration.MaxLimit)),
	})
	http.Handle("/fizzbuzz/stats", &stats.HTTPHandler{Store: store})
	appAddr := fmt.Sprintf(":%d", configuration.AppPort)
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
