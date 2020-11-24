package biz

// Version ...
type Version struct{}

// Ping ...
func (me *Version) Ping(arg interface{}, ret *interface{}) error {
	*ret = "pong"
	return nil
}
