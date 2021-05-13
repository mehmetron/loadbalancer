package main

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3"
)

// Delete all key:values with _subdomain_ as prefix
func DeleteKV(env *Env, subdomain string) (*clientv3.DeleteResponse, error) {
	delResp, err := env.etcd.Delete(context.TODO(), subdomain, clientv3.WithPrefix())
	return delResp, err
}

func CreateKVs(env *Env, subdomain, port string) (*clientv3.PutResponse, error) {
	putResp, err := env.etcd.Put(context.TODO(), subdomain, fmt.Sprintf("http://localhost:%s", port))
	return putResp, err
}
