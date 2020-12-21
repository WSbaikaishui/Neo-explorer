package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"neophora/lib/cli"
	"net/rpc"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	defer client.Close()
	defer wg.Wait()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			close(queue)
			return
		}
		queue <- line
	}
}

var client *cli.T
var queue chan string
var wg sync.WaitGroup

func init() {
	address := os.ExpandEnv("${NEODB_ADDRESS}")
	bufsize := os.ExpandEnv("${COLLECTOR_BUFSIZE}")
	threads := os.ExpandEnv("${COLLECTOR_THREADS}")
	thread, err := strconv.Atoi(threads)
	if err != nil {
		log.Fatalln(err)
	}
	client = &cli.T{
		Address: address,
		Pool:    make(chan *rpc.Client, thread),
	}
	size, err := strconv.Atoi(bufsize)
	if err != nil {
		log.Fatalln(err)
	}

	queue = make(chan string, size)
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go worker()
	}
}

func run(f func() error) {
	for i := time.Millisecond; i < time.Hour; i = i * 2 {
		if err := f(); err != nil {
			log.Println("[????][WARN]", err)
			time.Sleep(i)
			continue
		}
		return
	}
	log.Println("[!!!!][ERROR][FAILED]")
}

func worker() {
	defer wg.Done()
	for v := range queue {
		task(v)
	}
}

func task(line string) {
	var data map[string]string
	if err := json.Unmarshal([]byte(line), &data); err != nil {
		log.Println("[!!!!][JSON]", err)
		return
	}
	args := make(map[string][]byte)
	for k, v := range data {
		val, err := hex.DecodeString(v)
		if err != nil {
			log.Println("[!!!!][HEX]", v, err)
			continue
		}
		args[k] = val
	}
	var ret bool
	for i := time.Second; i < time.Hour; i++ {
		if err := client.Call("T.Put", args, &ret); err != nil || ret == false {
			log.Println("[????][CALL]", ret, err)
			time.Sleep(i)
			continue
		}
		return
	}
	log.Println("[!!!!][LOST]")
}
