package trans

import (
	"encoding/hex"
	"neophora/var/stderr"
)

// T ...
type T struct {
	V interface{}
}

// AddressToHash ...
func (me *T) AddressToHash() error {
	switch address := me.V.(type) {
	case string:
		_ = address
		// TODO
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToHex ...
func (me *T) BytesToHex() error {
	switch bytes := me.V.(type) {
	case []byte:
		me.V = hex.EncodeToString(bytes)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}
