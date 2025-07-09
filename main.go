package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gautamb02/sso-service/confreader"
)

var configPath string

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
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config loaded: %+v\n", config)
}
