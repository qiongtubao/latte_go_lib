package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"

	"github.com/pingcap/kvproto/pkg/kvrpcpb"

	"github.com/tikv/client-go/v2/txnkv"
)

type KvClient struct {
	*txnkv.Client
}

func NewKvClient(pdAddr string) (*KvClient, error) {
	client, err := txnkv.NewClient([]string{pdAddr}, txnkv.WithAPIVersion(kvrpcpb.APIVersion_V2))
	return &KvClient{Client: client}, err
}

func (c *KvClient) Set(key string, value string) error {
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	err = tx.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	return err
}

func (c *KvClient) Get(key string) (string, error) {
	tx, err := c.Begin()
	if err != nil {
		return "", err
	}
	v, err := tx.Get(context.TODO(), []byte(key))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

type TaskResult struct {
	Result string
	Err    error
	Id     int
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
func main() {

	c, err := NewKvClient("127.0.0.1:12379")
	if err != nil {
		panic(err)
	}
	//c.Set("hello", "world")
	//v, err := c.Get("hello")
	//if err != nil {
	//	panic(err)
	//}
	//if v != "world" {
	//	panic("set wrong")
	//}
	fmt.Println("========= Task start =========")
	count := 2
	var wg sync.WaitGroup
	results := make(chan TaskResult, count)
	wg.Add(count)
	key := "test"
	i := 0
	for i < count {
		go doTask(c, key, i, &wg, results)
	}
	// 等待所有任务完成
	wg.Wait()
	close(results)
	// 处理任务结果
	r := false
	for result := range results {
		if result.Err != nil {
			fmt.Println("Task failed:", result.Id, result.Err)
		} else {
			r = true
			//fmt.Println("Task result:", result.Result)
			v, err := c.Get(key)
			if err != nil {
				panic(err)
			}
			if v != result.Result {
				fmt.Println("set fail", v, result.Result)
			}
		}
	}
	if r == false {
		fmt.Println("========= Task fail =========")
	}
}

func doTask(client *KvClient, key string, id int, wg *sync.WaitGroup, results chan<- TaskResult) {
	defer wg.Done()
	v := randomString(10)
	err := client.Set(key, v)
	if err != nil {
		results <- TaskResult{Id: id, Err: err}
		return
	}
	results <- TaskResult{Id: id, Result: v}
}
