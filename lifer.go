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
		// TODO здесь нужно разбирать коды (inbox from browser)
		log.Println("message: ", string(message))
		l.send <- []byte("ko: " + string(message))
	}
}

func (l *Lifer) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		l.conn.Close() // TODO дублирует то что в read ?
	}()

	for {
		select {
		case message, ok := <-l.send: // читаем из буфера записи лайфера
			l.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// канал l.send был закрыт из сектора
				l.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := l.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// добавляем остальные сообщения в очереди l.send (если есть)
			// к данному сообщению (чтобы одним блоком ушло все из буфера лайфера)
			// логично: ядер конечное число, и горутины ждут своей очереди
			// поэтому используем возможность, чтобы не стоять снова
			n := len(l.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-l.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C: // обработка ping/pong
			l.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := l.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
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
	go lifer.write()
}