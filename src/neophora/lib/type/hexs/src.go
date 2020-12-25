package hexs

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

// T ...
type T string

// Valid ...
func (me T) Valid() bool {
	return re.MatchString(me.Val())
}

// Val ...
func (me T) Val() string {
	return string(me)
}

// Decode ...
func (me T) Decode() []byte {
	data, _ := hex.DecodeString(me.Val())
	return data
}

// H256 ...
func (me T) H256() string {
	data := me.Decode()
	l1 := sha256.Sum256(data)
	l2 := sha256.Sum256(l1[:])
	return hex.EncodeToString(l2[:])
}

var re = regexp.MustCompile(`^[0-9a-f]*$`)
