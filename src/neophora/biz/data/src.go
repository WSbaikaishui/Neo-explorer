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

// Get ...
func (me *T) Get(args struct {
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

// GetArgs ...
func (me *T) GetArgs(args struct {
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

// GetLastKeyEx ...
func (me *T) GetLastKeyEx(args struct {
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

// GetLastEx ...
func (me *T) GetLastEx(args struct {
	Target string
	Index  string
	Keys   []string
	Max    uint64
	Min    uint64
}, ret *[]byte) error {
	var result []byte
	if err := me.GetLastKeyEx(args, &result); err != nil {
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

// GetLastKey ...
func (me *T) GetLastKey(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	return me.GetLastKeyEx(struct {
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

// GetLast ...
func (me *T) GetLast(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *[]byte) error {
	return me.GetLastEx(struct {
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

// GetHex ...
func (me *T) GetHex(args struct {
	Key string
}, ret *string) error {
	var result []byte
	if err := me.Get(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetArgsHex ...
func (me *T) GetArgsHex(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetArgs(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetLastExHex ...
func (me *T) GetLastExHex(args struct {
	Target string
	Index  string
	Keys   []string
	Max    uint64
	Min    uint64
}, ret *string) error {
	var result []byte
	if err := me.GetLastEx(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}

// GetLastHex ...
func (me *T) GetLastHex(args struct {
	Target string
	Index  string
	Keys   []string
}, ret *string) error {
	var result []byte
	if err := me.GetLast(args, &result); err != nil {
		return err
	}
	*ret = hex.EncodeToString(result)
	return nil
}
