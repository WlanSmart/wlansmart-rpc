package rpc

import (
	"time"

	"github.com/abxuz/b-tools/brpc/http"
)

type Portal struct {
	Client *http.Client
}

type (
	PortalBindRequest struct {
		Store     string
		Mode      string
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
		MaxClient   int
		BindClient  int
	}
)

func (p *Portal) Bind(req *PortalBindRequest) (resp *PortalBindResponse, err error) {
	err = p.Client.Call("portal.Bind", req, &resp)
	return
}

type (
	PortalSendCodeRequet struct {
		Phone string
	}

	PortalSendCodeResponse struct {
		Code string
	}
)

func (p *Portal) SendCode(req *PortalSendCodeRequet) (resp *PortalSendCodeResponse, err error) {
	err = p.Client.Call("portal.SendCode", req, &resp)
	return
}
