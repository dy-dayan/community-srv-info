package handler

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl"
	atomicid "github.com/dy-dayan/community-srv-info/idl/dayan/common/srv-atomicid"
	srv "github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
	"github.com/micro/go-micro"
)

func (h *Handle) AddAsset(ctx context.Context, req *srv.AddAssetReq,
	resp *srv.AddAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	cl := atomicid.NewAtomicIDService("dayan.common.srv.atomicid", micro.Client())
	req1 := &atomicid.GetIDReq{Label: "dayan.community.srv.asset.asset_id"}
	rsp1, err := cl.GetID(ctx, req1)

	if err != nil {
		logrus.Errorf("atomicid.GetID error:%v", err)
		return err
	}

	if rsp1.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Warnf("atomicid.GetID resp code:%v, msg:%s", rsp1.BaseResp.Code, rsp1.BaseResp.Msg)
		resp.BaseResp = rsp1.BaseResp
		return nil
	}

	data := db.Asset{
		ID:          rsp1.Id,
		Serial:      req.Asset.Serial,
		Category:    req.Asset.Category,
		Loc:         req.Asset.Loc,
		State:       req.Asset.State,
		CommunityID: req.Asset.CommunityID,
		Brand:       req.Asset.Brand,
		Desc:        req.Asset.Desc,
		OperatorID:  req.Asset.OperatorID,
	}

	err := db.UpsertAsset(&data)

	resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)

	return nil
}

func (h *Handle) DelAsset(ctx context.Context, req *srv.DelAssetReq,
	resp *srv.DelAssetResp) error {

	return nil
}

func (h *Handle) GetAsset(ctx context.Context, req *srv.GetAssetReq,
	resp *srv.GetAssetResp) error {

	return nil
}
