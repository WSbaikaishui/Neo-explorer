package strval

import (
	"crypto/sha256"
	"encoding/hex"
)

// T ...
type T string

// Valid ...
func (me T) Valid() bool {
	return true
}

// Val ...
func (me T) Val() string {
	return string(me)
}

// Bytes ...
func (me T) Bytes() []byte {
	return []byte(me.Val())
}

// Hex ...
func (me T) Hex() string {
	return hex.EncodeToString(me.Bytes())
}

// H256 ...
func (me T) H256() string {
	data := me.Bytes()
	l1 := sha256.Sum256(data)
	l2 := sha256.Sum256(l1[:])
	return hex.EncodeToString(l2[:])
}
