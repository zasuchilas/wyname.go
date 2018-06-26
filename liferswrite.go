package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// writing into websocket
func (l *Lifer) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		l.conn.Close() // TODO дублирует то что в read ?
		log.Println("defer write", l.hash)
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
