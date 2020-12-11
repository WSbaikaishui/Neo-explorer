package biz

import (
	"fmt"
	"neophora/cli"
	"neophora/var/stderr"
	"net/url"
)

// Data ...
type Data struct {
	Client *cli.T
}

// GetCountInUInt64 ...
func (me *Data) GetCountInUInt64(args struct {
}, ret *uint64) error {
	var uri url.URL

	uri.Scheme = "hash"
	uri.Host = "height"
	uri.Path = "/ffffffffffffffff"

	var result []byte

	urs := uri.String()
	if err := me.Client.Calls("DB.GetLastKey", struct {
		Key    []byte
		Prefix int
	}{
		Key:    []byte(urs),
		Prefix: len(urs) - 16,
	}, &result); err != nil {
		return stderr.ErrUnknown
	}

	if len(result) == 0 {
		return stderr.ErrNotFound
	}

	key, err := url.Parse(string(result))
	if err != nil {
		return stderr.ErrUnknown
	}

	var count uint64

	fmt.Sscanf(key.Path, "/%x", &count)

	*ret = count
	return nil
}
