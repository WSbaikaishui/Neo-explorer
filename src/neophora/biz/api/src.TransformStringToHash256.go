package api

import (
	"neophora/lib/type/strval"
	"neophora/var/stderr"
)

// TransformStringToHash256 ...
// as an example:
//
// ```
// curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "TransformStringToHash256","params":{"String":"I am a boy"}}'
// {"id":1,"result":"b22dade0e36dc2978dd290d46b4a10797482874ff0350c31535910893e449bca","error":null}
// ```
func (me *T) TransformStringToHash256(args struct {
	String strval.T
}, ret *string) error {
	if args.String.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	*ret = args.String.H256()
	return nil
}
