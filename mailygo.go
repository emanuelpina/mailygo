package main

import (
	"log"
	"net/http"
	"strconv"
)

var appConfig *config

func init() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}
	appConfig = cfg
}

func main() {
	if !checkRequiredConfig(appConfig) {
		log.Fatal("Not all required configurations are set")
	}
	http.HandleFunc("/", FormHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(appConfig.Port), nil))
}
