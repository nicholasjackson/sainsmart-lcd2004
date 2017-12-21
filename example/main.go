package main

import (
	"log"
	"os"
	"time"

	lcd2004 "github.com/nicholasjackson/sainsmart-lcd2004"
	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("/dev/i2c-1")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	dev := i2c.Dev{b, 0x27}
	var _ conn.Conn = &dev

	d := lcd2004.New(&dev)
	d.Init()
	d.Clear()

	for _, l := range os.Args[1:] {
		log.Println(l)
		d.Println(l)
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)
	d.DisplayOff()
}
