# dyson-mqtt-listen
A quick tool for listening to Dyson devices over MQTT

## Installation

### Using "go install"

If you happen to have Go installed, you can just use it to install.

```bash
go install github.com/dotvezz/dyson-mqtt-listen/cmd/dyson-mqtt-listen@latest
```

### Without "go install"

You can download the executable binary file for your platform in [this repository's releases](https://github.com/dotvezz/dyson-mqtt-listen/releases).

## Running the tool

Once you've downloaded or installed the binary, run it with the correct flags for your device.

### Devices with wifi stickers:
Use the flag `-address` (The IP address of your device on your network), with `-ssid`, and `-wifi-password` using values from your sticker.

### Devices without wifi stickers:
Use the flag `-address` with `-serial`, `-password`, and `-device`. Use the [get_devices.py](https://github.com/libdyson-wg/libdyson-neon/blob/main/get_devices.py) script in libdyson-neon to find the correct values.

```bash
# if you used `go install`
dyson-mqtt-listen -address "192.168.1.207" -ssid "DYSON-NM7-US-REA2128R-475" -wifi-password "zxcvasdf" 

# if you downloaded the linux binary
chmod +r ./dyson-mqtt-listen
./dyson-mqtt-listen -address "192.168.1.207" -ssid "DYSON-NM7-US-REA2128R-475" -wifi-password "zxcvasdf" 

# if you downloaded the macOS binary
chmod +r ./dyson-mqtt-listen-darwin
./dyson-mqtt-listen -address "192.168.1.207" -ssid "DYSON-NM7-US-REA2128R-475" -wifi-password "zxcvasdf" 

# if you downloaded the Windows binary
dyson-mqtt-listen.exe -address "192.168.1.207" -ssid "DYSON-NM7-US-REA2128R-475" -wifi-password "zxcvasdf" 
```

Let it run for a few minutes. Hopefully you see something like the following:


```
00:46:35: Connected to device...
00:46:35: Subscribed to 475/NM7-US-REA2128R/status/current
00:46:35: Subscribed to 475/NM7-US-REA2128R/command
00:46:35: Press Ctrl+C to exit.
00:46:49|475/NM7-US-REA2128R/status/current: {"msg":"ENVIRONMENTAL-CURRENT-SENSOR-DATA","time":"2023-06-18T04:46:50.000Z","data":{"tact":"2930","hact":"0062","pact":"0004","vact":"0000","sltm":"OFF"}}
00:47:19|475/NM7-US-REA2128R/status/current: {"msg":"ENVIRONMENTAL-CURRENT-SENSOR-DATA","time":"2023-06-18T04:47:20.000Z","data":{"tact":"2930","hact":"0062","pact":"0004","vact":"0000","sltm":"OFF"}}
```
