package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"net"
	"os"

	"github.com/tarm/serial"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {

	ServerAddr,err := net.ResolveUDPAddr("udp","localhost:9000")
	CheckError(err)
	fmt.Println(ServerAddr)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
	

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
		//fmt.Printf("%s\n",f.Name())
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
		fmt.Printf("%q\n", buf[:n])

		time.Sleep(5 * time.Millisecond)

		// fmt.Println(mqttClient)
		// mqttClient.Publish(topic, 0, false, string(out))
	}

}