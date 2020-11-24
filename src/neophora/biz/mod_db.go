package biz

import (
	"encoding/hex"
	"neophora/dat"
	"neophora/var/stderr"
	"reflect"
)

// DB ...
type DB struct {
	DB *dat.T
}

// PutInStringValue ...
func (me *DB) PutInStringValue(arg map[string]string, ret *interface{}) error {
	for k, v := range arg {
		err := me.DB.Put([]byte(k), []byte(v))
		if err != nil {
			return err
		}
	}
	*ret = "ok"
	return nil
}

// GetInStringValue ...
func (me *DB) GetInStringValue(arg []string, ret *interface{}) error {
	result := make(map[string]interface{})
	for _, k := range arg {
		value, err := me.DB.Get([]byte(k))
		if err != nil {
			return err
		}
		result[k] = string(value)
	}
	*ret = result
	return nil
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

// Get ...
func (me *DB) Get(data []byte, ret *[]byte) error {
	value, err := me.DB.Get(data)
	if err != nil {
		return err
	}
	*ret = value
	return nil
}

// Batch ...
func (me *DB) Batch(data struct {
	Method string
	Args   []interface{}
}, ret *[]interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = stderr.ErrUnknown
		}
	}()
	cls := reflect.TypeOf(me)
	method, ok := cls.MethodByName(data.Method)
	if ok == false {
		return stderr.ErrUnsupportedMethod
	}
	result := make([]interface{}, len(data.Args))
	for i, v := range data.Args {
		typ := method.Type.In(2).Elem()
		val := reflect.New(typ)
		ret := method.Func.Call([]reflect.Value{
			reflect.ValueOf(me),
			reflect.ValueOf(v),
			val,
		})
		if len(ret) != 1 {
			return stderr.ErrUnknown
		}
		if ret[0].IsNil() == false {
			return ret[0].Interface().(error)
		}
		result[i] = val.Elem().Interface()
	}
	*ret = result
	return nil
}

// Ping ...
func (me *DB) Ping(arg []interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}
