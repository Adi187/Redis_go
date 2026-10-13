package main

import (
	"encoding/base64"
	"fmt"
	"redis_go/store"
	"redis_go/store/kv"
	// "redis_go/store/ttl"
	// "time"
)

func main() {

	// expected :=1000

	// var wg sync.WaitGroup

	// for i:=0;i < expected;i++{
	// 	wg.Add(1)
	// 	go func(){
	// 		defer wg.Done()
	// 		increment()
	// 	}()
	// }
	// wg.Wait()

	s := kv.NewStore(0)
	s.Set("pageviews", "0")
	const hits = 1000
	cmds := make([]Command, hits)

	for i := range cmds {
		cmds[i] = Command{Op: "INCR", Key: "pageviews"}
	}

	// cmds := []Command{
	// 	{Op: "SET", Key: "env", Value: "production"},
	// 	{Op: "SET", Key: "version", Value: "0.0.1"},
	// 	{Op: "SET", Key: "debug", Value: "true"},
	// 	{Op: "GET", Key: "env"},
	// 	{Op: "SET", Key: "region", Value: "eu-west-1"},
	// 	{Op: "GET", Key: "version"},
	// }
	// s := CreateStore()
	// plainstore := kv.NewStore(20)
	RestoreBoot(s, cmds)
	got, _, _ := s.Get("pageviews")
    
	fmt.Printf("expected : %d\n", hits)
	fmt.Printf("got : %s\n", got)
    
	// fmt.Println("restored, stored size : ",plainstore.Len())
	// clonePlainstore := kv.Clone(plainstore)
	// metrics := NewMetricsMiddleware(plainstore)
	// for _, c := range cmds {
	// 	dispatch(metrics, c)
	// }
	// fmt.Println(clonePlainstore)
	// clonePlainstore.Set("a", "value")
	// encryptedValue, err := SetKeyWithEncryption(metrics, "secret", "mysecretvalue")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Encoded Value:", encryptedValue)
	// }

	// ttlStore := ttl.NewTTLStore(5 * time.Second)
	// encryptedTTLValue, err := SetKeyWithEncryption(ttlStore, "secret", "wqdwefwrfw")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Encoded Value:", encryptedTTLValue)
	// }

	// metrics.Report()

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
