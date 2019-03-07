package dal

type House struct {
	ID         string `bson:"_id"`
	acreage    面积
	state      使用状态
	rental     出租状态
	created_at 添加时间
	updated_at 更新时间
}
