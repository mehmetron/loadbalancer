package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/coreos/etcd/clientv3"
)


type Env struct {
	etcd   clientv3.KV
}

var (
	dialTimeout    = 5 * time.Second
)

// Create key:value pair
func (env *Env) Create(w http.ResponseWriter, r *http.Request) {

	// Serialize request body
	var toCreate CreateSandbox
	err := json.NewDecoder(r.Body).Decode(&toCreate)
	check(err)

	// Generate subdomains
	rand.Seed(time.Now().UnixNano())
	demoSubdomain := RandStringRunes(10)
	apiSubdomain := fmt.Sprintf("%sgalatatower", demoSubdomain)
	fmt.Println("gen subdomains", demoSubdomain, apiSubdomain)

	// Find ports
	port1, port2 := GeneratePorts()
	demoPort := strconv.Itoa(port1)
	apiPort := strconv.Itoa(port2)

	fmt.Println("demoSubdomain, apiSubdomain, demoPort, apiPort, toCreate.LangID", demoSubdomain, apiSubdomain, demoPort, apiPort, toCreate.LangID)

	// Put demoPort and apiPort into etcd
	putDemoResp, err := CreateKVs(env, demoSubdomain, demoPort)
	fmt.Println("CreateKVs putDemoResp: ", putDemoResp)
	check(err)

	putApiResp, err := CreateKVs(env, apiSubdomain, apiPort)
	fmt.Println("CreateKVs putApiResp: ", putApiResp)
	check(err)

	// // Put demoPort in etcd
	// putResp, err := env.etcd.Put(context.TODO(), demoSubdomain, fmt.Sprintf("http://localhost:%s", demoPort))
	// fmt.Println("Added :", putResp)
	// check(err)

	// // Put apiPort in etcd
	// putResp2, err := env.etcd.Put(context.TODO(), apiSubdomain, fmt.Sprintf("http://localhost:%s", apiPort))
	// fmt.Println("Added :", putResp2)
	// check(err)

	// Create docker container
	res := Create(env, demoPort, apiPort, toCreate.LangID)
	fmt.Println("res", res)

	fmt.Fprintf(w, "Create World! %s %s", demoSubdomain, apiSubdomain, res)

}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: dialTimeout,
	})
	check(err)
	defer cli.Close()
	kv := clientv3.NewKV(cli)

	env := &Env{etcd: kv, docker: docker}

}
