package rpc

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/abxuz/b-tools/bcrypt"
	rpcHttp "github.com/abxuz/b-tools/brpc/http"
)

const TimeoutDefault = time.Second * 5

var (
	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 5 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 0,
	}

	ErrInternal = errors.New("server internal error")
)

func NewClient(endpont string, clientPrivateKey string, serverPublicKey string) (*rpcHttp.Client, error) {
	var cPrivKey bcrypt.NoisePrivateKey
	if err := cPrivKey.FromString(clientPrivateKey); err != nil {
		return nil, err
	}

	var sPubKey bcrypt.NoisePublicKey
	if err := sPubKey.FromString(serverPublicKey); err != nil {
		return nil, err
	}

	c := rpcHttp.NewClient(
		rpcHttp.WithClientPrivateKey(cPrivKey),
		rpcHttp.WithServerPublicKey(sPubKey),
		rpcHttp.WithEndpoint(endpont),
		rpcHttp.WithHttpClient(httpClient),
	)
	return c, nil
}
