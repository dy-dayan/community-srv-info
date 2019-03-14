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

//convertHouse 将db model 转为pb model
func convertHouse(house *db.House) *srv.House {
	return &srv.House{
		Common: &srv.HouseCommon{
			Id:         house.ID,
			BuildingID: house.BuildingID,
			Unit:       house.Unit,
			Acreage:    house.Acreage,
			State:      house.State,
			Rental:     house.Rental,
			OperatorID: house.OperatorID,
		},
		CreatedAt: house.CreatedAt,
		UpdatedAt: house.UpdatedAt,
	}
}

//AddHouse 添加一个房屋信息
func (h *Handle) AddHouse(ctx context.Context, req *srv.AddHouseReq, resp *srv.AddHouseResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	cl := atomicid.NewAtomicIDService("dayan.common.srv.atomicid", micro.Client())
	idReq := &atomicid.GetIDReq{Label: "dayan.community.srv.community.building_id"}
	idResp, err := cl.GetID(ctx, idReq)
	if err != nil {
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
	if err != nil {
		logrus.Warnf("db.UpsertHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	resp.Id = idResp.Id
	return nil
}

//DelHouse 删除一个房屋信息
func (h *Handle) DelHouse(ctx context.Context, req *srv.DelHouseReq, resp *srv.DelHouseResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	err := db.DelHouseByID(req.Id)
	if err != nil {
		logrus.Warnf("db.DelHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	return nil
}

//GetHouseByID 通过房屋ID 查询房屋具体信息
func (h *Handle) GetHouseByID(ctx context.Context, req *srv.GetHouseByIDReq, resp *srv.GetHouseByIDResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}

	ret, err := db.GetHouseByID(req.Id)
	if err != nil {
		logrus.Warnf("db.DelHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}
	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}

	resp.House = convertHouse(ret)
	return nil
}

//查询社区中所有房屋
func (h *Handle) GetHouse(ctx context.Context, req *srv.GetHouseReq, resp *srv.GetHouseResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	ret, err := db.GetHouse(int(req.Limit), int(req.Offset), req.CommunityID)

	if err != nil {
		logrus.Warnf("db.GetHouse error :%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
	}

	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}
	for _, item := range *ret {
		tmpItem := item
		tmp := convertHouse(&tmpItem)
		resp.Houses = append(resp.Houses, tmp)
	}
	return nil
}
