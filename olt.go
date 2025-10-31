package rpc

import (
	"github.com/abxuz/b-tools/brpc/http"
)

const RpcNameZxanOlt = "ZxanOlt"

type ZxanOltPon struct {
	Id   int
	Name string
}

type ZxanOltOnu struct {
	PonId int
	OnuId int
	Sn    string
}

type ZxanOltOnuDetail struct {
	PonId      int
	OnuId      int
	Sn         string
	Name       string
	Desc       string
	State      string
	RxPower    float32
	UptimeSecs int64
}

type ZxanOltOnuUncfg struct {
	PonId    int
	OnuId    int
	Sn       string
	Password string
	Loid     string
	Model    string
	Version  string
}

type ZxanOltVlan struct {
	UserVlan int
	Vlan     int
}

type ZxanOltServicePort struct {
	PonId    int
	OnuId    int
	Index    int
	Vlan     int
	UserVlan int
}

type ZxanOlt struct {
	Client *http.Client
}

type ZxanOltListPonRequest struct{}
type ZxanOltListPonResponse struct {
	List []ZxanOltPon
}

func (olt *ZxanOlt) ListPon(req *ZxanOltListPonRequest) (resp *ZxanOltListPonResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".ListPon", req, &resp)
	return
}

type ZxanOltListOnuRequest struct{}
type ZxanOltListOnuResponse struct {
	List []ZxanOltOnu
}

func (olt *ZxanOlt) ListOnu(req *ZxanOltListOnuRequest) (resp *ZxanOltListOnuResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".ListOnu", req, &resp)
	return
}

type ZxanOltListOnuDetailRequest struct{}
type ZxanOltListOnuDetailResponse struct {
	List []ZxanOltOnuDetail
}

func (olt *ZxanOlt) ListOnuDetail(req *ZxanOltListOnuDetailRequest) (resp *ZxanOltListOnuDetailResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".ListOnuDetail", req, &resp)
	return
}

type ZxanOltAddOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Type  string
	Name  string
	Desc  string
	Tcont string
	Vlan  []ZxanOltVlan
}
type ZxanOltAddOnuResponse struct{}

func (olt *ZxanOlt) AddOnu(req *ZxanOltAddOnuRequest) (resp *ZxanOltAddOnuResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".AddOnu", req, &resp)
	return
}

type ZxanOltUpdateOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Type  string
	Name  string
	Desc  string
	Tcont string
	Vlan  []ZxanOltVlan
}
type ZxanOltUpdateOnuResponse struct{}

func (olt *ZxanOlt) UpdateOnu(req *ZxanOltUpdateOnuRequest) (resp *ZxanOltUpdateOnuResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".UpdateOnu", req, &resp)
	return
}

type ZxanOltDeleteOnuRequest struct {
	PonId int
	OnuId int
}
type ZxanOltDeleteOnuResponse struct{}

func (olt *ZxanOlt) DeleteOnu(req *ZxanOltDeleteOnuRequest) (resp *ZxanOltDeleteOnuResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".DeleteOnu", req, &resp)
	return
}

type ZxanOltListOnuUncfgRequest struct{}
type ZxanOltListOnuUncfgResponse struct {
	List []ZxanOltOnuUncfg
}

func (olt *ZxanOlt) ListOnuUncfg(req *ZxanOltListOnuUncfgRequest) (resp *ZxanOltListOnuUncfgResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".ListOnuUncfg", req, &resp)
	return
}

type ZxanOltListOnuServicePortRequest struct {
	PonId int
	OnuId int
}
type ZxanOltListOnuServicePortResponse struct {
	List []ZxanOltServicePort
}

func (olt *ZxanOlt) ListOnuServicePort(req *ZxanOltListOnuServicePortRequest) (resp *ZxanOltListOnuServicePortResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".ListOnuServicePort", req, &resp)
	return
}

type ZxanOltSaveConfigRequest struct{}
type ZxanOltSaveConfigResponse struct{}

