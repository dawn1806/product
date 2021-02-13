package common

import "encoding/json"

// 通过 json tag 进行结构体赋值
func SwapTo(req, category interface{}) error {
	bts, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, category)
}
