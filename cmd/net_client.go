package cmd

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	dialTout      = 5
	httpTout      = 10
	handshakeTout = 10
)

var (
	netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: dialTout * time.Second,
		}).Dial,
		TLSHandshakeTimeout: handshakeTout * time.Second,
	}

	netClient = &http.Client{
		Timeout:   time.Second * httpTout,
		Transport: netTransport,
	}

	serverAddr = ""
)

// SetServerAddr - set addr using host and port
func SetServerAddr(HTTPHost string, HTTPPort int) {
	serverAddr = "http://" + HTTPHost + ":" + strconv.Itoa(HTTPPort)
}
