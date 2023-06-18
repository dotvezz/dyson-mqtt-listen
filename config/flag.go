package config

import (
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"log"
	"strings"
)

var (
	address    *string
	password   *string
	deviceType *string
	serial     *string
	wifiPass   *string
	ssid       *string
)

func init() {
	address = flag.String("address", "", "(Required) The address of the device to connect to.")
	password = flag.String("password", "", "Only needed if using MQTT credentials")
	deviceType = flag.String("device", "", "Only needed if using MQTT credentials")
	serial = flag.String("serial", "", "Only needed if using MQTT credentials")
	wifiPass = flag.String("wifi-password", "", "The Wi-Fi password from the product sticker")
	ssid = flag.String("ssid", "", "The SSID from the product sticker")
	flag.Parse()
}

func Address() string {
	if *address == "" {
		log.Fatal("-address is required")
	}
	return *address
}

func Username() string {
	name := *serial
	if name == "" {
		parts := strings.Split(strings.ToUpper(*ssid), "-")
		name = strings.Join(parts[1:len(parts)-1], "-")
	}
	return name
}

func Password() string {
	pass := *password

	if pass == "" {
		hash := sha512.Sum512([]byte(*wifiPass))
		pass = base64.StdEncoding.EncodeToString(hash[:])
	}

	return pass
}

func Device() string {
	device := *deviceType
	if device == "" {
		parts := strings.Split(strings.ToUpper(*ssid), "-")
		device = parts[len(parts)-1]
	}
	return device
}
