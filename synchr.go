package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

// var upgraderForSync = websocket.Upgrader{} // use default options

func synchronize(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, []byte(now()))
	if err != nil {
		log.Println("write:", err)
	}
}

func now() string {
	n := time.Now().UnixNano()
	s := strconv.FormatInt(n, 10)
	return s
	// log.Println(n.Unix())
	// log.Println(n.UnixNano())
	// 1529048205774
	// 1529057006311727000
	// 1529057032862392400
}
