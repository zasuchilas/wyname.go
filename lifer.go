package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Lifer пользователь поверх websocket
type Lifer struct {
	// sector
	// subscription
	hash string

	conn *websocket.Conn // websocket connection
	send chan []byte     // buffered channel of outbound messages

	initsamf bool
	initgps  bool
	started  bool
	// insecur bool

	samf   int
	sex    int
	age    int
	sa     int
	filter int
	mark   string

	gps *Gps

	// text represantations for logs
	inboundLat  string
	inboundLon  string
	inboundSamf string

	secache map[string]*Sector // lifers cache of *Sectors
	cmember string             // current lifer member sector
	csubscr map[string]bool    // current lifer subscribe sectors
	nmember string             // new inbound lifer member sector
	nsubscr map[string]bool    // new inbound lifersubsribe sectors
}

// reading from websocket
func (l *Lifer) read() {
	defer func() {
		l.conn.Close() // закрываем соединение websockets
		statminus()
		// away func
		if l.cmember != "" {
			if awaySector, found := l.secache[l.cmember]; found {
				awaySector.broadcast <- newawayjob(l)
			}
		}
		if len(l.csubscr) > 0 {
			awayUnsubscrJob := newUnsubscribeJob(l)
			for awayname := range l.csubscr {
				if awaySubs, found := l.secache[awayname]; found {
					awaySubs.broadcast <- awayUnsubscrJob
				}
			}
		}
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
								// connect first

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
										// connect first -> l.started = true
										l.secache[l.cmember].broadcast <- newcomejob(l)
										for secname := range l.csubscr {
											l.secache[secname].broadcast <- newSubscribeJob(l)
										}
										l.started = true
										// -> log : part, hash, samf, sex, age, lat, lon
										log.Println("B," + l.hash + "," + l.inboundSamf + "," + strconv.Itoa(l.sex) + "," + strconv.Itoa(l.age) + "," + l.inboundLat + "," + l.inboundLon)
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
	}

	statplus()

	lifer.send <- []byte("12," + lifer.hash)

	// -> log : part, hash, remote_addr, host, cookies, useragent, wskey, tls
	log.Println("A," + lifer.hash + "," + r.RemoteAddr + "," + strings.Replace(r.Host, ",", " ", -1) + "," + strings.Replace(fmt.Sprint(r.Cookies()), ",", " ", -1) + "," + strings.Replace(r.UserAgent(), ",", "", -1) + "," + r.Header.Get("Sec-WebSocket-Key") + "," + strings.Replace(fmt.Sprint(r.TLS), ",", "", -1))

	go lifer.read()
	go lifer.write()
}
