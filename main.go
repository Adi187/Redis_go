package main

import (
	"encoding/base64"
	"fmt"
	"redis_go/store"
	"redis_go/store/kv"
	"sync"

	// "redis_go/store/ttl"
	"context"
	"time"
)



func main() {
    //with normal mutex : concurreny = 2ms
	store := kv.NewStore(0)
	store.Set("name", "Adi")
	const readers = 100
	const readsEach = 50
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < readers; i++ {
		wg.Go(func() {
			for j := 0; j < readsEach; j++ {
				store.Get("name")
			}
		})
	}
	wg.Wait()
	fmt.Printf("%d readers did %d concurrent reads in %s\n",
		readers, readers*readsEach, time.Since(start).Round(time.Millisecond))
    
	 s2:=kv.NewStore(0)
	 s2.Set("pageviews","0")

	 const hits = 1000
	 cmds:= make([]Command, hits)
	 for i:= range cmds{
		cmds[i] = Command{Op:"INCR", Key:"pageviews"}
	 }
	 
	 cmds = append(cmds,
		Command{Op:"INCR",Key:""},
		Command{Op: "WAT",Key:"pageviews"},
	)
    
	RestoreBoot(s2,cmds)
	got,_ ,_ := s2.Get("pageviews")
	fmt.Printf("expected: %d\n",hits)
	fmt.Printf("got : %s\n",got)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second* 4)
    defer cancel()
	
	var wg1 sync.WaitGroup

	for i:=1;i<=3;i++{
		wg1.Add(1)
		go func(id int){
				defer wg1.Done()
                makeRequest(ctx, id)
			}(i)
		}

    wg1.Wait() //
	fmt.Println("all done")
	}


func CreateStore() store.Storer {
	s := kv.NewStore(5)
	logger := NewLoggingMiddleware(s)
	return logger
}

// SetKeyWithEncryption stores a Base64-encoded value. Base64 is encoding, not encryption.
func SetKeyWithEncryption(s store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}

	s.Len()
	value, _, err := s.Get(key)
	return value, err
}


func makeRequest(ctx context.Context,id int){
    fmt.Printf(" [req %d] making the request...\n", id)
	done:=make(chan struct{})
	go func(){
		time.Sleep(3*time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Printf("[req %d] request done\n", id)
	case <-ctx.Done():
		fmt.Printf("[req %d] request timed out : %v\n", id,ctx.Err())
	}
}
// func PopulateDefaults(s store.Storer) error{
//  defaults := map[string]string{
//     "env" : "development",
// 	"version" : "0.01",
// 	"debug" : "true",
//  }
//  for k,v:= range defaults{
// 	if err:= s.Set(k,v);err!=nil{
//       return fmt.Errorf("PopulateDefaults : %v",err)
// 	}
//  }
//  return nil
// }
