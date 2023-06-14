package main

import (
	"go-playground/play003_error"
	"go-playground/play005_float"
	"go-playground/play006_tail"
	"log"
)

func main() {
	//log.Printf("hello world")
	//play001_error.Test()
	//play002_error.Test()
	play003_error.Test()

	x := play005_float.AddFloat(0.300, 0.6000)
	log.Println(x)

	y := play005_float.AddInt(300000, 60000)
	log.Println(y)

	play006_tail.Test()
	log.Println(y)
}
