package main

import (
	"go-playground/play041_gorm"
	"log"
)

func init() {
	log.Println("init 1")
}

func init() {
	log.Println("init 2")
}

func main() {
	play041_gorm.Play()

	log.Println("waiting select{}")
	select {}
}
