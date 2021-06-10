package api

import (
	"neo3fura/biz/data"
)

// T ...
type T struct {
	Data *data.T
}

// Ping ...
func (me *T) Ping(args struct{}, ret *string) error {
	*ret = "pong"
	return nil
}

func (me *T) Filter(data map[string]interface{}, filter map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	if filter == nil {
		return data, nil
	}
	if len(filter) == 0 {
		return data, nil
	}
	for k, _ := range filter {
		if data[k] != nil {
			switch data[k].(type) {
			case string:
				res[k] = data[k]
			case int:
				res[k] = data[k]
			case interface{}:
				r, err := me.Filter(data[k].(map[string]interface{}), filter[k].(map[string]interface{}))
				if err != nil {
					return nil, err
				}
				res[k] = r
			}
		}
	}
	return res, nil
}

func (me *T) FilterArray(data []map[string]interface{}, filter map[string]interface{}) ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0)
	for _, item := range data {
		r, err := me.Filter(item, filter)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
