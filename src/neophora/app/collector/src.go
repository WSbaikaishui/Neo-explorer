package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		var data struct {
			Key string `json:"key"`
			Val string `json:"val"`
		}
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			continue
		}
		uri, err := url.Parse(data.Key)
		if err != nil {
			log.Println("[!!!!][URI]", data.Key)
			continue
		}
		value, err := hex.DecodeString(data.Val)
		if err != nil {
			log.Println("[!!!!][HEX]", data.Val)
			continue
		}

		log.Println("[DATA]", uri)

		run(func() error {
			_, err := db.Get().Do("SET", uri.String(), value)
			return err
		})

		switch fmt.Sprintf("%s://%s", uri.Scheme, uri.Host) {
		case "hash://height",
			"contract://hash-height",
			"storage://hash-height",
			"adhocassetstate://hash-height",
			"adhoccontractstate://hash-height",
			"adhocclaimable://account-height",
			"adhocunspents://account-height",
			"adhocnep5balances://account-height",
			"adhocaccountstate://account-height":
			var score uint64
			if n, err := fmt.Sscanf(path.Base(uri.Path), "%016x", &score); err != nil || n != 1 {
				log.Println("[!!!!][INDEX]", uri)
				continue
			}
			urc := &url.URL{}
			*urc = *uri
			urc.Path = path.Join(path.Dir(urc.Path), "_")
			key := &url.URL{
				Scheme: "index",
				Host:   "keys",
				Path:   urc.String(),
			}
			run(func() error {
				_, err := db.Get().Do("ZADD", key.String(), score, uri.String())
				return err
			})
		}
	}
}

var db *redis.Pool

func init() {
	netowrk := os.ExpandEnv("${REDIS_NETWORK}")
	address := os.ExpandEnv("${REDIS_ADDRESS}")
	maxidle, err := strconv.Atoi(os.ExpandEnv("${REDIS_MAXIDLE}"))
	if err != nil {
		log.Fatalln(err)
	}
	db = redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(netowrk, address)
	}, maxidle)
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
