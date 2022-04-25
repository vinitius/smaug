package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload" // Loading .env file
)

func GetProducts() []string {
	products, found := os.LookupEnv("COINBASE_PRODUCT_IDS")
	if !found {
		log.Fatalf("could not find required config: COINBASE_PRODUCT_IDS")
	}

	return strings.Split(products, "|")
}

func GetChannels() []string {
	channels, found := os.LookupEnv("COINBASE_CHANNELS")
	if !found {
		log.Fatalf("could not find required config: COINBASE_CHANNELS")
	}

	return strings.Split(channels, "|")
}

func GetCoinbaseServiceAddress() string {
	addr, found := os.LookupEnv("COINBASE_SERVICE_ADDRESS")
	if !found {
		log.Fatalf("could not find required config: COINBASE_SERVICE_ADDRESS")
	}

	return addr
}

func GetSlidingWindowSize() int {
	size, found := os.LookupEnv("SLIDING_WINDOW_SIZE")
	if !found {
		log.Fatalf("could not find required config: SLIDING_WINDOW_SIZE")
	}

	asInt, err := strconv.Atoi(size)
	if err != nil {
		log.Fatalf("SLIDING_WINDOW_SIZE must be a valid number: %s", err.Error())
	}

	return asInt
}
