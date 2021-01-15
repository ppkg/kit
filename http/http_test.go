package http

import (
	"testing"
)

func Test_httpHandle_Get(t *testing.T) {
	tests := []struct {
		name    string
		handle  *httpHandle
		want    []byte
		wantErr bool
	}{
		{
			name:   "get请求测试",
			handle: NewHttpHandle("http://www.baidu.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.handle.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(string(got))
		})
	}
}
