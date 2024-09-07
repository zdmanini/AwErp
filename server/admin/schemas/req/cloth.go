package req

import "Awesome/core/request"

type ClothStyleQueryReq struct {
	Name   string `form:"name"`
	Code   string `form:"code"`
	Year   string `form:"year"`
	Season string `form:"season"`
	Status int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type ClothStyleAddReq struct {
	Name        string   `json:"name" binding:"required,min=1,max=30"`
	Code        string   `json:"code"`
	Picture     string   `json:"picture"`
	Colors      []string `json:"colors"`
	Sizes       []string
	Year        string        `json:"year"`
	Season      string        `json:"season"`
	UnitPrice   float64       `json:"unit_price"`
	Procedures  []Procedure   `json:"procedures"`
	Consumption []Consumption `json:"consumption"`
	Crafts      []Craft       `json:"crafts"`
	Status      uint8         `json:"status" binding:"oneof=0 1"`
	Remark      string        `json:"remark"`
}

type ClothStyleEditReq struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" binding:"required,min=1,max=30"`
	Code        string   `json:"code"`
	Picture     string   `json:"picture"`
	Colors      []string `json:"colors"`
	Sizes       []string
	Year        string        `json:"year"`
	Season      string        `json:"season"`
	UnitPrice   float64       `json:"unit_price"`
	Procedures  []Procedure   `json:"procedures"`
	Consumption []Consumption `json:"consumption"`
	Crafts      []Craft       `json:"crafts"`
	Status      uint8         `json:"status" binding:"oneof=0 1"`
	Remark      string        `json:"remark"`
}

type Procedure struct {
	Name             string  `json:"name" form:"name"`
	Sort             uint    `json:"sort" form:"sort"`
	Price            float64 `json:"price" form:"price"`
	Indiscriminately bool    `json:"indiscriminately" form:"indiscriminately"`
	IsCompleted      bool    `json:"is_completed" form:"is_completed"`
}

type Consumption struct {
	Name       string  `json:"name" form:"name"`
	Part       string  `json:"part" form:"part"`
	Nums       uint    `json:"nums" form:"nums"`
	Unit       string  `json:"unit" form:"unit"`
	UnitPrice  float64 `json:"unit_price" form:"unit_price"`
	Size       string  `json:"size" form:"size"`
	SizeAllow  bool    `json:"size_allow" form:"size_allow"`
	Color      string  `json:"color" form:"color"`
	ColorAllow bool    `json:"color_allow" form:"color_allow"`
}

type Craft struct {
	Name        string  `json:"name,omitempty" form:"name"`
	Sort        uint    `json:"sort,omitempty" form:"sort"`
	Description string  `json:"description,omitempty" form:"description"`
	Children    []Craft `json:"children" form:"children"`
}

type ClothOrderQueryReq struct {
	Name         string `form:"name"`
	Code         string `form:"code"`
	CustomerId   uint8  `form:"customer_id"`
	DeliveryDate int64  `form:"delivery_date"`
	Status       int8   `form:"status,default=-1" binding:"oneof=-1 0 1"`
	request.TimeRangeReq
}

type ClothOrderAddReq struct {
	Name         string            `json:"name" binding:"required,min=1,max=30"`
	Code         string            `json:"code"`
	CustomerId   uint8             `json:"customer_id"`
	DeliveryDate int64             `json:"delivery_date" binding:"required"`
	OrderType    string            `json:"order_type"`
	Salesman     string            `json:"salesman"`
	Cloth        *ClothStyleAddReq `json:"cloth"`
	Total        uint              `json:"total"`
	TotalPrice   float64           `json:"total_price"`
	Contains     []OrderContain    `json:"contains"`
	Procedure    []Procedure       `json:"procedure"`
	Consumption  []Consumption     `json:"consumption"`
	Crafts       []Craft           `json:"crafts"`
	Status       uint8             `json:"status" binding:"oneof=0 1"`
	Remark       string            `json:"remark"`
}

type OrderContain struct {
	Size  string `json:"size" form:"size" default:""`
	Color string `json:"color" form:"color" default:""`
	Nums  uint   `json:"nums" form:"nums" default:"0"`
}

type ClothOrderEditReq struct {
	ID           string             `json:"id" binding:"required"`
	Name         string             `json:"name" binding:"required,min=1,max=30"`
	Code         string             `json:"code"`
	CustomerId   uint8              `json:"customer_id"`
	DeliveryDate int64              `json:"delivery_date"`
	OrderType    string             `json:"order_type"`
	Salesman     string             `json:"salesman"`
	Cloth        *ClothStyleEditReq `json:"cloth"`
	Total        uint               `json:"total"`
	TotalPrice   float64            `json:"total_price"`
	Contains     []OrderContain     `json:"contains"`
	Procedure    []Procedure        `json:"procedure"`
	Consumption  []Consumption      `json:"consumption"`
	Crafts       []Craft            `json:"crafts"`
	Status       uint8              `json:"status" binding:"oneof=0 1"`
	Remark       string             `json:"remark"`
}

type TailorCuttingQueryReq struct {
	OrderNo  string `form:"order_no"`
	BedNo    string `form:"bed_no"`
	TailorId uint   `form:"tailor_id"`
	request.TimeRangeReq
}

type TailorCuttingAddReq struct {
	OrderNo  string               `form:"json:"order_no" binding:"required"`
	BedNo    string               `form:"json:"bed_no" binding:"required"`
	TailorId uint                 `form:"json:"tailor_id" binding:"required"`
	StabNo   uint                 `json:"stab_no"`
	Piece    []TailorCuttingPiece `json:"piece"`
}
type TailorCuttingPiece struct {
	Color string `json:"color"`
	Size  string `json:"size"`
	Nums  uint   `json:"nums"`
}

type TailorCuttingEditReq struct {
	ID       string               `form:"id" binding:"required"`
	OrderNo  string               `form:"json:"order_no" binding:"required"`
	BedNo    string               `form:"json:"bed_no" binding:"required"`
	TailorId uint                 `form:"json:"tailor_id" binding:"required"`
	StabNo   uint                 `json:"stab_no"`
	Piece    []TailorCuttingPiece `json:"piece"`
}
