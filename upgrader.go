package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Println("Origin:", r.Header.Get("Origin"))
		log.Println("RemoteAddr:", r.RemoteAddr, "host:", r.Host, "Cookies:", r.Cookies(), "UserAgent:", r.UserAgent(), "Sec-WebSocket-Key:", r.Header.Get("Sec-WebSocket-Key"), "TLS:", r.TLS)
		return true
	},
}
