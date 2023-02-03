package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"
)

func main() {
	fmt.Println(os.Getenv("IOTHUB_DEVICE_CONNECTION_STRING"))
	c, err := iotdevice.NewFromConnectionString(
		iotmqtt.New(), os.Getenv("IOTHUB_DEVICE_CONNECTION_STRING"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to the iothub
	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	s := fmt.Sprintf("%d", time.Now().UnixNano())
	v, err := c.UpdateTwinState(context.Background(), map[string]interface{}{
		"ts": s,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("version: %d", v)

	// send a device-to-cloud message
	if err = c.SendEvent(context.Background(), []byte("hello")); err != nil {
		log.Fatal(err)
	}
}
