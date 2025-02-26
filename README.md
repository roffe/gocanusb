# goCANUSB

Golang implementation for using the LAWICEL CANUSB Windows DLL

[examples](https://github.com/roffe/gocanusb/tree/master/examples)

Used in [goCAN](https://github.com/roffe/gocan)

## Usage

Download the appropriate DLL and place next to your built Go binary or Windows system folder

- [32-bit](canusbdrv.dll) GOARCH="386"
- [64-bit](canusbdrv64.dll) GOARCH="amd64"

```sh
go get github.com/roffe/gocanusb@latest
```

```go
package main

import (
	"log"

	"github.com/roffe/gocanusb"
)

func main() {
	// open the first available device
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

	ver, err := ch.VersionInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CANUSB version: ", ver)
}
```

## Disclaimer

**LAWICEL CANUSB and LAWICEL CAN232 are trademarks of ELMICRO Computer GmbH & Co. KG**




