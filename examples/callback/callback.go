package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/roffe/gocanusb"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	quitChan := make(chan os.Signal, 2)
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)

	log.Println("Press CTRL-C to exit")
	time.Sleep(1 * time.Second)

	log.Println("Opening first available adapter")
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

	sig := <-quitChan
	log.Println("Signal received:", sig)

	log.Println("Closing")
	if err := ch.Close(); err != nil {
		log.Fatal(err)
	}
}

func callbackHandler(msg *gocanusb.CANMsg) uintptr {
	log.Println("Callback:", msg.String())
	return 0
}
