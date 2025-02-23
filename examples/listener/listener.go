package main

import (
	"log"

	"github.com/roffe/gocanusb"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	// Open the first adapter found
	ch, err := gocanusb.Open(
		"",
		"500",
		gocanusb.ACCEPTANCE_CODE_ALL,
		gocanusb.ACCEPTANCE_MASK_ALL,
		gocanusb.FLAG_TIMESTAMP|gocanusb.FLAG_BLOCK|gocanusb.FLAG_NO_LOCAL_SEND|gocanusb.FLAG_SLOW,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	if err := ch.SetReceiveCallback(callbackHandler); err != nil {
		log.Fatal(err)
	}

	log.Println("Flushing")
	if err := ch.Flush(gocanusb.FLUSH_WAIT); err != nil {
		log.Println(err)
	}
	log.Println("Closing")
	if err := ch.Close(); err != nil {
		log.Fatal(err)
	}
}

func callbackHandler(msg *gocanusb.CANMsg) uintptr {
	// Beware of the callback, it's running in the C world and will replace the content of msg
	// with the next message. If you want to keep the message to use later you need to copy it.
	// For example if you send the message in a channel or similar.
	msgCopy := &gocanusb.CANMsg{
		Id:        msg.Id,
		Timestamp: msg.Timestamp,
		Flags:     msg.Flags,
		Len:       msg.Len,
	}
	copy(msgCopy.Data[:], msg.Data[:])
	log.Println("Callback: ", msgCopy.String())
	return 0
}
