package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
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

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("conn path:", r.URL)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lifer := &Lifer{
		hash:     fmt.Sprint(&conn)[2:],
		conn:     conn,
		send:     make(chan []byte, 256),
		initsamf: false,
		initgps:  false,
		started:  false,
		secache:  make(map[string]*Sector, 21),
		mutex:    new(sync.RWMutex),
	}

	statplus()

	lifer.send <- []byte("12," + lifer.hash)

	// -> log : part, hash, remote_addr, host, cookies, useragent, wskey, tls
	log.Println("A," + lifer.hash + "," + r.RemoteAddr + "," + strings.Replace(r.Host, ",", " ", -1) + "," + strings.Replace(fmt.Sprint(r.Cookies()), ",", " ", -1) + "," + strings.Replace(r.UserAgent(), ",", "", -1) + "," + r.Header.Get("Sec-WebSocket-Key") + "," + strings.Replace(fmt.Sprint(r.TLS), ",", "", -1))

	go lifer.read()
	go lifer.write()
}
