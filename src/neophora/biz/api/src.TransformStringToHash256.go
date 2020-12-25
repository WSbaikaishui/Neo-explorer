package api

import (
	"neophora/lib/type/strval"
	"neophora/var/stderr"
)

// TransformStringToHash256 ...
// as an example:
//
// ```
// TODO
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
