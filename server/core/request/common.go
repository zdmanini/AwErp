package request

// PageReq 分页请求参数
type PageReq struct {
	PageNo   int `form:"pageNo,default=1" validate:"omitempty,gte=1"`            // 页码
	PageSize int `form:"pageSize,default=20" validate:"omitempty,gt=0,lte=1000"` // 每页大小
}

// TimeRangeReq 时间范围请求参数
type TimeRangeReq struct {
	StartTime int64 `form:"startTime" validate:"omitempty"` // 开始时间
	EndTime   int64 `form:"endTime" validate:"omitempty"`   // 结束时间
}

// IDReq 主键请求参数
type IDReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// IDsReq 主键列表请求参数
type IDsReq struct {
	Ids []string `form:"ids" binding:"required"` // 主键列表
}
