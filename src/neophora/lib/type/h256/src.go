package h256

import (
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

// RevVal ...
func (me T) RevVal() string {
	bytes, _ := hex.DecodeString(me.Val())
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return hex.EncodeToString(bytes)
}

var re = regexp.MustCompile(`^([0-9a-f]{64})$`)
