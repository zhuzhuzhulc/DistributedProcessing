package main

import (
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	nodeName	string
	nodeGroup	string
}

// Reads info from config file
func ReadConfig(configfile string) Config {

	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}


