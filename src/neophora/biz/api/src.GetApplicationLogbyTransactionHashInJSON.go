package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetApplicationLogByTransactionHashBlockHeightInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetApplicationLogByTransactionHashInJSON","params":{"TransactionHash":"769bdbc97a875e0403668fb6feabf007d6728b9086d9a65344c11d6023eb870b"}}'
// {"id":1,"result":{"executions":[{"contract":"0x64783fd1fd437115d3287886d887735ab2f917e4","gas_consumed":"2.104","notifications":[{"contract":"0xc3361fef35233f9db6354b259be9cde34ba667c5","state":{"type":"Array","value":[{"type":"ByteArray","value":"7472616e73666572"},{"type":"ByteArray","value":"6c4ccbce0b0a8f0429fae4588ea60e21dbf327da"},{"type":"ByteArray","value":"d34bea918239d8b405c96c096443ff72eeb8761e"},{"type":"ByteArray","value":"004a2e1815"}]}}],"stack":[],"trigger":"Application","vmstate":"HALT"}],"txid":"0x0b87eb23601dc14453a6d986908b72d607f0abfeb68f6603045e877ac9db9b76"},"error":null}
// ```

func (me *T) GetApplicationLogByTransactionHashInJSON(args struct {
	TransactionHash h256.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLastValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.apl",
		Index:  "h256.trx",
		Keys:   []string{args.TransactionHash.Val()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaApplicationLog()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}