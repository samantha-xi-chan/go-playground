package main

import (
	"go-playground/play033_rmq_ttl"
	"log"
)

func init() {
	log.Println("init...")
}

func main() {
	play033_rmq_ttl.Play()
	//play602_http_dl.Play()

	select {}
}
