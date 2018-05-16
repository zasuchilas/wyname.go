package main

import (
	"flag"
	"fmt"
	"wyname/utils"
)

var (
	addr = flag.String("addr", "localhost", "Domain address")
	port = flag.Int("port", 6969, "Port")
	ssl  = flag.Bool("ssl", false, "Use SSL")
	// Origin for secure
	Origin string
)

// go run main.go --ssl=true --addr="whatsyourna.me" --port=8888

func main() {
	flag.Parse()
	fmt.Println("wyname", "ok")
	fmt.Println("addr", *addr)
	fmt.Println("port", *port)
	fmt.Println("ssl", *ssl)

	Origin = origing()
	fmt.Println("origin", Origin)
	utils.Abc("qqq")

}

func origing() string {
	var o string
	if *ssl {
		o = "https://" + *addr
	} else {
		o = "http://" + *addr
	}
	return o
}
