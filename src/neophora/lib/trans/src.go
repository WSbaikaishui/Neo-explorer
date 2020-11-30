package trans

import (
	"crypto/sha256"
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

// HexToBytes ...
func (me *T) HexToBytes() error {
	switch enc := me.V.(type) {
	case string:
		var err error
		me.V, err = hex.DecodeString(enc)
		return err
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToHash ...
func (me *T) BytesToHash() error {
	switch bytes := me.V.(type) {
	case []byte:
		l1 := sha256.Sum256(bytes)
		l2 := sha256.Sum256(l1[:])
		me.V = l2[:]
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}
