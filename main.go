package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

var (
	org  = flag.String("org", "https://localhost:6969", "http service address")
	addr = flag.String("addr", ":6970", "ws service address")
)

// go run main.go --addr="whatsyourna.me:8888"

func main() {
	flag.Parse()

	http.HandleFunc("/", serveAll)

	fmt.Println(*addr, "started")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func serveAll(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Upgrade") == "websocket" {
		if r.URL.Path == "/sync" {
			synchronize(w, r)
		} else {
			// security path
			// ...
			serveWs(w, r)
		}
	} else {
		// home.html for dev
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "home.html")
	}
}

// stat is total lifers counter
var stat int64

// statplus increases lifers counter
func statplus() {
	atomic.AddInt64(&stat, 1)
}

// statminus reduces lifers counter
func statminus() {
	atomic.AddInt64(&stat, -1)
}

// startget returns stat value
func statget() string {
	st := atomic.LoadInt64(&stat)
	return strconv.FormatInt(st, 10)
}
