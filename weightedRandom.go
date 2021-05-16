package main

import (
	"math/rand"
	"sync"
	"time"
)

// randWeighted is a wrapped weighted item that is used to implement weighted random algorithm.
type randWeighted struct {
	Item   interface{}
	Weight int
}

// RandW is a struct that contains weighted items implement weighted random algorithm.
type RandW struct {
	items        []*randWeighted
	n            int
	sumOfWeights int
	r            *rand.Rand
	sync.Mutex
}

// NewRandW creates a new RandW with a random object.
func NewRandW() *RandW {
	return &RandW{r: rand.New(rand.NewSource(int64(time.Now().UnixNano())))}
}

// Next returns next selected item.
func (rw *RandW) Next() (item interface{}) {
	if rw.n == 0 {
		return nil
	}
	if rw.sumOfWeights <= 0 {
		return nil
	}
	randomWeight := rw.r.Intn(rw.sumOfWeights) + 1
	for _, item := range rw.items {
		randomWeight = randomWeight - item.Weight
		if randomWeight <= 0 {
			return item.Item
		}
	}

	return rw.items[len(rw.items)-1].Item
}

// Add adds a weighted item for selection.
func (rw *RandW) Add(item interface{}, weight int) {
	rItem := &randWeighted{Item: item, Weight: weight}
	rw.items = append(rw.items, rItem)
	rw.sumOfWeights += weight
	rw.n++
}

// Finds an item.
func (rw *RandW) Find(item interface{}) int {
	for index, x := range rw.items {
		if item == x.Item {
			return index
		}
	}
	return len(rw.items) - 1
}

// Removes an item.
func (rw *RandW) Remove(item interface{}) {
	index := rw.Find(item)
	rw.sumOfWeights -= rw.items[index].Weight

	//// Supposedly more efficient but doesn't work
	//rw.items[len(rw.items)-1], rw.items[index] = rw.items[index], rw.items[len(rw.items)-1]
	//rw.items = rw.items[:len(rw.items)-1]

	rw.items = append(rw.items[:index], rw.items[index+1:]...)
	rw.n--
}

// All returns all items.
func (rw *RandW) All() map[interface{}]int {
	m := make(map[interface{}]int)
	for _, i := range rw.items {
		m[i.Item] = i.Weight
	}
	return m
}

// RemoveAll removes all weighted items.
func (rw *RandW) RemoveAll() {
	rw.items = make([]*randWeighted, 0)
	rw.r = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
}

// Reset resets the balancing algorithm.
func (rw *RandW) Reset() {
	rw.r = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
}
