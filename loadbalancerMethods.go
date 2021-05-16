package main

// Algorithm implementations taken from https://github.com/smallnest/weighted
import (
	"fmt"
	"net/url"
)

// W is a interface that implement a weighted round robin algorithm.
type W interface {
	// Next gets next selected item.
	// Next is not goroutine-safe. You MUST use the snchronization primitive to protect it in concurrent cases.
	Next() (item interface{})
	// Add adds a weighted item for selection.
	Add(item interface{}, weight int)

	// All returns all items.
	All() map[interface{}]int

	// RemoveAll removes all weighted items.
	RemoveAll()
	// Reset resets the balancing algorithm.
	Reset()
}

func weightedRoundRobin() *url.URL {
	w := &RRW{}
	w.Add("http://milk.com/", 4)
	w.Add("http://milk.com/fuzzboy/", 2)
	w.Add("http://milk.com/value/", 3)
	w.Add("http://milk.com/hourlykitten/", 1)

	rawUrl := w.Next().(string)
	parsedurl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	}
	return parsedurl
}

var w *RandW

func init() {
	w = NewRandW()
	w.Add("http://milk.com/", 4)
	w.Add("http://milk.com/fuzzboy/", 2)
	w.Add("http://milk.com/value/", 3)
	w.Add("http://milk.com/hourlykitten/", 1)
	w.Add("http://localhost:8006/", 3)
	w.Add("http://localhost:8005/", 2)
	w.Add("http://localhost:8004/", 1)
	w.Remove("http://milk.com/fuzzboy/")
}

func weightedRandom() *url.URL {

	fmt.Println("51 ", w.items)

	rawUrl := w.Next().(string)
	parsedurl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	}
	return parsedurl
}

func weightedSmooth() *url.URL {
	w := &SW{}
	w.Add("http://milk.com/", 4)
	w.Add("http://milk.com/fuzzboy/", 2)
	w.Add("http://milk.com/value/", 3)
	w.Add("http://milk.com/hourlykitten/", 1)

	rawUrl := w.Next().(string)
	parsedurl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
	}
	return parsedurl
}
