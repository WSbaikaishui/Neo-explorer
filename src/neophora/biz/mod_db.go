package biz

import (
	"encoding/hex"
	"neophora/dat"
)

// DB ...
type DB struct {
	DB *dat.T
}

// PutInHexValue ...
func (me *DB) PutInHexValue(arg map[string]string, ret *interface{}) error {
	for k, v := range arg {
		key := []byte(k)
		value, err := hex.DecodeString(v)
		if err != nil {
			return err
		}
		err = me.DB.Put(key, value)
		if err != nil {
			return err
		}
	}
	*ret = "ok"
	return nil
}

// GetInHexValue ...
func (me *DB) GetInHexValue(arg []string, ret *interface{}) error {
	result := make(map[string]interface{})
	for _, k := range arg {
		value, err := me.DB.Get([]byte(k))
		if err != nil {
			return err
		}
		result[k] = hex.EncodeToString(value)
	}
	*ret = result
	return nil
}

// Ping ...
func (me *DB) Ping(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}
