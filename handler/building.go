package handler

import (
	"context"
	"github.com/dy-dayan/community-srv-info/dal/db"
	"github.com/dy-dayan/community-srv-info/idl"
	srv "github.com/dy-dayan/community-srv-info/idl/dayan/community/srv-info"
	atomicid "github.com/dy-dayan/community-srv-info/idl/dayan/common/srv-atomicid"
	"github.com/sirupsen/logrus"
	"github.com/dy-gopkg/kit/micro"
)

func(h *Handle) AddBuilding(ctx context.Context, req *srv.AddBuildingReq, resp *srv.AddBuildingResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	// 获取一个自增id
	cl := atomicid.NewAtomicIDService("dayan.common.srv.automicid", micro.Client())
	idReq := &atomicid.GetIDReq{Label: "dayan.community.srv.community.building_id"}
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

	building := db.Building{
		ID:          idResp.Id,
		Name:        req.Building.Name,
		Loc:         req.Building.Loc,
		ElevatorIDs: req.Building.ElevatorID,
		CommunityID: req.Building.CommunityID,
		Period:      req.Building.Period,
		OperatorID:  req.Building.OperatorID,
	}

	err = db.UpsertBuilding(&building)
	if err != nil{
		logrus.Warnf("db.UpsertBuiling error : %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}


func (h *Handle)DelBuilding(ctx context.Context, req *srv.DelBuilingReq, resp *srv.DelBuilingResp)error{
	resp.BaseResp = &base.Resp{
		Code:                 int32(base.CODE_OK),
	}
	err := db.DelBuildingByID(req.ID)
	if err != nil{
		logrus.Warnf("db.DelBuilding error %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

func (h *Handle)GetBuilding(ctx context.Context, req *srv.GetBuildingReq, resp *srv.GetBuildingResp)error{
	resp.BaseResp = &base.Resp{
		Code:                 int32(base.CODE_OK),
	}
	ret, err := db.GetBuildingByID(req.ID)
	if err != nil{
		logrus.Warnf("db.GetBuildingByID error %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	resp.Building.OperatorID = ret.OperatorID
	resp.Building.Period = ret.Period
	resp.Building.CommunityID = ret.CommunityID
	resp.Building.Loc = ret.Loc
	resp.Building.Name = ret.Name

	return nil
}