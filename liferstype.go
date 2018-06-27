package main

import "github.com/gorilla/websocket"

// Lifer пользователь поверх websocket
type Lifer struct {
	// sector
	// subscription
	hash string

	conn *websocket.Conn // websocket connection
	send chan []byte     // buffered channel of outbound messages

	initsamf bool // is samf values are set
	initgps  bool // is gps values are set
	started  bool // is lifer already launched in sector
	// insecur bool

	samf    int
	sex     int
	age     int
	sa      int
	filter  int
	filters []int
	mark    string

	gps *Gps

	// text represantations for logs
	inboundLat  string
	inboundLon  string
	inboundSamf string

	secache       map[string]*Sector // lifers cache of *Sectors
	membership    string             // current lifer member sector
	subscriptions map[string]bool    // current lifer subscribe sectors
}
