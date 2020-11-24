package biz

import (
	"neophora/dat"
)

// DB ...
type DB struct {
	DB *dat.T
}

// Put ...
func (me *DB) Put(data struct {
	Key   []byte
	Value []byte
}, ret *bool) error {
	if err := me.DB.Put(data.Key, data.Value); err != nil {
		return err
	}
	*ret = true
	return nil
}

// Puts ...
func (me *DB) Puts(data []struct {
	Key   []byte
	Value []byte
}, ret *bool) error {
	for _, v := range data {
		if err := me.DB.Put(v.Key, v.Value); err != nil {
			return err
		}
	}
	*ret = true
	return nil
}

// Get ...
func (me *DB) Get(data []byte, ret *[]byte) error {
	value, err := me.DB.Get(data)
	if err != nil {
		return err
	}
	*ret = value
	return nil
}

// Gets ...
func (me *DB) Gets(data [][]byte, ret *[][]byte) error {
	result := make([][]byte, len(data))
	for i, v := range data {
		value, err := me.DB.Get(v)
		if err != nil {
			return err
		}
		result[i] = value
	}
	*ret = result
	return nil
}

// Ping ...
func (me *DB) Ping(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}
