package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	neturl "net/url"

	"github.com/pkg/errors"
)

func HttpGet(ctx context.Context, url string, header http.Header, param neturl.Values) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	urlValues := neturl.Values{}
	for k, v := range param {
		urlValues.Add(k, v[0])
	}
	_url.RawQuery = urlValues.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, _url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "new http request error")
	}
	if header == nil {
		header = http.Header{}
	}
	request.Header = header
	request.Header.Add("Accept", "*/*")
	request.Header.Add("User-Agent", "xiaochen/0.0.1")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body error")
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("http status code not ok: %d, body: %s", resp.StatusCode, string(body))
	}
	return body, nil
}

func HttpPost(ctx context.Context, url string, header http.Header, param map[string]string, body interface{}) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	urlValues := neturl.Values{}
	for k, v := range param {
		urlValues.Add(k, v)
	}
	_url.RawQuery = urlValues.Encode()

	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrapf(err, "request body json marshal error, body: %+v", body)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, _url.String(), bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "new http request error")
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
		return nil, errors.Wrap(err, "http request error")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body error")
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("http status code not ok: %d, body: %s", resp.StatusCode, string(respBody))
	}
	return respBody, nil
}