func (olt *ZxanOlt) SaveConfig(req *ZxanOltListOnuServicePortRequest) (resp *ZxanOltListOnuServicePortResponse, err error) {
	err = olt.Client.Call(RpcNameZxanOlt+".SaveConfig", req, &resp)
	return
}

const RpcNameHwOlt = "HwOlt"

type HwOlt struct {
	Client *http.Client
}

type HwOltPon struct {
	Id   int
	Name string
}

type HwOltOnu struct {
	PonId int
	OnuId int
	Sn    string
}

type HwOltOnuDetail struct {
	PonId      int
	OnuId      int
	Sn         string
	Desc       string
	Line       string
	State      string
	RxPower    float32
	OnlineTime int64
}

type HwOltOnuUncfg struct {
	PonId    int
	OnuId    int
	Sn       string
	Password string
	Loid     string
	Model    string
	Version  string
}

type HwOltVlan struct {
	UserVlan int
	Vlan     int
}

type HwOltServicePort struct {
	PonId    int
	OnuId    int
	Index    int
	Vlan     int
	UserVlan int
}

type HwOltListPonRequest struct{}
type HwOltListPonResponse struct {
	List []HwOltPon
}

func (olt *HwOlt) ListPon(req *HwOltListPonRequest) (resp *HwOltListPonResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".ListPon", req, &resp)
	return
}

type HwOltListOnuRequest struct{}
type HwOltListOnuResponse struct {
	List []HwOltOnu
}

func (olt *HwOlt) ListOnu(req *HwOltListOnuRequest) (resp *HwOltListOnuResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".ListOnu", req, &resp)
	return
}

type HwOltListOnuDetailRequest struct{}
type HwOltListOnuDetailResponse struct {
	List []HwOltOnuDetail
}

func (olt *HwOlt) ListOnuDetail(req *HwOltListOnuDetailRequest) (resp *HwOltListOnuDetailResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".ListOnuDetail", req, &resp)
	return
}

type HwOltAddOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Desc  string
	Line  string
	Vlan  []HwOltVlan
}
type HwOltAddOnuResponse struct{}

func (olt *HwOlt) AddOnu(req *HwOltAddOnuRequest) (resp *HwOltAddOnuResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".AddOnu", req, &resp)
	return
}

type HwOltUpdateOnuRequest struct {
	PonId int
	OnuId int
	Sn    string
	Desc  string
	Line  string
	Vlan  []HwOltVlan
}
type HwOltUpdateOnuResponse struct{}

func (olt *HwOlt) UpdateOnu(req *HwOltUpdateOnuRequest) (resp *HwOltUpdateOnuResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".UpdateOnu", req, &resp)
	return
}

type HwOltDeleteOnuRequest struct {
	PonId int
	OnuId int
}
type HwOltDeleteOnuResponse struct{}

func (olt *HwOlt) DeleteOnu(req *HwOltDeleteOnuRequest) (resp *HwOltDeleteOnuResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".DeleteOnu", req, &resp)
	return
}

type HwOltListOnuUncfgRequest struct{}
type HwOltListOnuUncfgResponse struct {
	List []HwOltOnuUncfg
}

func (olt *HwOlt) ListOnuUncfg(req *HwOltListOnuUncfgRequest) (resp *HwOltListOnuUncfgResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".ListOnuUncfg", req, &resp)
	return
}

type HwOltListOnuServicePortRequest struct {
	PonId int
	OnuId int
}
type HwOltListOnuServicePortResponse struct {
	List []HwOltServicePort
}

func (olt *HwOlt) ListOnuServicePort(req *HwOltListOnuServicePortRequest) (resp *HwOltListOnuServicePortResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".ListOnuServicePort", req, &resp)
	return
}

type HwOltSaveConfigRequest struct{}
type HwOltSaveConfigResponse struct{}

func (olt *HwOlt) SaveConfig(req *HwOltListOnuServicePortRequest) (resp *HwOltListOnuServicePortResponse, err error) {
	err = olt.Client.Call(RpcNameHwOlt+".SaveConfig", req, &resp)
	return
}
