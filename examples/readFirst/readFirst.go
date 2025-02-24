package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	canusb "github.com/roffe/gocanusb"
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
	ch, err := canusb.Open(
		"",
		"500",
		canusb.ACCEPTANCE_CODE_ALL,
		canusb.ACCEPTANCE_MASK_ALL,
		canusb.FLAG_TIMESTAMP|canusb.FLAG_BLOCK|canusb.FLAG_BLOCK|canusb.FLAG_NO_LOCAL_SEND|canusb.FLAG_SLOW,
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

	if err := ch.SetTimeouts(100, 100); err != nil {
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
		msg, err := ch.ReadFirst(0x123, canusb.CANMSG_STANDARD)
		if err != nil {
			if err == canusb.ErrNoMessage || err == canusb.ErrTimeout {
				continue
			}

			log.Println(err)
			return
		}
		log.Println(msg.String())
	}
}
