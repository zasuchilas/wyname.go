package main

import (
	"fmt"
	"log"
	"net/http"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lifer := &Lifer{
		// hash:     fmt.Sprint(&conn)[2:],
		hash:     fmt.Sprint(&conn),
		conn:     conn,
		send:     make(chan []byte, 512),
		initsamf: false,
		initgps:  false,
		started:  false,
		secache:  make(map[string]*Sector, 21),
		mutex:    new(sync.RWMutex),
	}

	statplus()

	lifer.send <- []byte("12," + lifer.hash)

	// -> log : part, hash, remote_addr
	log.Println("A," + lifer.hash + "," + r.RemoteAddr)

	go lifer.read()
	go lifer.write()
}
