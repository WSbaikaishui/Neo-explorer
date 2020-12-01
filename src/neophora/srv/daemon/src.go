package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
	"time"
)

// T ...
type T struct {
	IP        string
	ExecPath  string
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
	err := me.exe.Run()
	if err != nil {
		log.Println("[Service][Daemon]", err)
	}
}

// Reset ...
func (me *T) Reset() {
	me.lock.Lock()
	defer me.lock.Unlock()
	dir := path.Dir(me.ExecPath)
	{
		d := path.Join(dir, "Chain_00746E41")
		cmd := fmt.Sprintf("rm -rf %s", d)
		log.Println("[Service][Daemon][CMD]", cmd)
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			log.Println("[!!!!][Service][Daemon][Clean][Chain]", err)
		}
	}
	{
		d := path.Join(dir, "Index_00746E41")
		cmd := fmt.Sprintf("rm -rf %s", d)
		log.Println("[Service][Daemon][CMD]", cmd)
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			log.Println("[!!!!][Service][Daemon][Clean][Index]", err)
		}
	}

	exec.Command("bash", "-c", "rm -rf ")
	port := os.ExpandEnv("${NEO_PORT}")
	cmd := fmt.Sprintf("server=%s port=%s startFrom=%d %s", me.IP, port, me.StartFrom, me.ExecPath)
	me.exe = exec.Command("bash", "-c", cmd)
	log.Println("[Service][Daemon][CMD]", cmd)
}
