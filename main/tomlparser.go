package main

import (
	"os"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type NodeList struct{
	Nodes		[]NodeInfo
}

type NodeInfo struct {
	NodeName	string
	NodeGroup	string
	NodeAddr	string
	NodePort	int
}

// Reads info from config file
func ReadConfig(configfile string) NodeList {

	_, err := os.Stat(configfile)
	check(err)

	dat, err := ioutil.ReadFile(configfile)
	tomlstr := string(dat[:len(dat)])
	check(err)

	println(tomlstr)
	/*
	var configs NodeList
	_, err = toml.Decode(tomlstr, &configs)

	check(err)

	return configs
	*/
	var alice NodeInfo
	var bob NodeInfo
	var cat NodeInfo
	var deb NodeInfo
	var nodes NodeList
	alice.NodeName = "alice"
	alice.NodeAddr = "127.0.0.1"
	alice.NodePort = 13371
	alice.NodeGroup = "group1"
	bob.NodeName = "bob"
	bob.NodeAddr = "127.0.0.1"
	bob.NodePort = 13372
	bob.NodeGroup = "group1"
	cat.NodeName = "alice"
	cat.NodeAddr = "127.0.0.1"
	cat.NodePort = 13373
	cat.NodeGroup = "group1"
	cat.NodeName = "deb"
	cat.NodeAddr = "127.0.0.1"
	cat.NodePort = 13374
	cat.NodeGroup = "group1"
	nodes.Nodes = []NodeInfo{alice, bob, cat, deb}
	return nodes
}


