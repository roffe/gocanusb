package main

import (
	"log"
	"time"

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
	log.Printf("# adapters: %d, %s", len(adapters), adapters)

	ch, err := gocanusb.Open(
		adapters[0],
		"500",
		gocanusb.ACCEPTANCE_CODE_ALL,
		gocanusb.ACCEPTANCE_MASK_ALL,
		gocanusb.FLAG_TIMESTAMP|gocanusb.FLAG_BLOCK|gocanusb.FLAG_NO_LOCAL_SEND|gocanusb.FLAG_SLOW,
	)
	if err != nil {
		log.Fatal(err)
	}

	cb := func(msg *gocanusb.CANMsg) uintptr {
		// Beware of the callback, it's running in the C world and will replace the content of msg
		// with the next message. If you want to keep the message to use later you need to copy it.
		// For example if you send the message in a channel or similar.
		var dataCopy [8]byte
		copy(dataCopy[:], msg.Data[:])

		msgCopy := &gocanusb.CANMsg{
			Id:        msg.Id,
			Timestamp: msg.Timestamp,
			Flags:     msg.Flags,
			Len:       msg.Len,
			Data:      dataCopy,
		}

		log.Println("Callback: ", msgCopy.String())
		return 0
	}
	if err := ch.SetReceiveCallback(cb); err != nil {
		log.Fatal(err)
	}

	for i := range 5 {
		log.Println("Writing message: ", i)
		if err := ch.Write(&gocanusb.CANMsg{Id: 0x123, Len: 8, Data: [8]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}}); err != nil {
			log.Println(err)
		}
	}

	time.Sleep(10 * time.Second)
	log.Println("Flushing")
	if err := ch.Flush(gocanusb.FLUSH_WAIT); err != nil {
		log.Println(err)
	}
	log.Println("Closing")
	if err := ch.Close(); err != nil {
		log.Fatal(err)
	}

}
