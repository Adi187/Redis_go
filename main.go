package main

import (
	"encoding/base64"
	"fmt"
	"redis_go/store"
	"redis_go/store/kv"
	"sync"
	// "redis_go/store/ttl"
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
