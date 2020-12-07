package biz

import (
	"encoding/hex"
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// Data ...
type Data struct {
	Client *cli.T
}

// Ping ...
func (me *Data) Ping(arg interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}

// GetDataInHex ...
func (me *Data) GetDataInHex(key string, ret *string) error {
	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(key), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetDataInString ...
func (me *Data) GetDataInString(key string, ret *string) error {
	var result []byte

	if err := me.Client.Calls("DB.Get", []byte(key), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = string(result)

	return nil
}

// GetBlockByHeightInHex ...
func (me *Data) GetBlockByHeightInHex(index uint64, ret *string) error {
	uri := &url.URL{
		Scheme: "block",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", index),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetBlockByHashInHex ...
func (me *Data) GetBlockByHashInHex(hash string, ret *string) error {
	uri := &url.URL{
		Scheme: "block",
		Host:   "hash",
		Path:   fmt.Sprintf("/%s", hash),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetHeaderByHeightInHex ...
func (me *Data) GetHeaderByHeightInHex(index uint64, ret *string) error {
	uri := &url.URL{
		Scheme: "header",
		Host:   "height",
		Path:   fmt.Sprintf("/%016x", index),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetHeaderByHashInHex ...
func (me *Data) GetHeaderByHashInHex(hash string, ret *string) error {
	uri := &url.URL{
		Scheme: "header",
		Host:   "hash",
		Path:   fmt.Sprintf("/%s", hash),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.Get", []byte(urs), &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}

// GetStorageByDBKeyHeightInHex ...
func (me *Data) GetStorageByDBKeyHeightInHex(args struct {
	DBKey  string
	Height uint64
}, ret *string) error {
	uri := &url.URL{
		Scheme: "storage",
		Host:   "dbkey-height",
		Path:   fmt.Sprintf("/%s/%016x", args.DBKey, args.Height),
	}

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLast", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	*ret = hex.EncodeToString(result)

	return nil
}
