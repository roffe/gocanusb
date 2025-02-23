# goCANUSB

Golang implementation for using the LAWICEL CANUSB Windows DLL

[examples](https://github.com/roffe/gocanusb/tree/master/examples)

Used in [goCAN](https://github.com/roffe/gocan)

## Installation

You need the DLL for your CANUSB which can be found at [https://www.canusb.com/support/canusb-support/](https://www.canusb.com/support/canusb-support/)

The bare minimum needed is the [32-bit](canusbdrv.dll) or [64-bit](canusbdrv64.dll) DLL placed in the same directory as your Go binary.

Download the appropriate based on your target GOARCH 386=32-bit, amd64=64-bit

**Make sure .NET support is installed**: Windows users will need to enable .NET framework 3.5 (which includes v2.0) before using CANUSB DLL.

To install .NET 3.5 support go to "Control Panel" then "Programs and Features" and then "Turn Windows features on or off" (on left menu), then you can enable Microsoft .NET Framework v3.5 which also adds support for 2.0 which is required by the CANUSB DLL. Then reboot PC and proceed.

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




