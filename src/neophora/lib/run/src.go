package run

import "log"

// T ...
type T struct {
	Service interface {
		Task()
	}
	Slaves int
}

// Run ...
func (me *T) Run() {
	log.Println("[LIB][RUN]", me.Slaves)
	for i := 0; i <= me.Slaves; i++ {
		go me.Thread()
	}
}

// Thread ...
func (me *T) Thread() {
	for {
		me.Safe()
	}
}

// Safe ...
func (me *T) Safe() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[!!!!][PANIC]", r)
		}
	}()

	me.Service.Task()
}
