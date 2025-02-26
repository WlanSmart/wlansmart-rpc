package rpc

import (
	"errors"

	"github.com/abxuz/b-tools/brpc/http"
	"golang.org/x/sync/singleflight"
)

var (
	ErrModemNotFound = errors.New("modem not found")
)

type Gateway struct {
	Client *http.Client
	sg     singleflight.Group
}

type (
	GatewayPullConfigRequest struct {
		SN       string
		Firmware string
	}

	GatewayPullConfigResponse struct {
		Radio1      RadioConfig
		Radio2      RadioConfig
		WiredEnable bool
		WANEnable   bool
		BridgeBit   byte
		Command     string
	}

	RadioConfig struct {
		SSID    string
		Key     string
		Channel string
		Enable  bool
	}
)

func (g *Gateway) PullConfigSingleflight(req *GatewayPullConfigRequest) (resp *GatewayPullConfigResponse, err error) {
	v, err, _ := g.sg.Do("PullConfig."+req.SN, func() (any, error) {
		return g.PullConfig(req)
	})
	if err != nil {
		return nil, err
	}
	return v.(*GatewayPullConfigResponse), nil
}

func (g *Gateway) PullConfig(req *GatewayPullConfigRequest) (resp *GatewayPullConfigResponse, err error) {
	err = g.Client.Call("gateway.PullConfig", req, &resp)
	return
}
