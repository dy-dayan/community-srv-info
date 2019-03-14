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

//convertBuilding db model to pb model
func convertBuilding(item *db.Building) *srv.Building {
	return &srv.Building{
		Common: &srv.BuildingCommon{
			Id:          item.ID,
			Name:        item.Name,
			CommunityID: item.CommunityID,
			Period:      item.Period,
			ElevatorIDs: item.ElevatorIDs,
			Loc:         item.Loc,
			OperatorID:  item.OperatorID,
		},
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

//AddBuilding 添加一个建筑物信息
func (h *Handle) AddBuilding(ctx context.Context, req *srv.AddBuildingReq, resp *srv.AddBuildingResp) error {
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
		ElevatorIDs: req.Building.ElevatorIDs,
		CommunityID: req.Building.CommunityID,
		Period:      req.Building.Period,
		OperatorID:  req.Building.OperatorID,
	}

	err = db.UpsertBuilding(&building)
	if err != nil {
		logrus.Warnf("db.UpsertBuiling error : %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	resp.Id = idResp.Id
	return nil
}

//DelBuilding 删除一个建筑物信息
func (h *Handle) DelBuilding(ctx context.Context, req *srv.DelBuildingReq, resp *srv.DelBuildingResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	err := db.DelBuildingByID(req.Id)
	if err != nil {
		logrus.Warnf("db.DelBuilding error %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	return nil
}

//GetBuilding 获得一个社区的建筑物信息
func (h *Handle) GetBuilding(ctx context.Context, req *srv.GetBuildingReq, resp *srv.GetBuildingResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	ret, err := db.GetBuilding(int(req.Limit), int(req.Offset), req.CommunityID)
	if err != nil {
		logrus.Warnf("db.GetBuildingByID error %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}

	for _, item := range *ret {
		tmpItem := item
		tmp := convertBuilding(&tmpItem)
		resp.Buildings = append(resp.Buildings, tmp)
	}
	return nil
}

//GetBuildingByID 查询具体建筑物信息
func (h *Handle) GetBuildingByID(ctx context.Context, req *srv.GetBuildingByIDReq, resp *srv.GetBuildingByIDResp) error {

	ret, err := db.GetBuildingByID(req.Id)
	if err != nil {
		logrus.Warnf("db.GetBuildingByID error %v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}

	if ret == nil {
		resp.BaseResp.Msg = "not find datadl"
		return nil
	}

	resp.Building = convertBuilding(ret)

	return nil
}
