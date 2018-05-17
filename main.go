package main

import (
	"flag"
	"fmt"
	"wyname/kernel"
	"wyname/utils"
)

var (
	addr   = flag.String("addr", "localhost", "Domain address")
	port   = flag.Int("port", 6969, "Port")
	ssl    = flag.Bool("ssl", false, "Use SSL")
	origin string
)

// go run main.go --ssl=true --addr="whatsyourna.me" --port=8888

func main() {

	// loading
	configsrv()

	utils.Abc("qqq")

}

func configsrv() {
	flag.Parse()
	fmt.Println("wyname", "ok")
	fmt.Println("addr", *addr)
	fmt.Println("port", *port)
	fmt.Println("ssl", *ssl)

	origin = origing()
	fmt.Println("origin", origin)

	kernel.Load(*addr, *port, *ssl)

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
