package trans

import (
	"crypto/sha256"
	"encoding/hex"
	"neophora/var/stderr"
	"regexp"
	"strings"
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

// BytesReverse ...
func (me *T) BytesReverse() error {
	switch bytes := me.V.(type) {
	case []byte:
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// StringToLowerCase ...
func (me *T) StringToLowerCase() error {
	switch str := me.V.(type) {
	case string:
		me.V = strings.ToLower(str)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// Remove0xPrefix ...
func (me *T) Remove0xPrefix() error {
	switch str := me.V.(type) {
	case string:
		matches := libTransReg0x.FindStringSubmatch(str)
		if len(matches) != 3 {
			return stderr.ErrInvalidArgs
		}
		me.V = matches[2]
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

var libTransReg0x = regexp.MustCompile(`^(0x)?([0-9a-f]{64})$`)
