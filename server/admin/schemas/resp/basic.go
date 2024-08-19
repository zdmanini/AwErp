package resp

import (
	"Awesome/core"
	uuid "github.com/satori/go.uuid"
)

type BasicCustomerResp struct {
	ID         uuid.UUID   `json:"id" structs:"id"`
	Name       string      `json:"name" structs:"name"`
	Code       string      `json:"code" structs:"code"`
	Phone      string      `json:"phone" structs:"phone"`
	Company    string      `json:"company" structs:"company"`
	Address    string      `json:"address" structs:"address"`
	Gender     string      `json:"gender" structs:"gender"`
	Remark     string      `json:"remark" structs:"remark"`
	Status     uint8       `json:"status" structs:"status"`
	CreateTime core.TsTime `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime `json:"update_time" structs:"update_time"`
}

type BasicSupplierResp struct {
	ID         uuid.UUID   `json:"id" structs:"id"`
	Name       string      `json:"name" structs:"name"`
	Code       string      `json:"code" structs:"code"`
	Phone      string      `json:"phone" structs:"phone"`
	Company    string      `json:"company" structs:"company"`
	Address    string      `json:"address" structs:"address"`
	Gender     string      `json:"gender" structs:"gender"`
	Remark     string      `json:"remark" structs:"remark"`
	Status     uint8       `json:"status" structs:"status"`
	CreateTime core.TsTime `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime `json:"update_time" structs:"update_time"`
}

type BasicMaterialResp struct {
	ID       uuid.UUID `json:"id" structs:"id"`
	Name     string    `json:"name" structs:"name"`
	Code     string    `json:"code" structs:"code"`
	Type     string    `json:"type" structs:"type"`
	Supplier struct {
		ID   uuid.UUID `json:"id" structs:"id"`
		Name string    `json:"name" structs:"name"`
		Code string    `json:"code" structs:"code"`
	} `json:"supplier" structs:"supplier"`
	UnitPrice  float64     `json:"unit_price" structs:"unit_price"`
	Unit       string      `json:"unit" structs:"unit"`
	Number     int64       `json:"number" structs:"number"`
	Remark     string      `json:"remark" structs:"remark"`
	Status     uint8       `json:"status" structs:"status"`
	CreateTime core.TsTime `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime `json:"update_time" structs:"update_time"`
}

type BasicInfoItemResp struct {
	ID         uuid.UUID   `json:"id" structs:"id"`
	Name       string      `json:"name" structs:"name"`
	Code       string      `json:"code" structs:"code"`
	Type       string      `json:"type" structs:"type"`
	Status     uint8       `json:"status" structs:"status"`
	Remark     string      `json:"remark" structs:"remark"`
	CreateTime core.TsTime `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime `json:"update_time" structs:"update_time"`
}

type BasicInfoTypeResp struct {
	ID         uuid.UUID           `json:"id" structs:"id"`
	Name       string              `json:"name" structs:"name"`
	Parent     string              `json:"parent" structs:"parent"`
	CreateTime core.TsTime         `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime         `json:"update_time" structs:"update_time"`
	Children   []BasicInfoTypeResp `json:"children" structs:"children"`
}
