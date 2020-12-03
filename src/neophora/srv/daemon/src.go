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
	ExecDir   string
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
	err := me.exe.Run()
	if err != nil {
		log.Println("[Service][Daemon]", err)
	}
}

// Reset ...
func (me *T) Reset() {
	me.lock.Lock()
	defer me.lock.Unlock()
	{
		d := path.Join(me.ExecDir, "Chain_00746E41")
		cmd := fmt.Sprintf("rm -rf %s", d)
		log.Println("[Service][Daemon][CMD]", cmd)
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			log.Println("[!!!!][Service][Daemon][Clean][Chain]", err)
		}
	}
	{
		d := path.Join(me.ExecDir, "Index_00746E41")
		cmd := fmt.Sprintf("rm -rf %s", d)
		log.Println("[Service][Daemon][CMD]", cmd)
		if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
			log.Println("[!!!!][Service][Daemon][Clean][Index]", err)
		}
	}

	port := os.ExpandEnv("${NEO_PORT}")
	cmd := fmt.Sprintf("cd %s && sleep 5 && SERVER=%s PORT=%s STARTFROM=%d %s", me.ExecDir, me.IP, port, me.StartFrom, me.ExecCmd)
	me.exe = exec.Command("bash", "-c", cmd)
	log.Println("[Service][Daemon][CMD]", cmd)
}
