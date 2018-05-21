package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second    // allowed to write
	pongWait       = 60 * time.Second    // allowed to read
	pingPeriod     = (pongWait * 9) / 10 // pings period (must be less than pongWait)
	maxMessageSize = 512                 // maximum message size
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Lifer пользователь поверх websocket
type Lifer struct {
	// sector
	// subscription

	conn *websocket.Conn // websocket connection

	// buffered channel of outbound messages
	send chan []byte
}

func (l *Lifer) read() {
	defer func() {
		// l.sector.unregister <- l
		l.conn.Close() // закрываем соединение websockets
	}()

	l.conn.SetReadLimit(maxMessageSize)
	l.conn.SetReadDeadline(time.Now().Add(pongWait))
	l.conn.SetPongHandler(func(string) error { l.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := l.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// c.hub.broadcast <- message
		log.Println("message: ", string(message))
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("uh", r.URL)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lifer := &Lifer{
		conn: conn,
		send: make(chan []byte, 256),
	}

	go lifer.read()

}
