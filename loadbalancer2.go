package main

//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"net/url"
//	"sync"
//	"sync/atomic"
//)
//
//// Backend holds the data about a server
//type Backend struct {
//	URL          *url.URL
//	Alive        bool
//	mux          sync.RWMutex
//	//ReverseProxy *httputil.ReverseProxy
//}
//
//// SetAlive for this backend
//func (b *Backend) SetAlive(alive bool) {
//	b.mux.Lock()
//	b.Alive = alive
//	b.mux.Unlock()
//}
//
//// IsAlive returns true when backend is alive
//func (b *Backend) IsAlive() (alive bool) {
//	b.mux.RLock()
//	alive = b.Alive
//	b.mux.RUnlock()
//	return
//}
//
//// ServerPool holds information about reachable backends
//type ServerPool struct {
//	backends []*Backend
//	current  uint64
//}
//
//
//
//// NextIndex atomically increase the counter and return an index
//func (s *ServerPool) NextIndex() int {
//	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
//}
//
//// MarkBackendStatus changes a status of a backend
//func (s *ServerPool) MarkBackendStatus(backendUrl *url.URL, alive bool) {
//	for _, b := range s.backends {
//		if b.URL.String() == backendUrl.String() {
//			b.SetAlive(alive)
//			break
//		}
//	}
//}
//
//// GetNextPeer returns next active peer to take a connection
//func (s *ServerPool) GetNextPeer() *Backend {
//	// loop entire backends to find out an Alive backend
//	next := s.NextIndex()
//	l := len(s.backends) + next // start from next and move a full cycle
//	for i := next; i < l; i++ {
//		idx := i % len(s.backends) // take an index by modding
//		if s.backends[idx].IsAlive() { // if we have an alive backend, use it and store if its not the original one
//			if i != next {
//				atomic.StoreUint64(&s.current, uint64(idx))
//			}
//			return s.backends[idx]
//		}
//	}
//	return nil
//}
//
//
//
//
//var serverPool ServerPool
//
//var tokens = []string{
//	"http://localhost:8004/",
//	"http://localhost:8005/",
//	"http://localhost:8006/",
//}
//
//
//func lb() {
//
//
//	box := serverPool.GetNextPeer()
//	resp, _ := http.Get(box.URL.String())
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//	fmt.Println("105", string(body))
//
//}
//
//func init() {
//	for _, tok := range tokens {
//		serverUrl, err := url.Parse(tok)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		serverPool.backends = append(serverPool.backends, &Backend{
//			URL:          serverUrl,
//			Alive:        true,
//		})
//
//	}
//}
//
//// Portions taken from https://github.com/kasvith/simplelb/blob/master/main.go
//func Loadbalancer2() {
//
//	lb()
//
//	// start health checking
//	//go healthCheck()
//
//
//}
