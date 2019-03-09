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

func(h *Handle) AddHouse(ctx context.Context, req *srv.AddHouseReq, resp *srv.AddHouseResp) error {
	resp.BaseResp = &base.Resp{
		Code:                 int32(base.CODE_OK),
	}
	cl := atomicid.NewAtomicIDService("dayan.common.srv.atomicid", micro.Client())
	idReq := &atomicid.GetIDReq{Label: "dayan.community.srv.community.building_id"}
	idResp, err := cl.GetID(ctx, idReq)
	if err != nil{
		logrus.Warnf("atomicid.GetID resp code:%v, msg:%s", idResp.BaseResp.Code, idResp.BaseResp.Msg)
		resp.BaseResp = idResp.BaseResp
		return nil
	}

	house := db.House{
		ID:         idResp.Id,
		BuildingID: req.House.BuildingID,
		Unit:       req.House.Unit,
		Acreage:    req.House.Acreage,
		State:      req.House.State,
		Rental:     req.House.Rental,
		OperatorID: req.House.OperatorID,
	}

	err = db.UpsertHouse(&house)
	if err != nil{
		logrus.Warnf("db.UpsertHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	return nil
}

func(h *Handle) DelHouse(ctx context.Context, req *srv.DelHouseReq, resp *srv.DelHouseResp)error{
	resp.BaseResp = &base.Resp{
		Code:                 int32(base.CODE_OK),
	}

	err := db.DelHouseByID(req.ID)

	if err != nil{
		logrus.Warnf("db.DelHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	return nil
}


func (h *Handle)GetHouse(ctx context.Context, req *srv.GetHouseReq, resp *srv.GetHouseResp)error{
resp.BaseResp = &base.Resp{
		Code:                 int32(base.CODE_OK),
	}
	ret , err := db.GetHouseByID(req.ID)

	if err != nil{
logrus.Warnf("db.GetHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	resp.House.OperatorID = ret.OperatorID
	resp.House.Rental = ret.Rental
	resp.House.State = ret.State
	resp.House.Acreage = ret.Acreage
	resp.House.Unit = ret.Unit
	resp.House.BuildingID = ret.BuildingID
	resp.UpdatedAt = ret.UpdatedAt
	resp.CreatedAt = ret.CreatedAt
	return nil
}