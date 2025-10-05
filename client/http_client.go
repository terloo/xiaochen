package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	neturl "net/url"
	"time"

	"github.com/pkg/errors"
)

func HttpGet(ctx context.Context, url string, header http.Header, param neturl.Values) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	_url.RawQuery = param.Encode()

	timeoutCtx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()
	request, err := http.NewRequestWithContext(timeoutCtx, http.MethodGet, _url.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "new http request error")
	}
	if header == nil {
		header = http.Header{}
	}
	request.Header = header
	request.Header.Add("Accept", "*/*")
	request.Header.Add("User-Agent", "xiaochen/0.0.1")

	log.Printf("http GET request, url: %s, header: %+v\n", _url.String(), header)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response respBody error")
	}
	if resp.ContentLength != -1 && resp.ContentLength != int64(len(respBody)) {
		return nil, errors.Errorf("read response respBody error, except length [%d], but got length [%d]", resp.ContentLength, len(respBody))
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("http status code not ok: %d, respBody: %s", resp.StatusCode, string(respBody))
	}
	log.Printf("http GET request, url: %s, resp: %s\n", _url.String(), string(respBody))
	return respBody, nil
}

func HttpPost(ctx context.Context, url string, header http.Header, param neturl.Values, body interface{}) ([]byte, error) {
	_url, err := neturl.Parse(url)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	_url.RawQuery = param.Encode()

	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrapf(err, "request body json marshal error, body: %+v", body)
	}

	timeoutCtx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	request, err := http.NewRequestWithContext(timeoutCtx, http.MethodPost, _url.String(), bytes.NewReader(b))
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

	log.Printf("http POST request, url: %s, header: %+v, body: %+v\n", _url.String(), header, string(b))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body error")
	}
	if resp.ContentLength != -1 && resp.ContentLength != int64(len(respBody)) {
		return nil, errors.Errorf("read response body error, except length [%d], but got length [%d]", resp.ContentLength, len(respBody))
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("http status code not ok: %d, body: %s", resp.StatusCode, string(respBody))
	}
	log.Printf("http POST request, url: %s, req: %+v, resp: %s\n", _url.String(), body, string(respBody))
	return respBody, nil
}
