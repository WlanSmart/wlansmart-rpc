package portal

import (
	"context"
	"time"

	rpc "github.com/WlanSmart/wlansmart-rpc/v2"
	"github.com/abxuz/b-tools/v2/brpc/http"
)

const (
	RpcName        = "portal"
	DefaultTimeout = 5 * time.Second
)

const (
	BindModePassword = byte(0)
	BindModeCode     = byte(1)
)

type BindRequest struct {
	Store     string
	Mode      byte
	Username  string
	Password  string
	ClientIP  string
	ClientMAC string
	Hostname  string
}
type BindResponse struct {
	PlanName    string
	ExpireTime  time.Time
	WiredEnable bool
	MaxClient   uint
	BindClient  uint
}

type SendCodeRequet struct {
	Store     string
	Phone     string
	Code      string
	ClientIP  string
	ClientMAC string
}
type SendCodeResponse = struct{}

type PortalService interface {
	Bind(*BindRequest, *BindResponse) error
	SendCode(*SendCodeRequet, *SendCodeResponse) error
}

func RegisterService(server rpc.RegisterNameServer, service PortalService) error {
	return server.RegisterName(RpcName, service)
}

type PortalAdapter struct {
	cli *http.Client
}

func NewAdapter(cli *http.Client) *PortalAdapter {
	return &PortalAdapter{cli: cli}
}

func (a *PortalAdapter) Bind(req *BindRequest) (resp *BindResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), DefaultTimeout)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".Bind", req, &resp)
	return
}

func (a *PortalAdapter) SendCode(req *SendCodeRequet) (err error) {
	ctx, done := context.WithTimeout(context.Background(), DefaultTimeout)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".SendCode", req, &SendCodeResponse{})
	return
}
