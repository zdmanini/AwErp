package req

import (
	"likeadmin/core/request"
)

type BasicCustomerQueryReq struct {
	Name   string `form:"name"`
	Code   string `form:"code"`
	Gender string `form:"gender"`
	Phone  string `form:"phone"`
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type BasicCustomerAddReq struct {
	Name    string `json:"name" binding:"required,min=2,max=30"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Remark  string `json:"remark"`
	Status  uint8  `json:"status" binding:"oneof=0 1"`
}

type BasicCustomerEditReq struct {
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required,min=2,max=30"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Remark  string `json:"remark"`
	Status  uint8  `json:"status" binding:"oneof=0 1"`
}

type BasicSupplierQueryReq struct {
	Name   string `form:"name"`
	Code   string `form:"code"`
	Gender string `form:"gender"`
	Phone  string `form:"phone"`
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type BasicSupplierAddReq struct {
	Name    string `json:"name" binding:"required,min=2,max=30"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Remark  string `json:"remark"`
	Status  uint8  `json:"status" binding:"oneof=0 1"`
}

type BasicSupplierEditReq struct {
	ID      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required,min=2,max=30"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Remark  string `json:"remark"`
	Status  uint8  `json:"status" binding:"oneof=0 1"`
}

type BasicMaterialQueryReq struct {
	Name       string `form:"name"`
	Code       string `form:"code"`
	Type       string `form:"type"`
	SupplierId int    `form:"supplier_id,default=-1"`
	Status     int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type BasicMaterialAddReq struct {
	Name       string  `json:"name" binding:"required,min=2,max=30"`
	Code       string  `json:"code"`
	Type       string  `json:"type"`
	SupplierId int     `json:"supplier_id"`
	UnitPrice  float64 `json:"unit_price" binding:"required,gt=0"`
	Unit       string  `json:"unit"`
	Remark     string  `json:"remark"`
	Status     uint8   `json:"status" binding:"oneof=0 1"`
}

type BasicMaterialEditReq struct {
	ID         string  `json:"id" binding:"required"`
	Name       string  `json:"name" binding:"required,min=2,max=30"`
	Code       string  `json:"code"`
	Type       string  `json:"type"`
	SupplierId int     `json:"supplier_id"`
	UnitPrice  float64 `json:"unit_price" binding:"required,gt=0"`
	Unit       string  `json:"unit"`
	Remark     string  `json:"remark"`
	Status     uint8   `json:"status" binding:"oneof=0 1"`
}

type BasicInfoItemQueryReq struct {
	Name   string `form:"name"`
	Code   string `form:"code"`
	Type   string `form:"type"`
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type BasicInfoItemAddReq struct {
	Name   string `json:"name" binding:"required,min=1,max=30"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Status uint8  `json:"status" binding:"oneof=0 1"`
	Remark string `json:"remark"`
}

type BasicInfoItemEditReq struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required,min=1,max=30"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Status uint8  `json:"status" binding:"oneof=0 1"`
	Remark string `json:"remark"`
}

type BasicInfoTypeQueryReq struct {
	Name string `form:"name"`
	request.TimeRangeReq
}

type BasicInfoTypeAddReq struct {
	Name   string `json:"name" binding:"required,min=1,max=30"`
	Parent string `json:"parent"`
}

type BasicInfoTypeEditReq struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required,min=1,max=30"`
	Parent string `json:"parent"`
}

type BasicInfoQueryReq struct {
	Type string `form:"type"`
}
