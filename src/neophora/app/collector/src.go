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
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	netowrk := os.ExpandEnv("${REDIS_NETWORK}")
	address := os.ExpandEnv("${REDIS_ADDRESS}")
	reader := bufio.NewReader(os.Stdin)
	db := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial(netowrk, address)
	}, 2)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if len(line) == 0 {
			continue
		}
		var data struct {
			Key string `json:"key"`
			Val string `json:"val"`
		}
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			log.Println("[????][Parse]", err, string(line))
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

		for {
			if _, err := db.Get().Do("SET", uri.String(), value); err != nil {
				log.Println("[!!!!][REQ]", err)
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}

		switch uri.Scheme {
		case "hash":
			switch uri.Host {
			case "height":
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
				for {
					if _, err := db.Get().Do("ZADD", key.String(), score, uri.String()); err != nil {
						log.Println("[!!!!][REQ]", err)
						time.Sleep(1 * time.Second)
						continue
					}
					break
				}
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "block", "blockhashdata", "header", "tx":
			switch uri.Host {
			case "height", "hash":
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "adhocblockinfo", "adhocheaderinfo", "adhocsysfee", "adhoctxinfo", "adhocstateroot":
			switch uri.Host {
			case "height", "hash":
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "height":
			switch uri.Host {
			case "tx":
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "adhoclog":
			switch uri.Host {
			case "tx":
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "contract", "storage":
			switch uri.Host {
			case "hash-height":
				var score uint64
				if n, err := fmt.Sscanf(path.Base(uri.Path), "%016x", &score); err != nil || n != 1 {
					log.Println("[!!!!][INDEX]", uri)
					time.Sleep(1 * time.Second)
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
				for {
					if _, err := db.Get().Do("ZADD", key.String(), score, uri.String()); err != nil {
						log.Println("[!!!!][REQ]", err)
						continue
					}
					break
				}
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		case "adhocaccountstate", "adhocassetstate", "adhoccontractstate", "adhocclaimable", "adhocunspents", "adhocnep5balances":
			switch uri.Host {
			case "account-height", "hash-height":
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
				for {
					if _, err := db.Get().Do("ZADD", key.String(), score, uri.String()); err != nil {
						log.Println("[!!!!][REQ]", err)
						time.Sleep(1 * time.Second)
						continue
					}
					break
				}
			default:
				log.Println("[!!!!][KEY]", uri)
			}
		default:
			log.Println("[!!!!][KEY]", uri)
		}
	}
}
