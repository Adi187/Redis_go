package main

import (
	"log"
	"redis_go/store"
	"strings"
	"sync"
)

type Command struct {
	Op    string
	Key   string
	Value string

}


func dispatch(s store.Storer, c Command) {
	switch strings.ToUpper(c.Op) {
	case "SET":
		log.Printf("SET key %s",c.Key)
		s.Set(c.Key,c.Value)
	case "GET":
		s.Get(c.Key)
	case "INCR":
		s.Incr(c.Key)
	}
}


// RestoreBoot implements a slice of commands against a store
func RestoreBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	for _, c := range cmds {
     wg.Add(1)
	 go func(){
		//defers the code to the end of 
		//the function
	   defer wg.Done()
       dispatch(s,c)
	 }()
	}

    wg.Wait()

}
