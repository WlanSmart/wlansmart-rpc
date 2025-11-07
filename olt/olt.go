package olt

import (
	"context"
	"time"

	rpc "github.com/WlanSmart/wlansmart-rpc/v2"
	"github.com/abxuz/b-tools/v2/brpc/http"
)

const (
	RpcName = "olt"

	TimeoutDefault       = 5 * time.Second
	TimeoutListOnu       = 30 * time.Second
	TimeoutListOnuDetail = time.Minute
	TimeoutListOnuVlan   = 15 * time.Second
	TimeoutSaveConfig    = 10 * time.Second
)

type (
	Pon struct {
		Id   int
		Name string
	}

	OnuUncfg struct {
		PonId    int
		OnuId    int
		Sn       string
		Password string
		Loid     string
		Model    string
		Version  string
	}

	Onu struct {
		PonId int
		OnuId int
		Sn    string
	}

	OnuDetail struct {
		PonId   int
		OnuId   int
		Sn      string
		Name    string // Only for Zxan
		Desc    string
		Line    string // Only fir Hw
		State   string
		RxPower float32
		Uptime  int64 // up seconds for Zxan, online timestamp for Hw
	}

	Vlan struct {
		UserVlan int
		Vlan     int
	}
)

type ListPonRequest = struct{}
type ListPonResponse = []*Pon

type ListOnuUncfgRequest = struct{}
type ListOnuUncfgResponse = []*OnuUncfg

type ListOnuRequest = struct{}
type ListOnuResponse = []*Onu

type ListOnuDetailRequest = struct{}
type ListOnuDetailResponse = []*OnuDetail

type GetOnuDetailRequest struct {
	PonId int
	OnuId int
}
type GetOnuDetailResponse = OnuDetail

type AddOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Type  string // required by zxan
	Name  string // required by zxan
	Desc  string
	Tcont string // required by zxan
	Line  string // required by hw
	Vlan  []Vlan
}
type AddOnuResponse = struct{}

type UpdateOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Type  string // required by zxan
	Name  string // required by zxan
	Desc  string
	Tcont string // required by zxan
	Line  string // required by hw
	Vlan  []Vlan
}
type UpdateOnuResponse = struct{}

type UpdateOnuDescRequest struct {
	PonId int
	OnuId int
	Name  string // required by zxan
	Desc  string
}
type UpdateOnuDescResponse = struct{}

type DeleteOnuRequest struct {
	PonId int
	OnuId int
}
type DeleteOnuResponse = struct{}

type ListOnuVlanRequest struct {
	PonId int
	OnuId int
}
type ListOnuVlanResponse = []*Vlan

type SaveConfigRequest = struct{}
type SaveConfigResponse = struct{}

type OltService interface {
	ListPon(*ListPonRequest, *ListPonResponse) error
	ListOnuUncfg(*ListOnuUncfgRequest, *ListOnuUncfgResponse) error
	ListOnu(*ListOnuRequest, *ListOnuResponse) error
	ListOnuDetail(*ListOnuDetailRequest, *ListOnuDetailResponse) error
	GetOnuDetail(*GetOnuDetailRequest, *GetOnuDetailResponse) error
	AddOnu(*AddOnuRequest, *AddOnuResponse) error
	UpdateOnu(*UpdateOnuRequest, *UpdateOnuResponse) error
	UpdateOnuDesc(*UpdateOnuDescRequest, *UpdateOnuDescResponse) error
	DeleteOnu(*DeleteOnuRequest, *DeleteOnuResponse) error
	ListOnuVlan(*ListOnuVlanRequest, *ListOnuVlanResponse) error
	SaveConfig(*SaveConfigRequest, *SaveConfigResponse) error
}

func RegisterService(server rpc.RegisterNameServer, service OltService) error {
	return server.RegisterName(RpcName, service)
}

type OltAdapter struct {
	cli *http.Client
}

func New(cli *http.Client) *OltAdapter {
	return &OltAdapter{cli: cli}
}

func (a *OltAdapter) ListPon() (resp ListPonResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".ListPon", ListPonRequest{}, &resp)
	return
}

func (a *OltAdapter) ListOnuUncfg() (resp ListOnuUncfgResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".ListOnuUncfg", ListOnuUncfgRequest{}, &resp)
	return
}

func (a *OltAdapter) ListOnu() (resp ListOnuResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutListOnu)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".ListOnu", ListOnuRequest{}, &resp)
	return
}

func (a *OltAdapter) ListOnuDetail() (resp ListOnuDetailResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutListOnuDetail)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".ListOnuDetail", ListOnuDetailRequest{}, &resp)
	return
}

func (a *OltAdapter) GetOnuDetail(req *GetOnuDetailRequest) (resp *GetOnuDetailResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".GetOnuDetail", req, &resp)
	return
}
func (a *OltAdapter) AddOnu(req *AddOnuRequest) (err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".AddOnu", req, &AddOnuResponse{})
	return
}

func (a *OltAdapter) UpdateOnu(req *UpdateOnuRequest) (err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".UpdateOnu", req, &UpdateOnuResponse{})
	return
}

func (a *OltAdapter) UpdateOnuDesc(req *UpdateOnuDescRequest) (err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".UpdateOnuDesc", req, &UpdateOnuDescResponse{})
	return
}

func (a *OltAdapter) DeleteOnu(req *DeleteOnuRequest) (err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutDefault)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".DeleteOnu", req, &DeleteOnuResponse{})
	return
}

func (a *OltAdapter) ListOnuVlan(req *ListOnuVlanRequest) (resp ListOnuVlanResponse, err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutListOnuVlan)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".ListOnuVlan", req, &resp)
	return
}

func (a *OltAdapter) SaveConfig() (err error) {
	ctx, done := context.WithTimeout(context.Background(), TimeoutSaveConfig)
	defer done()
	err = a.cli.CallContext(ctx, RpcName+".SaveConfig", SaveConfigRequest{}, &SaveConfigResponse{})
	return
}
