# goCANUSB

Golang implementation for using the ELMICRO CANUSB Windows DLL

## Installation

You need the DLL for your CANUSB which can be found at [https://www.canusb.com/support/canusb-support/](https://www.canusb.com/support/canusb-support/)

Download the appropriate setup based on your target GOARCH 32-bit = 386, 64-bit = amd64

**Make sure .NET support is installed**: Windows 10 users will need to enable .NET framework 3.5 (which includes v2.0) before using CANUSB DLL.

To install .NET 3.5 support go to “Control Panel” then “Programs and Features” and then “Turn Windows features on or off” (on left menu), then you can enable Microsoft .NET Framework v3.5 which also adds support for 2.0 which is required by the CANUSB DLL. Then reboot PC and proceed with the following step.

## Disclaimer

**LAWICEL CANUSB and LAWICEL CAN232 are trademarks of ELMICRO Computer GmbH & Co. KG**

## Showcase

Used in [goCAN](https://github.com/roffe/gocan)

example:

```go
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
```