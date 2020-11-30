package collector

import (
	"log"
	"neophora/cli"
)

// T ...
type T struct {
	Queue chan struct {
		Key   []byte
		Value []byte
	}
	Client *cli.T
}

// Add ...
func (me *T) Add(key []byte, value []byte) {
	me.Queue <- struct {
		Key   []byte
		Value []byte
	}{
		Key:   key,
		Value: value,
	}
}

// Task ...
func (me *T) Task() {
	var result bool
	task := <-me.Queue
	if err := me.Client.Calls("DB.Put", task, &result); err != nil {
		log.Println("[!!!!][Call][DBS]", task)
		me.Queue <- task
	}
}

// Thread ...
func (me *T) Thread() {
	var result interface{}
	for {

		task := <-me.Queue
		if err := me.Client.Calls("DB.Put", task, &result); err != nil {
			log.Println("[!!!!][Call][DBS]", task)
			me.Queue <- task
		}
	}
}
