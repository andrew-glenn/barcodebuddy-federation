package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFilePath = "config/config.json"

var globalConfig Configuration

type Configuration struct {
	RedisUrl      string `json:"RedisUrl"`
	RedisSize     int    `json:"RedisSize"`
	ApiDailyCalls int    `json:"ApiDailyCalls"`
	AdminUser     string `json:"AdminUser"`
	AdminPassword string `json:"AdminPassword"`
	WebserverPort string `json:"WebserverPort"`
}

func loadConfig() {
	if !fileExists(configFilePath) {
		generateDefaultConfig()
	}
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	globalConfig = Configuration{}
	err = decoder.Decode(&globalConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func generateDefaultConfig() {
	config := Configuration{
		RedisUrl:      "127.0.0.1:6379",
		RedisSize:     200,
		ApiDailyCalls: 150,
		AdminUser:     "admin",
		AdminPassword: "admin",
		WebserverPort: ":8080",
	}
	fmt.Println("First start, generated initial configuration")
	saveConfig(config)
}

func saveConfig(config Configuration) {
	file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		os.Exit(1)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&config)
	if err != nil {
		fmt.Println("Error writing configuration:", err)
		os.Exit(1)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
