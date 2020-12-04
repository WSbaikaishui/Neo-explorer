package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

// T ...
type T struct {
	IP        string
	ExecCmd   string
	StartFrom int
	last      time.Time
	lock      sync.Mutex
	exe       *exec.Cmd
}

// Call ...
func (me *T) Call() {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.last = time.Now()
}

// Commit ...
func (me *T) Commit(n int) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.StartFrom = n
}

// Task ...
func (me *T) Task() {
	me.Reset()
	log.Println("[Service][Daemon] Running ...")
	me.exe.Stdin = os.Stdin
	me.exe.Stdout = os.Stdout
	err := me.exe.Run()
	if err != nil {
		log.Println("[Service][Daemon]", err)
	}
}

// Reset ...
func (me *T) Reset() {
	me.lock.Lock()
	defer me.lock.Unlock()
	cmd := fmt.Sprintf("export STARTFROM=%d && %s", me.StartFrom, me.ExecCmd)
	me.exe = exec.Command("bash", "-c", cmd)
	log.Println("[Service][Daemon][CMD]", cmd)
}
