package data

import (
	"encoding/hex"
	"neophora/var/stderr"
	"net/url"
	"path"

	"github.com/gomodule/redigo/redis"
)

// T ...
type T struct {
	DB *redis.Pool
}

// GET ...
func (me *T) GET(args struct {
	Key string
}, ret *[]byte) error {
	reply, err := me.DB.Get().Do("GET", args.Key)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	*ret = result
	return nil
}

// GETARGS ...
func (me *T) GETARGS(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	uri := &url.URL{
		Scheme: args.Target,
		Host:   args.Index,
		Path:   path.Join(args.Keys...),
	}
	reply, err := me.DB.Get().Do("GET", uri.String())
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	result, ok := reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	*ret = result
	return nil
}

// GETLASTKEYEX ...
func (me *T) GETLASTKEYEX(args struct {
	Target string
	Index  string
	Keys   []string
	Max    uint64
	Min    uint64
}, ret *[]byte) error {
	uri := &url.URL{
		Scheme: args.Target,
		Host:   args.Index,
		Path:   path.Join(args.Keys...),
	}
	uri.Path = uri.String()
	uri.Host = "keys"
	uri.Scheme = "index"

	reply, err := me.DB.Get().Do("ZREVRANGEBYSCORE", uri.String(), args.Max, args.Min, "LIMIT", 0, 1)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	replies, ok := reply.([]interface{})
	if ok == false {
		return stderr.ErrUnknown
	}

	if len(replies) != 1 {
		return stderr.ErrNotFound
	}

	result, ok := replies[0].([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	*ret = result
	return nil
}

// GETLASTEX ...
func (me *T) GETLASTEX(args struct {
	Target string
	Index  string
	Keys   []string
	Max    uint64
	Min    uint64
}, ret *[]byte) error {
	var result []byte
	if err := me.GETLASTKEYEX(args, &result); err != nil {
		return err
	}

	if len(result) == 0 {
		return stderr.ErrUnknown
	}

	reply, err := me.DB.Get().Do("GET", result)
	if err != nil {
		return stderr.ErrUnknown
	}

	if reply == nil {
		return stderr.ErrNotFound
	}

	var ok bool
	result, ok = reply.([]byte)
	if ok == false {
		return stderr.ErrUnknown
	}

	*ret = result
	return nil
}

// GETLASTKEY ...
func (me *T) GETLASTKEY(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	return me.GETLASTKEYEX(struct {
		Target string
		Index  string
		Keys   []string
		Max    uint64
		Min    uint64
	}{
		Target: args.Target,
		Index:  args.Index,
		Keys:   args.Keys,
		Max:    0xffffffff,
		Min:    0,
	}, ret)
}

// GETLAST ...
func (me *T) GETLAST(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	return me.GETLASTEX(struct {
		Target string
		Index  string
		Keys   []string
		Max    uint64
		Min    uint64
	}{
		Target: args.Target,
		Index:  args.Index,
		Keys:   args.Keys,
		Max:    0xffffffff,
		Min:    0,
	}, ret)
}

// GETHEX ...
func (me *T) GETHEX(args struct {
	Key string
}, ret *string) error {
	var result []byte
	if err := me.GET(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GETARGSHEX ...
func (me *T) GETARGSHEX(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GETARGS(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GETLASTEXHEX ...
func (me *T) GETLASTEXHEX(args struct {
	Target string
	Index  string
	Keys   []string
	Max    uint64
	Min    uint64
}, ret *string) error {
	var result []byte
	if err := me.GETLASTEX(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GETLASTHEX ...
func (me *T) GETLASTHEX(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GETLAST(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}
