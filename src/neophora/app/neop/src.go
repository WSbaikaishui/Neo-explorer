package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"neophora/biz/api"
	"net/http"
	"os"
	"reflect"
)

func main() {
	endpoint := os.ExpandEnv("${ENDPOINT}")
	if len(endpoint) == 0 {
		endpoint = "https://mainnet.neophora.io:4443"
	}
	typ := reflect.TypeOf(&api.T{})
	if len(os.Args) < 2 {
		usage(typ)
		return
	}
	method := os.Args[1]
	fun, ok := typ.MethodByName(method)
	if ok == false {
		usage(typ)
		return
	}
	arg := fun.Type.In(1)
	set := flag.NewFlagSet(method, flag.ExitOnError)
	params := make(map[string]interface{})
	for i := 0; i < arg.NumField(); i++ {
		switch arg.Field(i).Type.Kind() {
		case reflect.String:
			params[arg.Field(i).Name] = set.String(arg.Field(i).Name, "", "")
		case reflect.Uint64:
			params[arg.Field(i).Name] = set.Uint64(arg.Field(i).Name, 0, "")
		}
	}
	set.Parse(os.Args[2:])
	obj := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}
	req, err := json.Marshal(obj)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(req))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func usage(typ reflect.Type) {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "%s <COMMAND> [-k=v]...\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "To See Help Message of Each Command, Try:\n")
	fmt.Fprintf(os.Stderr, "%s <COMMAND> -help\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Availiable Commmands are:\n")
	for i := 0; i < typ.NumMethod(); i++ {
		fun := typ.Method(i)
		fmt.Fprintf(os.Stderr, "\t%s\n", fun.Name)
	}
}
