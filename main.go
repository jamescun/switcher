package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	listenAddress  = flag.String("listen", ":80", "Server Listen Address")
	sshAddress     = flag.String("ssh", "127.0.0.1:22", "SSH Server Address")
	defaultAddress = flag.String("default", "127.0.0.1:8080", "Default Server Address")
)

func usage() {
	fmt.Println("Switcher 1.0.1")
	fmt.Println("usage: switcher [options]\n")

	fmt.Println("Options:")
	fmt.Println("  --listen   <:80>            Server Listen Address")
	fmt.Println("  --ssh      <127.0.0.1:22>   SSH Server Address")
	fmt.Println("  --default  <127.0.0.1:8080>  Default Server Address\n")

	fmt.Println("Examples:")
	fmt.Println("  To serve SSH(127.0.0.1:22) and HTTP(127.0.0.1:8080) on port 80")
	fmt.Println("  $ switcher\n")

	fmt.Println("  To serve SSH(127.0.0.1:2222) and HTTPS(127.0.0.1:443) on port 443")
	fmt.Println("  $ switcher --listen :443 --ssh 127.0.0.1:2222 --default 127.0.0.1:443")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	mux := NewMux()

	mux.Handle(SSH(*sshAddress))
	mux.Handle(TCP(*defaultAddress))

	log.Printf("[INFO] listen: %s\n", *listenAddress)
	err := mux.ListenAndServe(*listenAddress)
	if err != nil {
		log.Fatalf("[FATAL] listen: %s\n", err)
	}
}
