package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/tecbot/gorocksdb"
)

func main() {
	dbpath := os.ExpandEnv("${DATABASE_PATH}")

	pe := gorocksdb.NewFixedPrefixTransform(8)
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	opts.SetPrefixExtractor(pe)
	defer opts.Destroy()

	db, err := gorocksdb.OpenDb(opts, dbpath)
	if err != nil {
		log.Fatalln("[OPENDB]", err)
	}
	defer db.Close()

	wo := gorocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()
	reader := bufio.NewReader(os.Stdin)
	batch := gorocksdb.NewWriteBatch()
	for i := 0; true; i++ {
		log.Println("[HEIGHT]", i)
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}

		var data map[string]string

		if err := json.Unmarshal([]byte(line), &data); err != nil {
			log.Println("[!!!!][JSON]", err)
			continue
		}

		batch.Clear()
		for k, v := range data {
			key := []byte(k)
			val, err := hex.DecodeString(v)
			if err != nil {
				log.Println("[!!!!][HEX]", err)
				continue
			}
			batch.Put(key, val)
		}
		if err := db.Write(wo, batch); err != nil {
			log.Fatalln(err)
		}
	}
}
