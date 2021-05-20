package joh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"neophora/lib/rwio"
	"neophora/lib/scex"
	"net/http"
	"net/http/httputil"
	"net/rpc"
	"net/url"
	"path/filepath"
	"sort"

	// "sort"
)

// T ...
type T struct {
}

type Config struct {
	Methods struct {
		Realized []string `yaml:"realized"`
	} `yaml:"methods"`
}

func (me *T) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err!=nil {
		log.Printf("Error in reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
	}
	r := req.Clone(req.Context())
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	request := make(map[string]interface{})
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error decoding in JOSN: %v", err)
		http.Error(w, "can't decoding in JSON", http.StatusBadRequest)
	}
	c, err := me.OpenConfigFile()
	if err != nil {
		log.Fatalln(err)
	}
	sort.Strings(c.Methods.Realized)
	index := sort.SearchStrings(c.Methods.Realized, fmt.Sprintf("%v", request["method"]))
	if index < len(c.Methods.Realized) && c.Methods.Realized[index] == request["method"] {
		conn := &rwio.T{R: req.Body, W: w}
		codec := &scex.T{}
		codec.Init(conn)
		rpc.ServeCodec(codec)
	} else {
		ForwardHandler(w, r)
	}
}

func ForwardHandler(writer http.ResponseWriter, request *http.Request) {
	u, err := url.Parse("https://seed1t.neo.org:20332")
	if nil != err {
		log.Println(err)
		return
	}
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
		},
	}
	proxy.ServeHTTP(writer, request)
}

func (me *T) OpenConfigFile() (Config, error) {
	absPath, _ := filepath.Abs("./config.yml")
	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
