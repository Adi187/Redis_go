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
	// s := CreateStore()
	plainstore:=kv.NewStore(30)
    metrics:= NewMetricsMiddleware(plainstore)

	
   
	encryptedValue, err := SetKeyWithEncryption(metrics, "secret", "mysecretvalue")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Encoded Value:", encryptedValue)
	}
    
	// ttlStore := ttl.NewTTLStore(5 * time.Second)
	// encryptedTTLValue, err := SetKeyWithEncryption(ttlStore, "secret", "wqdwefwrfw")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Encoded Value:", encryptedTTLValue)
	// }

	metrics.Report()
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



