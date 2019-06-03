package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheThingsIndustries/mystique/pkg/apex"
	"github.com/TheThingsIndustries/mystique/pkg/log"
	mqttnet "github.com/TheThingsIndustries/mystique/pkg/net"
	"github.com/TheThingsIndustries/mystique/pkg/server"
)

var (
	ctx    = context.Background()
	logger = apex.Log
	DEBUG  = false
	HTTP   = true
	TCP    = true
)

func InitMQTT() {

	mqttPort := os.Getenv("PORT")
	if mqttPort == "" {
		mqttPort = "1883"
	}

	httpPort := "1880"

	go StartClient(mqttPort)

	//fmt.Println(auth.Provider.Decode)

	if DEBUG == true {
		apex.SetLevelFromString("debug")
	}
	ctx = log.NewContext(ctx, logger)
	s := server.New(ctx)
	RunServer(s, mqttPort, httpPort)
}

// RunServer the server
func RunServer(s server.Server, mqttPort, httpPort string) {

	if TCP && mqttPort != "" {
		logger.WithField("address", ":"+mqttPort).Info("Starting MQTT server")
		lis, err := mqttnet.Listen("tcp", ":"+mqttPort)
		if err != nil {
			logger.WithError(err).Fatal("Could not start MQTT server")
		}
		defer lis.Close()

		go func() {
			for {
				conn, err := lis.Accept()
				if err != nil {
					logger.WithError(err).Error("Could not accept connection")
					return
				}
				go s.Handle(conn)
			}
		}()
	}

	wss := mqttnet.Websocket(s.Handle)
	mux := http.NewServeMux()
	mux.Handle("/mqtt", wss)

	if _, err := os.Stat("index.html"); err == nil {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("SERVE frontend")
			http.ServeFile(w, r, "index.html")
		})
	}

	if HTTP && httpPort != "" {
		logger.WithField("address", ":"+httpPort).Info("Starting HTTP+ws server")
		lis, err := net.Listen("tcp", ":"+httpPort)
		if err != nil {
			logger.WithError(err).Fatal("Could not start HTTP+ws server")
		}
		defer lis.Close()

		go func() {
			err := http.Serve(lis, mux)
			if err != nil {
				logger.WithError(err).Error("Could not serve HTTP+ws")
			}
		}()
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal := (<-sigChan).String()
	logger.WithField("signal", signal).Info("Signal received")
}
