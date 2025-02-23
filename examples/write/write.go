package main

import (
	"log"

	"github.com/roffe/gocanusb"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	ch, err := gocanusb.Open(
		"",
		"500",
		gocanusb.ACCEPTANCE_CODE_ALL,
		gocanusb.ACCEPTANCE_MASK_ALL,
		gocanusb.FLAG_TIMESTAMP|gocanusb.FLAG_BLOCK,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	for i := range 5 {
		log.Println("Writing message: ", i)
		if err := ch.Write(&gocanusb.CANMsg{Id: 0x123, Len: 8, Data: [8]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}); err != nil {
			log.Println(err)
		}
		if err := ch.Flush(gocanusb.FLUSH_WAIT); err != nil {
			log.Println(err)
		}
	}
}
