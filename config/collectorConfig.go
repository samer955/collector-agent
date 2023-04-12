package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type CollectorConfig struct {
	topics       []string
	discoveryTag string
}

var config CollectorConfig

func init() {

	if err := godotenv.Load("config.env"); err != nil {
		log.Println("Error loading .env file")
	}

	topics := os.Getenv("TOPICS")
	if topics == "" {
		log.Fatal("TOPICS ENV VARIABLE NOT FOUND")
	}
	config.topics = strings.Split(topics, ",")

	discovery := os.Getenv("DISCOVERY_TAG")
	if discovery == "" {
		log.Fatal("DISCOVERY_TAG ENV VARIABLE NOT FOUND")
	}
	config.discoveryTag = discovery

}

func GetConfig() CollectorConfig {
	return config
}

func (c *CollectorConfig) Topics() []string {
	return c.topics
}

func (c *CollectorConfig) DiscoveryTag() string {
	return c.discoveryTag
}
