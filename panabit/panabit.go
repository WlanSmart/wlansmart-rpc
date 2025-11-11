package panabit

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	rpc "github.com/WlanSmart/wlansmart-rpc/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	TimeoutDefault = time.Second * 5
)

type ApiResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

type PanabitAdapter struct {
	endpoint string
	key      string
	cli      *http.Client
}

func NewPanabitAdapter() *PanabitAdapter {
	return &PanabitAdapter{
		cli: rpc.DefaultHttpClient,
	}
}

func (a *PanabitAdapter) SetHttpClient(cli *http.Client) *PanabitAdapter {
	a.cli = cli
	return a
}

func (a *PanabitAdapter) SetRawEndpoint(endpoint string) *PanabitAdapter {
	a.endpoint = endpoint
	return a
}

func (a *PanabitAdapter) SetEndpoint(endpoint string) *PanabitAdapter {
	return a.SetRawEndpoint(endpoint + "/api/wlansmart.cgi")
}

func (a *PanabitAdapter) SetAddress(addr string) *PanabitAdapter {
	return a.SetEndpoint("https://" + addr)
}

func (a *PanabitAdapter) SetKey(key string) *PanabitAdapter {
	a.key = key
	return a
}

func (a *PanabitAdapter) RawRequest(ctx context.Context, dir, name, action string, params url.Values) ([]byte, error) {
	t := strconv.FormatInt(time.Now().Unix(), 10)
	token := md5.Sum([]byte(t + "," + a.key))
	url := a.endpoint +
		"?api_time=" + t +
		"&api_token=" + hex.EncodeToString(token[:]) +
		"&api_route=" + dir + "@" + name +
		"&api_action=" + action
	if len(params) > 0 {
		url += "&" + params.Encode()
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := a.cli.Do(request)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	contentType := response.Header.Get("Content-Type")
	if strings.Contains(contentType, "gb2312") {
		return io.ReadAll(transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder()))
	}
	return io.ReadAll(response.Body)
}

func (a *PanabitAdapter) Request(ctx context.Context, dir, name, action string, params url.Values, resp any) error {
	data, err := a.RawRequest(ctx, dir, name, action, params)
	if err != nil {
		return err
	}

	r := new(ApiResponse)
	err = json.Unmarshal(data, r)
	if err != nil {
		return errors.New(string(data))
	}
	if r.Code != 0 {
		if r.Msg != "" {
			return errors.New(r.Msg)
		}
		return errors.New(string(data))
	}
	if r.Data != nil && resp != nil {
		return json.Unmarshal(r.Data, resp)
	}
	return nil
}

func (a *PanabitAdapter) RunCmd(ctx context.Context, cmd string) (string, error) {
	params := make(url.Values)
	params.Set("cmd", cmd)
	data, err := a.RawRequest(context.Background(), "App/wlansmart_api", "wlansmart_api", "runcmd", params)
	return string(data), err
}
