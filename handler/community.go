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

type Handle struct{}

// convertCommunity 将db中的对象转换为pb对象
func convertCommunity(item *db.CommunityInfo) *srv.Community {
	return &srv.Community{
		Common: &srv.CommunityCommon{
			Id:           item.ID,
			Name:         item.Name,
			Province:     item.Province,
			City:         item.City,
			Region:       item.Region,
			Street:       item.Street,
			SerialNumber: item.SerialNumber,
			OrgID:        item.OrgID,
			HouseCount:   item.HouseCount,
			CheckInCount: item.CheckInCount,
			BuildingArea: item.BuildingArea,
			GreeningArea: item.GreeningArea,
			Loc:          item.Loc,
			State:        item.State,
			SealedState:  item.SealedState,
			OperatorID:   item.OperatorID,
		},
		UpdatedAt: item.UpdatedAt,
		CreatedAt: item.CreatedAt,
	}
}

//AddCommunity 添加一个社区信息
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
		SerialNumber: req.Community.SerialNumber,
		Province:     req.Community.Province,
		City:         req.Community.City,
		Region:       req.Community.Region,
		Street:       req.Community.Street,
		OrgID:        req.Community.OrgID,
		HouseCount:   req.Community.HouseCount,
		CheckInCount: req.Community.CheckInCount,
		BuildingArea: req.Community.BuildingArea,
		GreeningArea: req.Community.GreeningArea,
		Loc:          req.Community.Loc,
		State:        req.Community.State,
		SealedState:  req.Community.SealedState,
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

//DelCommunity 删除一个社区
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

//GetCommunityByID 获得一个社区信息
func (h *Handle) GetCommunityByID(ctx context.Context, req *srv.GetCommunityByIDReq, resp *srv.GetCommunityByIDResp) error {
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

	if info == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}

	resp.Community = convertCommunity(info)
	return nil
}

//GetCommunity 获得所有社区信息
func (h *Handle) GetCommunity(ctx context.Context, req *srv.GetCommunityReq, resp *srv.GetCommunityResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	ret, err := db.GetCommunityInfo(int(req.Limit), int(req.Offset))
	if err != nil {
		logrus.Warnf("db.GetCommunityInfoByID error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}
	for _, item := range *ret {
		tmpValue := item
		tmp := convertCommunity(&tmpValue)
		resp.Communitys = append(resp.Communitys, tmp)
	}
	return nil
}

//GetCommunityByLoc
func (h *Handle) GetCommunityByLoc(ctx context.Context, req *srv.GetCommunityByLocReq, resp *srv.GetCommunityByLocResp) error {
	resp.BaseResp = &base.Resp{
		Code: int32(base.CODE_OK),
	}
	ret, err := db.GetCommunityInfoByLoc(int(req.Limit), int(req.Offset), req.Loc, req.Distance)
	if err != nil {
		logrus.Warnf("db.GetCommunityInfoByID error:%v", err)
		resp.BaseResp.Code = int32(base.CODE_DATA_EXCEPTION)
		resp.BaseResp.Msg = err.Error()
		return nil
	}
	if ret == nil {
		resp.BaseResp.Msg = "not find data"
		return nil
	}
	for _, item := range *ret {
		tmpValue := item
		tmp := convertCommunity(&tmpValue)
		resp.Communitys = append(resp.Communitys, tmp)
	}
	return nil
}
