package etcd

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc-b"
)

const testKey = "/kratos/test/config"

func TestConfig(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = client.Close()
	}()
	if _, err = client.Put(context.Background(), testKey, "test config"); err != nil {
		t.Fatal(err)
	}

	source, err := New(client, WithPath(testKey))
	if err != nil {
		t.Fatal(err)
	}

	kvs, err := source.Load()
	if err != nil {
		t.Fatal(err)
	}

	if len(kvs) != 1 || kvs[0].Key != testKey || string(kvs[0].Value) != "test config" {
		t.Fatal("config error")
	}

	w, err := source.Watch()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = w.Stop()
	}()

	if _, err = client.Put(context.Background(), testKey, "new config"); err != nil {
		t.Error(err)
	}

	if kvs, err = w.Next(); err != nil {
		t.Fatal(err)
	}

	if len(kvs) != 1 || kvs[0].Key != testKey || string(kvs[0].Value) != "new config" {
		t.Fatal("config error")
	}

	if _, err := client.Delete(context.Background(), testKey); err != nil {
		t.Error(err)
	}
}

func TestExtToFormat(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = client.Close()
	}()

	tp := "/kratos/test/ext"
	tn := "a.bird.json"
	tk := tp + "/" + tn
	tc := `{"a":1}`
	if _, err = client.Put(context.Background(), tk, tc); err != nil {
		t.Fatal(err)
	}

	source, err := New(client, WithPath(tp), WithPrefix(true))
	if err != nil {
		t.Fatal(err)
	}

	kvs, err := source.Load()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(kvs))
	assert.Equal(t, tk, kvs[0].Key)
	assert.Equal(t, tc, string(kvs[0].Value))
	assert.Equal(t, "json", kvs[0].Format)
}
