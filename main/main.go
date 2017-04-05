package main

import (
	"os"
)

func main(){
	if len(os.Args) < 2{
		println("Need to pass in config file location!")
		return
	}
	configfile := os.Args[1]
	var configs = ReadConfig(configfile)
	println(configs.nodeName)
}