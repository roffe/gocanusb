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
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)

	log.Println("Press CTRL-C to exit")
	time.Sleep(1 * time.Second)

	log.Println("Opening first available adapter")
	ch, err := gocanusb.Open(
		"",
		"500",
		gocanusb.ACCEPTANCE_CODE_ALL,
		gocanusb.ACCEPTANCE_MASK_ALL,
		gocanusb.FLAG_TIMESTAMP|gocanusb.FLAG_BLOCK|gocanusb.FLAG_BLOCK|gocanusb.FLAG_NO_LOCAL_SEND|gocanusb.FLAG_SLOW,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	ver, err := ch.VersionInfo()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Version info:", ver)

	if err := ch.SetTimeouts(100, 50); err != nil {
		log.Println(err)
		return
	}

	for {
		select {
		case sig := <-quitChan:
			log.Println("Signal received:", sig)
			return
		default:
		}
		msg, err := ch.Read()
		if err != nil {
			if err == gocanusb.ErrNoMessage || err == gocanusb.ErrTimeout {
				continue
			}

			log.Println(err)
			return
		}
		log.Println(msg.String())
	}
}
