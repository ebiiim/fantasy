package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var ip string
	var port int
	var dir string
	var help bool
	flag.BoolVar(&help, "h", false, "show this message and exit")
	flag.StringVar(&ip, "ip", "localhost", "IP")
	flag.IntVar(&port, "port", 8080, "port")
	flag.StringVar(&dir, "dir", ".", "base directory")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	fmt.Printf("Serving %s on http://%s:%d\n", dir, ip, port)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), http.FileServer(http.Dir(dir)))

	<-ctx.Done()
}
