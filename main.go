package main

import (
	"go-playground/play301_x509"
	"log"
)

func init() {
	log.Println("init 1")
}

func init() {
	log.Println("init 2")
}

func main() {
	//play602_http_dl.Play()
	//play132_websocktserver.Play()
	play301_x509.Play()

	log.Println("waiting select{}")
	select {}
}
