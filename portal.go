package rpc

import (
	"github.com/abxuz/b-tools/brpc/http"
)

type Portal struct {
	Client *http.Client
}

type (
	PortalBindRequest struct {
		Username  string
		Password  string
		ClientIP  string
		ClientMAC string
	}

	PortalBindResponse struct {
		Errno int
		Msg   string
	}
)

func (p *Portal) Bind(req *PortalBindRequest) (resp *PortalBindResponse, err error) {
	err = p.Client.Call("portal.Bind", req, &resp)
	return
}
