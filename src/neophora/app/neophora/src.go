package main

import (
	"context"
	"fmt"
	"gopkg.in/neophora/biz/api"
	"gopkg.in/neophora/biz/data"
	"gopkg.in/neophora/lib/cli"
	"gopkg.in/neophora/lib/joh"
	"log"
	"net/http"
	"net/rpc"
	"os"
	// "strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)



func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	fmt.Println("connected")
	defer cancel()
	//address := os.ExpandEnv("${NEODB_ADDRESS}")
	//poolsize, err := strconv.Atoi(os.ExpandEnv("${NEODB_POOLSIZE}"))
	if err != nil {
		log.Fatalln(err)
	}
	client := &cli.T{
		C:   c,
		Ctx: ctx,
	}

	defer cancel()

	rpc.Register(&api.T{
		Data: &data.T{
			Client: client,
		},
	})

	listen := os.ExpandEnv("0.0.0.0:${NEOPHORA_PORT}")
	log.Println("[LISTEN]", listen)
	http.ListenAndServe(listen, &joh.T{})
}
