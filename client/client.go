package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	neturl "net/url"
)

func HttpGet(ctx context.Context, url string, header http.Header, param neturl.Values) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}
	urlValues := neturl.Values{}
	for k, v := range param {
		urlValues.Add(k, v[0])
	}
	_url.RawQuery = urlValues.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, _url.String(), nil)
	if err != nil {
		return nil, err
	}
	if header == nil {
		header = http.Header{}
	}
	request.Header = header
	request.Header.Add("Accept", "*/*")
	request.Header.Add("User-Agent", "xiaochen/0.0.1")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func HttpPost(ctx context.Context, url string, header http.Header, param map[string]string, body interface{}) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}
	urlValues := neturl.Values{}
	for k, v := range param {
		urlValues.Add(k, v)
	}
	_url.RawQuery = urlValues.Encode()

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, _url.String(), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if header == nil {
		header = http.Header{}
	}
	request.Header = header
	request.Header.Add("Accept", "*/*")
	request.Header.Add("User-Agent", "xiaochen/0.0.1")
	request.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
