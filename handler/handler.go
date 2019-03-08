package handler

import (
	"context"
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl"
	srv "github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
	atomicid "github.com/dy-dayan/community-srv-info/idl/dayan/common/srv-atomicid"
	"github.com/dy-gopkg/kit/micro"
	"github.com/sirupsen/logrus"
)

type Handle struct {

}

func (h *Handle) AddCommunity(ctx context.Context, req *srv.AddCommunityReq, rsp *srv.AddCommunityResp) error {
	rsp.BaseResp = &base.Resp{
		Code:int32(base.CODE_OK),
	}
	// 获取一个自增id
	cl := atomicid.NewAtomicIDService("dayan.common.srv.atomicid", micro.Client())
	req1 := &atomicid.GetIDReq{Label: "dayan.community.srv.proposal.proposal_id"}
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
		Name:         req.Name,
		Province:     req.Province,
		City:         req.City,
		Region:       req.Region,
		Street:       req.Street,
		CouncilID:    req.CouncilID,
		OrgID:        req.OrgID,
		HouseCount:   req.HouseCount,
		CheckInCount: req.CheckInCount,
		BuildingArea: req.BuildingArea,
		GreeningArea: req.GreeningArea,
		Loc:          req.Loc,
		State:        req.State,
		OperatorID:   req.OperatorID,
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
		Code:int32(base.CODE_OK),
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
		Code:int32(base.CODE_OK),
	}

	info, err := db.GetCommunityInfoByID(req.CommunityID)
	if err != nil {
		logrus.Warnf("db.GetCommunityInfoByID error:%v", err)
		rsp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		rsp.BaseResp.Msg = err.Error()
		return nil
	}

	rsp.Name = info.Name
	rsp.Province = info.Province
	rsp.City = info.City
	rsp.Region = info.Region
	rsp.Street = info.Street
	rsp.CouncilID = info.CouncilID
	rsp.OrgID = info.OrgID
	rsp.HouseCount = info.HouseCount
	rsp.CheckInCount = info.CheckInCount
	rsp.BuildingArea = info.BuildingArea
	rsp.GreeningArea = info.GreeningArea
	rsp.Loc = info.Loc
	rsp.State = info.State
	rsp.OperatorID = info.OperatorID
	rsp.CreatedAt = info.CreatedAt
	rsp.UpdatedAt = info.UpdatedAt

	return nil
}