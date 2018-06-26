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
		l.conn.Close() // закрываем соединение websockets
		statminus()
		l.awayFromMembers()
		l.unsubscribeEverywhere()
		// -> log : part, hash, samf, sex, age, lat, lon
		log.Println("C," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
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
						l.samf = inbsamf
						l.inboundSamf = inb[1]
						l.sex, l.age, l.sa, l.filter, l.mark = desamf(inbsamf)
						if l.started {
							// reconnect

						} else {
							l.initsamf = true
							if l.initgps == true {
								l.connectFirst()
							}
						}
					}
				}
			case codeGpsData:
				if len == 3 {
					inbla, ela := strconv.ParseFloat(inb[1], 64)
					inblo, elo := strconv.ParseFloat(inb[2], 64)
					if ela == nil && elo == nil && (inbla != l.gps.lat || inblo != l.gps.lon) {
						g, e := newGps(inbla, inblo)
						if e == nil {
							l.gps = g
							l.inboundLat = inb[1]
							l.inboundLon = inb[2]
							nmember, nsubscr, e := l.gps.calculate() // new member and subscribe sectors
							if e == nil {
								// secache update
								for subsc := range nsubscr {
									if _, found := l.secache[subsc]; !found {
										l.secache[subsc] = camp.sector(subsc)
									}
								}

								if l.started {
									// move

									// проверить не изменился ли набор секторов
									// for _, sec := range secs {
									// 	// camp.sector(sec) RMUTEX
									// 	// проверить не изменился ли набор секторов
									// 	log.Println(sec)

									// }
								} else {
									l.cmember = nmember
									l.csubscr = nsubscr
									l.initgps = true
									if l.initsamf == true {
										l.connectFirst() // connect first -> l.started = true and log B
									}
								}
							}
						}
					}
				}
			case codeGlobRequest:
				if len == 5 {
					tala, etala := strconv.ParseFloat(inb[1], 64)
					talo, etalo := strconv.ParseFloat(inb[2], 64)
					tcla, etcla := strconv.ParseFloat(inb[3], 64)
					tclo, etclo := strconv.ParseFloat(inb[4], 64)
					if etala == nil && etalo == nil && etcla == nil && etclo == nil {
						log.Println(tala, talo, tcla, tclo)

					}
				}
			default:
				// l.send <- []byte("18," + string(message))
				// c.hub.broadcast <- message
			}
		}
	}
}
