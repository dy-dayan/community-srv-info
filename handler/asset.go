package handler

import (
	"context"
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl"
	atomicid "github.com/dy-dayan/community-srv-info/idl/dayan/common/srv-atomicid"
	srv "github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
	"github.com/dy-gopkg/kit/micro"
	"github.com/sirupsen/logrus"
)

//convertAsset 转db model to pb model
func convertAsset(item *db.Asset) *srv.Asset {
	return &srv.Asset{
		Common: &srv.AssetCommon{
			Id:           item.ID,
			SerialNumber: item.SerialNumber,
			Category:     item.Category,
			State:        item.State,
			CommunityID:  item.CommunityID,
			Loc:          item.Loc,
			Brand:        item.Brand,
			Desc:         item.Desc,
			OperatorID:   item.OperatorID,
		},
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

//AddAsset  添加一个资产
func (h *Handle) AddAsset(ctx context.Context, req *srv.AddAssetReq,
	resp *srv.AddAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	cl := atomicid.NewAtomicIDService("dayan.common.srv.automicid", micro.Client())
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
		ID:           idResp.Id,
		SerialNumber: req.Asset.SerialNumber,
		Category:     req.Asset.Category,
		Loc:          req.Asset.Loc,
		State:        req.Asset.State,
		CommunityID:  req.Asset.CommunityID,
		Brand:        req.Asset.Brand,
		Desc:         req.Asset.Desc,
		OperatorID:   req.Asset.OperatorID,
	}

	err = db.UpsertAsset(&data)

	resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
	if err != nil {
		logrus.Warnf("db.UpsertAsset error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	resp.AssetID = idResp.Id
	return nil
}

//DelAsset 删除一个资产
func (h *Handle) DelAsset(ctx context.Context, req *srv.DelAssetReq,
	resp *srv.DelAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	err := db.DelAssetByID(req.AssetID)
	if err != nil {
		logrus.Warnf("db.DelAsset erroro:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

//GetAssetByID 通过资产ID 查询一个资产信息
func (h *Handle) GetAssetByID(ctx context.Context, req *srv.GetAssetByIDReq, resp *srv.GetAssetByIDResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	asset, err := db.GetAssetByID(req.AssetID)
	if err != nil {
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	//没有查询到数据
	if asset == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}
	resp.Asset = convertAsset(asset)
	return nil
}

//GetAsset 通过社区ID 查询社区内的资产
func (h *Handle) GetAsset(ctx context.Context, req *srv.GetAssetReq,
	resp *srv.GetAssetResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	ret, err := db.GetAsset(int(req.Limit), int(req.Offset), req.CommunityID)
	if err != nil {
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	//没有查询到数据
	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}

	for _, item := range *ret {
		tmpItem := item
		tmp := convertAsset(&tmpItem)
		resp.Assets = append(resp.Assets, tmp)
	}
	return nil
}
