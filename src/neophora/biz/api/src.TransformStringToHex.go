package api

import (
	"neophora/lib/type/strval"
	"neophora/var/stderr"
)

// TransformStringToHex ...
// as an example:
//
// ```
// TODO
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
