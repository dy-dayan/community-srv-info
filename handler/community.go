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

func (h *Handle) AddCommunity(ctx context.Context, req *srv.AddCommunityReq, rsp *srv.AddCommunityResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	// 获取一个自增id
	cl := atomicid.NewAtomicIDService("dayan.common.srv.atomicid", micro.Client())
	req1 := &atomicid.GetIDReq{Label: "dayan.community.srv.community.community_id"}
	rsp1, err := cl.GetID(ctx, req1)
	if err != nil {
		logrus.Errorf("atomicid.GetID error:%v", err)
		return err
	}

	if rsp1.BaseResp.Code != int32(base.CODE_OK) {
		logrus.Warnf("atomicid.GetID resp code:%v, msg:%s", rsp1.BaseResp.Code, rsp1.BaseResp.Msg)
		rsp.BaseResp = rsp1.BaseResp
		return nil
	}

	info := &db.CommunityInfo{
		ID:           rsp1.Id,
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
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) DelCommunity(ctx context.Context, req *srv.DelCommunityReq, rsp *srv.DelCommunityResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	err := db.DelCommunityInfo(req.CommunityID)
	if err != nil {
		logrus.Warnf("db.DelCommunityInfo error:%v", err)
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle) GetCommunity(ctx context.Context, req *srv.GetCommunityReq, rsp *srv.GetCommunityResp) error {
	rsp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	info, err := db.GetCommunityInfoByID(req.CommunityID)
	if err != nil {
		logrus.Warnf("db.GetCommunityInfoByID error:%v", err)
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = err.Error()
		return nil
	}

	rsp.Community.Name = info.Name
	rsp.Community.Province = info.Province
	rsp.Community.City = info.City
	rsp.Community.Region = info.Region
	rsp.Community.Street = info.Street
	rsp.Community.CouncilID = info.CouncilID
	rsp.Community.OrgID = info.OrgID
	rsp.Community.HouseCount = info.HouseCount
	rsp.Community.CheckInCount = info.CheckInCount
	rsp.Community.BuildingArea = info.BuildingArea
	rsp.Community.GreeningArea = info.GreeningArea
	rsp.Community.Loc = info.Loc
	rsp.Community.State = info.State
	rsp.Community.OperatorID = info.OperatorID
	rsp.CreatedAt = info.CreatedAt
	rsp.UpdatedAt = info.UpdatedAt

	return nil
}
