package rpc

import (
	"errors"

	"github.com/abxuz/b-tools/brpc/http"
	"golang.org/x/sync/singleflight"
)

const (
	RpcNameGateway = "gateway"
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
		SSID struct {
			SSID1 SSID
			SSID2 SSID
			SSID3 SSID
			SSID4 SSID
			SSID5 SSID
			SSID6 SSID
			SSID7 SSID
			SSID8 SSID
		}
		Eth struct {
			Eth1 Eth
			Eth2 Eth
			Eth3 Eth
			Eth4 Eth
		}
		Route struct {
			Enable bool
			Vlan   int
		}
		Command string
	}

	SSID struct {
		Name    string
		Key     string
		Enable  bool
		Channel string
		Hide    bool
		Bridge  Bridge
	}

	Eth struct {
		Enable bool
		Bridge Bridge
	}

	Bridge struct {
		Enable bool
		Vlan   int
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
	err = g.Client.Call(RpcNameGateway+".PullConfig", req, &resp)
	return
}
