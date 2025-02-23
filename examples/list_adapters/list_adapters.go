package main

import (
	"log"

	"github.com/roffe/gocanusb"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	adapters, err := gocanusb.GetAdapters()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d adapters", len(adapters))
	for i, adapter := range adapters {
		log.Printf("Adapter #%d: szID: %s", i, adapter)
	}
}
