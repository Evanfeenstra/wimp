package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func main() {

	c := &serial.Config{Name: findSerial(), Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("no midi found")
	} else {
		go loop(s)
	}

	InitMQTT()
}

func findSerial() string {
	contents, _ := ioutil.ReadDir("/dev")

	for _, f := range contents {
		fmt.Printf("%s\n", f.Name())
		if strings.Contains(f.Name(), "cu.usbmodem") {
			return "/dev/" + f.Name()
		}
	}
	// fail
	return ""
}

func loop(s *serial.Port) {
	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Printf("%q\n", buf[:n])
		//fmt.Println(string(buf[:n]))
		s := string(buf[:n])
		a := strings.Split(s, ":")
		last := a[len(a)-1]
		secondToLast := a[len(a)-2]
		thirdToLastWithSomeLeftover := a[len(a)-3]
		thirdToLast := thirdToLastWithSomeLeftover[len(thirdToLastWithSomeLeftover)-1:]

		msg := thirdToLast + ":" + secondToLast + ":" + last
		fmt.Println(msg)
		time.Sleep(5 * time.Millisecond)

		mqttClient.Publish("asdf", 0, false, msg)
	}

}
