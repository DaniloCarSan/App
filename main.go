package main

import (
	"app/app/config"
	"log"
)

func main() {
	log.Println(config.API().Host)
}
