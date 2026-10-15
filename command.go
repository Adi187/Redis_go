package main

import (
	"redis_go/store"
	"strings"
	"sync"
	"fmt"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) error {
	switch strings.ToUpper(c.Op) {
	case "SET":
		return s.Set(c.Key, c.Value)
	case "GET":
		_, _, err := s.Get(c.Key)
		return err
	case "INCR":
		_, err := s.Incr(c.Key)
		return err
	default:
		return fmt.Errorf("unknown command: %q", c.Op)
	}
}

// RestoreBoot implements a slice of commands against a store
func RestoreBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	errs := make(chan error,len(cmds))
	for _, c := range cmds {
		wg.Go(func() {
			errs <- dispatch(s, c)
		})
	}
    
	go func(){
		wg.Wait()
		close(errs)
	}()

	for err := range errs{
		if err!= nil {
			fmt.Printf("Error: %v ",err)
		}
	}
}
