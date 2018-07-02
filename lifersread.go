package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// reading from websocket
func (l *Lifer) read() {
	defer func() {
		statminus()
		l.awayFromMembers()
		l.unsubscribeEverywhere(false)
		l.logC()
		// we do not send anything to the client via websocket
		l.conn.Close() // close the websocket connection
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
		// inbox from browser
		inb := strings.Split(string(message), ",")
		len := len(inb)
		if len > 0 {
			switch inb[0] {
			case codeStatsRequest:
				st := statget()
				l.send <- []byte(codeStatsResponse + "," + st)
			case codeSamf:
				if len == 2 {
					inbsamf, e := strconv.Atoi(inb[1])
					if e == nil && inbsamf != l.samf {
						if l.started {
							// reconnect
							l.awayFromMembers()               // delete membership
							l.unsubscribeEverywhere(true)     // delete all subscribtions
							l.changeSamfData(inbsamf, inb[1]) // ! change samf data
							l.connectFirst()                  // connect with new samf data
						} else {
							l.changeSamfData(inbsamf, inb[1]) // ! change samf data
							l.initsamf = true
							if l.initgps == true {
								l.connectFirst()
								l.logB()
							}
						}
					}
				}
			case codeGpsData:
				if len == 3 {
					inbla, ela := strconv.ParseFloat(inb[1], 64)
					inblo, elo := strconv.ParseFloat(inb[2], 64)
					if ela == nil && elo == nil && (l.gps == nil || (l.gps != nil && (inbla != l.gps.lat || inblo != l.gps.lon))) {
						g, e := newGps(inbla, inblo)
						if e == nil {
							membership, subscriptions, e := g.calculate() // new member and subscribe sectors
							if e == nil {
								// save new data
								l.mutex.Lock()
								l.gps = g
								l.inboundLat = inb[1]
								l.inboundLon = inb[2]
								l.mutex.Unlock()
								// secache update
								for subsc := range subscriptions {
									if _, found := l.secache[subsc]; !found {
										l.secache[subsc] = camp.sector(subsc)
									}
								}
								// further processing
								if l.started { // if this is a continuation
									// notifications
									if membership == l.membership { // move
										l.secache[membership].broadcast <- createMoveJob(l)
									} else { // change membership
										l.awayFromMembers()
										l.membership = membership
										l.secache[membership].broadcast <- createComeJob(l)
									}
									// subscriptions
									for oldSubscrSector := range l.subscriptions { // remove needless subscriptions
										if _, found := subscriptions[oldSubscrSector]; !found {
											l.secache[oldSubscrSector].broadcast <- createUnsubscribeJob(l, true)
										}
									}
									for newSubscrSector := range subscriptions { // add new subscriptions
										if _, found := l.subscriptions[newSubscrSector]; !found {
											l.secache[newSubscrSector].broadcast <- createSubscribeJob(l)
										}
									}
									l.subscriptions = subscriptions
								} else { // if this is the beginning
									// save additional data
									l.membership = membership
									l.subscriptions = subscriptions
									l.initgps = true
									if l.initsamf == true { // if all data is ready
										l.connectFirst() // connect first -> l.started = true
										l.logB()
									}
								}
							}
						}
					}
				}
			case codeGlobRequest:
				if len == 6 {
					tala, etala := strconv.ParseFloat(inb[1], 64)
					talo, etalo := strconv.ParseFloat(inb[2], 64)
					tcla, etcla := strconv.ParseFloat(inb[3], 64)
					tclo, etclo := strconv.ParseFloat(inb[4], 64)
					globReqCode := inb[5]
					if etala == nil && etalo == nil && etcla == nil && etclo == nil {
						log.Println(tala, talo, tcla, tclo)
						a, ea := newGps(tala, talo)
						c, ec := newGps(tcla, tclo)
						if ea == nil && ec == nil {
							globsectors := screen(a, c)
							for globsec := range globsectors {
								globSector, found := l.secache[globsec]
								if found == false {
									globSector = camp.sector(globsec)
									l.secache[globsec] = globSector
								}
								globSector.broadcast <- createGlobJob(l, globReqCode)
							}
						}
					}
				} // codeGlobRequest
			} // switch
		}
	}
}
