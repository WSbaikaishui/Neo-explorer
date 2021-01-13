package api

import (
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// TransformHexToReversedHex ...
// as an example:
//
// ```
// $ curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "TransformHexToReversedHex","params":{"Hex":"abcd"}}'
// {"id":1,"result":"cdab","error":null}
// ```
func (me *T) TransformHexToReversedHex(args struct {
	Hex hexs.T
}, ret *string) error {
	if args.Hex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	*ret = args.Hex.RevVal()
	return nil
}
