package api

import (
	"encoding/json"
	"neo3fura/lib/type/h256"
	"neo3fura/lib/type/uintval"
	"neo3fura/var/stderr"
)

func (me *T) Getblockheader(args []interface{}, ret *json.RawMessage) error {
	if args[1] != true {
		return stderr.ErrInvalidArgs
	}
	switch args[0].(type) {
	case string:
		blockHash := h256.T(args[0].(string))
		if blockHash.Valid() == false {
			return stderr.ErrInvalidArgs
		}
		var raw1 map[string]interface{}
		err := me.GetBlockHeaderByBlockHash(struct {
			BlockHash h256.T
			Filter    map[string]interface{}
			Raw       *map[string]interface{}
		}{
			BlockHash: blockHash,
			Filter:    nil,
			Raw:       &raw1,
		}, ret)
		if err != nil {
			return err
		}
		r, err := json.Marshal(raw1)
		if err != nil {
			return err
		}
		*ret = json.RawMessage(r)
		return nil
	case float64:
		blockHeight := uintval.T(uint64(args[0].(float64)))
		if blockHeight.Valid() == false {
			return stderr.ErrInvalidArgs
		}
		var raw1 map[string]interface{}
		err := me.GetBlockHeaderByBlockHeight(struct {
			BlockHeight uintval.T
			Filter      map[string]interface{}
			Raw         *map[string]interface{}
		}{
			BlockHeight: blockHeight,
			Filter:      nil,
			Raw:         &raw1,
		}, ret)
		if err != nil {
			return err
		}
		r, err := json.Marshal(raw1)
		if err != nil {
			return err
		}
		*ret = json.RawMessage(r)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}
