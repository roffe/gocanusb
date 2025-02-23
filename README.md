# goCANUSB

Golang implementation for using the LAWICEL CANUSB Windows DLL

## Installation

You need the DLL for your CANUSB which can be found at [https://www.canusb.com/support/canusb-support/](https://www.canusb.com/support/canusb-support/)

Download the appropriate setup based on your target GOARCH 32-bit = 386, 64-bit = amd64

**Make sure .NET support is installed**: Windows 10 users will need to enable .NET framework 3.5 (which includes v2.0) before using CANUSB DLL.

To install .NET 3.5 support go to “Control Panel” then “Programs and Features” and then “Turn Windows features on or off” (on left menu), then you can enable Microsoft .NET Framework v3.5 which also adds support for 2.0 which is required by the CANUSB DLL. Then reboot PC and proceed with the following step.

## Disclaimer

**LAWICEL CANUSB and LAWICEL CAN232 are trademarks of ELMICRO Computer GmbH & Co. KG**

## Showcase

[examples](https://github.com/roffe/gocanusb/tree/master/examples)

Used in [goCAN](https://github.com/roffe/gocan)


