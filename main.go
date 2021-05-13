package main

import (
	"time"
	"github.com/coreos/etcd/clientv3"
)


type Env struct {
	etcd   clientv3.KV
}

var (
	dialTimeout    = 5 * time.Second
)

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
