package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/hexs"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEHexKeyBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEHexKeyBlockHeightInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","HexKey":"736b","BlockHeight":99999999}}'
// {"id":1,"result":"0007324861c968b43100","error":null}
//
// ```
func (me *T) GetStorageByContractHashLEHexKeyBlockHeightInHex(args struct {
	ContractHashLE h160.T
	HexKey         hexs.T
	BlockHeight    uintval.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.HexKey.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.HexKey.H256(), args.BlockHeight.Hex()},
	}, ret)
}
