package config

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetProducts() []string {
	products, found := os.LookupEnv("COINBASE_PRODUCTS")
	if !found {
		log.Panicf("could not find required config: COINBASE_PRODUCTS")
	}

	return strings.Split(products, "|")
}

func GetChannels() []string {
	channels, found := os.LookupEnv("COINBASE_CHANNELS")
	if !found {
		log.Panicf("could not find required config: COINBASE_CHANNELS")
	}

	return strings.Split(channels, "|")
}

func GetCoinbaseServiceAddress() string {
	addr, found := os.LookupEnv("COINBASE_SERVICE_ADDRESS")
	if !found {
		log.Panicf("could not find required config: COINBASE_SERVICE_ADDRESS")
	}

	return addr
}

func GetSlidingWindowSize() int {
	size, found := os.LookupEnv("SLIDING_WINDOW_SIZE")
	if !found {
		log.Panicf("could not find required config: SLIDING_WINDOW_SIZE")
	}

	asInt, err := strconv.Atoi(size)
	if err != nil {
		log.Panicf("SLIDING_WINDOW_SIZE must be a valid number: %s", err.Error())
	}

	return asInt
}
