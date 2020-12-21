package main

import (
	"log"
	"neophora/biz/database"
	"net"
	"net/rpc"
	"os"

	"github.com/tecbot/gorocksdb"
)

func main() {
	dbpath := os.ExpandEnv("${DATABASE_PATH}")
	address := os.ExpandEnv("0.0.0.0:${NEODB_PORT}")

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

	ro := gorocksdb.NewDefaultReadOptions()
	defer wo.Destroy()

	rpc.Register(&database.T{
		DB: db,
		WO: wo,
		RO: ro,
	})

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("[LISTEN]", err)
	}

	log.Println("[LISTEN]", address)
	rpc.DefaultServer.Accept(listener)
}
