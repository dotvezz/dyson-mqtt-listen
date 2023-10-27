package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dotvezz/dyson-mqtt-listen/config"
	paho "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := paho.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%s", config.Address(), "1883"))
	opts.SetClientID("mqtt-listen")
	opts.SetUsername(config.Username())
	opts.SetPassword(config.Password())

	c := paho.NewClient(opts)

	if t := c.Connect(); t.Wait() && t.Error() != nil {
		log.Fatal(t.Error())
	}

	fmt.Printf("%s: Connected to device...\n", time.Now().Format(time.TimeOnly))

	if token := c.Subscribe(fmt.Sprintf("%s/%s/status/current", config.Device(), config.Username()), 0, printMessage); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	} else {
		fmt.Printf("%s: Subscribed to %s\n", time.Now().Format(time.TimeOnly), fmt.Sprintf("%s/%s/status/current", config.Device(), config.Username()))
	}

	if token := c.Subscribe(fmt.Sprintf("%s/%s/command", config.Device(), config.Username()), 0, printMessage); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	} else {
		fmt.Printf("%s: Subscribed to %s\n", time.Now().Format(time.TimeOnly), fmt.Sprintf("%s/%s/command", config.Device(), config.Username()))
	}

	fmt.Println(fmt.Sprintf("%s: Press Ctrl+C to exit.", time.Now().Format(time.TimeOnly)))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-sig
		c.Disconnect(500)
		os.Exit(0)
	}()

	select {}
}

func printMessage(client paho.Client, msg paho.Message) {
	fmt.Printf("%s|%s: ", time.Now().Format(time.TimeOnly), msg.Topic())
	fmt.Printf("%s\n", msg.Payload())
}
