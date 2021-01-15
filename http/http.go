package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type httpHandle struct {
	url         string
	contentType string
	timeout     time.Duration
	retryTimes  int8
	tlsVerify   bool
}

func NewHttpHandle(url string, opts ...httpOptions) *httpHandle {
	hd := &httpHandle{
		url: url,
	}

	opts = append([]httpOptions{defaultHttpOption()}, opts...)
	for _, v := range opts {
		v.apply(hd)
	}

	return hd
}

//发送Post请求
func (h *httpHandle) Post(data []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", h.url, bytes.NewReader(data))
	if err != nil {
		return nil, HttpError{
			level:       1,
			error:       err.Error(),
			callerError: "http new request error",
		}
	}

	if len(h.contentType) > 0 {
		request.Header.Set("Content-Type", h.contentType)
	}

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: h.timeout,
		}).DialContext,
	}

	//跳过证书验证
	if !h.tlsVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	if err != nil {
		return nil, HttpError{
			level:       2,
			error:       err.Error(),
			callerError: "http client do error",
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, HttpError{
			level:       3,
			statusCode:  resp.StatusCode,
			error:       resp.Status,
			callerError: fmt.Sprint("http status error [%d]", resp.StatusCode),
		}
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, HttpError{
			level:       4,
			statusCode:  resp.StatusCode,
			error:       err.Error(),
			callerError: fmt.Sprint("http read body error"),
		}
	}
	return respBytes, nil
}

//发送Get请求
func (h *httpHandle) Get() ([]byte, error) {
	client := &http.Client{Timeout: h.timeout}
	resp, err := client.Get(h.url)
	if err != nil {
		return nil, HttpError{
			level:       2,
			error:       err.Error(),
			callerError: "http client do error",
		}
	}
	defer resp.Body.Close()

	var (
		buffer [1024]byte
		n      int
	)
	result := bytes.NewBuffer(nil)
	for {
		n, err = resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, HttpError{
				level:       4,
				statusCode:  resp.StatusCode,
				error:       err.Error(),
				callerError: fmt.Sprint("http read body error"),
			}
		}
	}

	return result.Bytes(), nil
}
