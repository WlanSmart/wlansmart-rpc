package rpc

import (
	"context"
	"time"

	"github.com/abxuz/b-tools/brpc/http"
)

const RpcNamePortal = "portal"

type Portal struct {
	Client *http.Client
}

const (
	PortalBindModePassword = byte(0)
	PortalBindModeCode     = byte(1)
)

type (
	PortalBindRequest struct {
		Store     string
		Mode      byte
		Username  string
		Password  string
		ClientIP  string
		ClientMAC string
		Hostname  string
	}

	PortalBindResponse struct {
		PlanName    string
		ExpireTime  time.Time
		WiredEnable bool
		MaxClient   uint
		BindClient  uint
	}
)

func (p *Portal) Bind(req *PortalBindRequest) (resp *PortalBindResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = p.Client.CallContext(ctx, RpcNamePortal+".Bind", req, &resp)
	return
}

type (
	PortalSendCodeRequet struct {
		Store     string
		Phone     string
		Code      string
		ClientIP  string
		ClientMAC string
	}

	PortalSendCodeResponse struct{}
)

func (p *Portal) SendCode(req *PortalSendCodeRequet) (resp *PortalSendCodeResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = p.Client.CallContext(ctx, RpcNamePortal+".SendCode", req, &resp)
	return
}
