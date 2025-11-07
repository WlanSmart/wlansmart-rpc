package gateway

import (
	"context"
	"errors"
	"time"

	rpc "github.com/WlanSmart/wlansmart-rpc/v2"
	"github.com/abxuz/b-tools/v2/brpc/http"
	"golang.org/x/sync/singleflight"
)

const (
	RpcName        = "gateway"
	DefaultTimeout = 5 * time.Second
)

var (
	ErrModemNotFound = errors.New("modem not found")
)

type (
	Config struct {
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

type PullConfigRequest struct {
	SN       string
	Firmware string
}

type PullConfigResponse = Config

type GatewayService interface {
	PullConfig(*PullConfigRequest, *PullConfigResponse) error
}

func RegisterService(server rpc.RegisterNameServer, service GatewayService) error {
	return server.RegisterName(RpcName, service)
}

type GatewayAdapter struct {
	cli *http.Client
	sg  *singleflight.Group
}

func NewAdapter(cli *http.Client) *GatewayAdapter {
	return &GatewayAdapter{
		cli: cli,
		sg:  new(singleflight.Group),
	}
}

func (a *GatewayAdapter) PullConfigSingleflight(req *PullConfigRequest) (resp *PullConfigResponse, err error) {
	v, err, _ := a.sg.Do("PullConfig."+req.SN, func() (any, error) {
		return a.PullConfig(req)
	})
	if err != nil {
		return nil, err
	}
	return v.(*PullConfigResponse), nil
}

func (a *GatewayAdapter) PullConfig(req *PullConfigRequest) (resp *PullConfigResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), DefaultTimeout)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".PullConfig", req, &resp)
	return
}
