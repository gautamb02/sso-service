package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gautamb02/sso-service/confreader"
	"github.com/gautamb02/sso-service/logger"
	"github.com/gautamb02/sso-service/server"
)

var configPath string
var Log *log.Logger

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
}

func main() {
	flag.Parse()

	if configPath == "" {
		fmt.Println("Error: --config is required")
		flag.Usage()
		os.Exit(1)
	}

	configReader := confreader.NewConfigReader(configPath)
	config, err := configReader.GetConfig()
	fmt.Printf("Logger : %s\n", config.Logger)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}
	err = logger.InitLogger(config.Logger)
	if err != nil {
		fmt.Printf("❌ Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	server := server.NewServer(config)
	server.Start()

}
