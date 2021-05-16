package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"time"
)

// isAlive checks whether a backend is Alive by establishing a TCP connection
func isBackendAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return false
	}
	_ = conn.Close()
	return true
}

// healthCheck runs a routine for check status of the backends every 2 mins
func healthCheck(w *RandW) {

	//t := time.NewTicker(time.Minute * 2)
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			log.Println("Starting health check...")
			for _, b := range w.items {
				status := "up"
				str := fmt.Sprintf("%v", b.Item)
				parsedUrl, err := url.Parse(str)
				if err != nil {
					fmt.Println(err)
				}
				alive := isBackendAlive(parsedUrl)
				if !alive {
					w.Lock()
					w.Remove(b.Item)
					w.Unlock()
					status = "down"
				}
				log.Printf("%s [%s]\n", b.Item, status)
			}
			log.Println("Health check completed")
		}
	}
}
