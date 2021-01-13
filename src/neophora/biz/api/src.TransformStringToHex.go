package api

import (
	"neophora/lib/type/strval"
	"neophora/var/stderr"
)

// TransformStringToHex ...
// as an example:
//
// ```
// $ curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "TransformStringToHex","params":{"String":"I am a boy"}}'
// {"id":1,"result":"4920616d206120626f79","error":null}
// ```
func (me *T) TransformStringToHex(args struct {
	String strval.T
}, ret *string) error {
	if args.String.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	*ret = args.String.Hex()
	return nil
}
