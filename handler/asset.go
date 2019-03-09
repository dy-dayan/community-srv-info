package handler

import (
	"context"
	"github.com/dy-gopkg/kit/micro"
	"github.com/sirupsen/logrus"
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl"
	atomicid "github.com/dy-dayan/community-srv-info/idl/dayan/common/srv-atomicid"
	srv "github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
)

func (h *Handle) AddAsset(ctx context.Context, req *srv.AddAssetReq,
	resp *srv.AddAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	cl := atomicid.NewAtomicIDService("dayan.common.srv.automicid",micro.Client())
	idReq := &atomicid.GetIDReq{Label: "dayan.community.srv.community.asset_id"}
	idResp, err := cl.GetID(ctx, idReq)

	if err != nil {
		logrus.Errorf("atomicid.GetID error:%v", err)
		return err
	}

	if idResp.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Warnf("atomicid.GetID resp code:%v, msg:%s", idResp.BaseResp.Code, idResp.BaseResp.Msg)
		resp.BaseResp = idResp.BaseResp
		return nil
	}

	data := db.Asset{
		ID:          idResp.Id,
		Serial:      req.Asset.Serial,
		Category:    req.Asset.Category,
		Loc:         req.Asset.Loc,
		State:       req.Asset.State,
		CommunityID: req.Asset.CommunityID,
		Brand:       req.Asset.Brand,
		Desc:        req.Asset.Desc,
		OperatorID:  req.Asset.OperatorID,
	}

	err = db.UpsertAsset(&data)

	resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
	if err != nil {
		logrus.Warnf("db.UpsertAsset error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) DelAsset(ctx context.Context, req *srv.DelAssetReq,
	resp *srv.DelAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code:int32(base.CODE_OK),
	}

	err := db.DelAssetByID(req.AssetID)
	if err != nil{
		logrus.Warnf("db.DelAsset erroro:%v",err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) GetAsset(ctx context.Context, req *srv.GetAssetReq,
	resp *srv.GetAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code:int32(base.CODE_OK),
	}

	asset, err := db.GetAssetByID(req.AssetID)
	if err != nil{
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	resp.Asset.OperatorID = asset.OperatorID
	resp.Asset.Desc = asset.Desc
	resp.Asset.Brand = asset.Brand
	resp.Asset.CommunityID = asset.CommunityID
	resp.Asset.State = asset.State
	resp.Asset.Loc = asset.Loc
	resp.Asset.Category = asset.Category
	resp.Asset.Serial = asset.Serial
	return nil
}
