package kit

import "github.com/ppkg/kit/cast"

//查询值是否在切片存在
func InSlice(value interface{}, list interface{}) bool {
	slice, err := cast.ToSliceE(list)
	if err != nil {
		panic(err)
	}

	for k := range slice {
		if slice[k] == value {
			return true
		}
	}
	return false
}
