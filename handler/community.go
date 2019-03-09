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

type Handle struct {
}

func (h *Handle) AddCommunity(ctx context.Context, req *srv.AddCommunityReq, resp *srv.AddCommunityResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	// 获取一个自增id
	cl := atomicid.NewAtomicIDService("dayan.common.srv.automicid", micro.Client())
	idReq := &atomicid.GetIDReq{Label: "dayan.community.srv.community.community_id"}
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

	info := &db.CommunityInfo{
		ID:           idResp.Id,
		Name:         req.Community.Name,
		Province:     req.Community.Province,
		City:         req.Community.City,
		Region:       req.Community.Region,
		Street:       req.Community.Street,
		CouncilID:    req.Community.CouncilID,
		OrgID:        req.Community.OrgID,
		HouseCount:   req.Community.HouseCount,
		CheckInCount: req.Community.CheckInCount,
		BuildingArea: req.Community.BuildingArea,
		GreeningArea: req.Community.GreeningArea,
		Loc:          req.Community.Loc,
		State:        req.Community.State,
		OperatorID:   req.Community.OperatorID,
	}
	err = db.UpsertCommunityInfo(info)
	if err != nil {
		logrus.Warnf("db.UpsertCommunityInfo error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) DelCommunity(ctx context.Context, req *srv.DelCommunityReq, resp *srv.DelCommunityResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	err := db.DelCommunityInfo(req.CommunityID)
	if err != nil {
		logrus.Warnf("db.DelCommunityInfo error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) GetCommunity(ctx context.Context, req *srv.GetCommunityReq, resp *srv.GetCommunityResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	info, err := db.GetCommunityInfoByID(req.CommunityID)
	if err != nil {
		logrus.Warnf("db.GetCommunityInfoByID error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	resp.Community.Name = info.Name
	resp.Community.Province = info.Province
	resp.Community.City = info.City
	resp.Community.Region = info.Region
	resp.Community.Street = info.Street
	resp.Community.CouncilID = info.CouncilID
	resp.Community.OrgID = info.OrgID
	resp.Community.HouseCount = info.HouseCount
	resp.Community.CheckInCount = info.CheckInCount
	resp.Community.BuildingArea = info.BuildingArea
	resp.Community.GreeningArea = info.GreeningArea
	resp.Community.Loc = info.Loc
	resp.Community.State = info.State
	resp.Community.OperatorID = info.OperatorID
	resp.CreatedAt = info.CreatedAt
	resp.UpdatedAt = info.UpdatedAt

	return nil
}
