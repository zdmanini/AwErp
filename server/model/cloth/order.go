package cloth

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"likeadmin/model"
	"likeadmin/model/basic"
	"strings"
)

type ClothOrder struct {
	model.Model
	Name         string               `gorm:"type:varchar(255);not null" json:"name"`
	Code         string               `gorm:"type:varchar(255);not null;default:''" json:"code"`
	CustomerId   uint                 `gorm:"type:int(11);not null" json:"customer_id"`
	Customer     *basic.BasicCustomer `gorm:"foreignkey:CustomerId" json:"customer,omitempty"`
	DeliveryDate int64                `gorm:"type:bigint(20);not null" json:"delivery_date"`
	OrderType    string               `gorm:"type:varchar(255);not null;default:''" json:"order_type"`
	Salesman     string               `gorm:"type:varchar(255);not null;default:''" json:"salesman"`
	Cloth        *OrderStyle          `gorm:"type:json;serializer:json" json:"cloth"`
	Total        uint                 `gorm:"type:int(11);not null;default:0" json:"total"`
	TotalPrice   float64              `gorm:"type:decimal(10,2);not null;default:0.00" json:"total_price"`
	Contains     []OrderContain       `gorm:"type:json;serializer:json" json:"contains"`
	Procedure    []Procedure          `gorm:"type:json;serializer:json" json:"procedure"`
	Consumption  []Consumption        `gorm:"type:json;serializer:json" json:"consumption"`
	Crafts       []Craft              `gorm:"type:json;serializer:json" json:"crafts"`
	Status       uint8                `gorm:"type:tinyint(1);not null;default:1" json:"status"`
	Remark       string               `gorm:"type:varchar(255);not null;default:''" json:"remark"`
	model.SoftDelete
}

type OrderStyle struct {
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Picture     string        `json:"picture"`
	Colors      []string      `json:"colors"`
	Sizes       []string      `json:"sizes"`
	Year        string        `json:"year"`
	Season      string        `json:"season"`
	UnitPrice   float64       `json:"unitPrice"`
	Procedures  []Procedure   `json:"procedures"`
	Crafts      []Craft       `json:"crafts"`
	Consumption []Consumption `json:"consumption"`
	Status      uint8         `json:"status"`
	Remark      string        `json:"remark"`
}

type OrderContain struct {
	Size  string `json:"size" form:"size" default:""`
	Color string `json:"color" form:"color" default:""`
	Nums  uint   `json:"nums" form:"nums" default:"0"`
}

func (u *ClothOrder) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "order")
	if u.Code == "" {
		u.Code = strings.Split(u.Model.ID.String(), "-")[0]
	}
	return
}

func (u *ClothOrder) AfterCreate(tx *gorm.DB) error {
	return checkOrderInfo(u, tx)
}

// AfterUpdate ...
func (u *ClothOrder) AfterUpdate(tx *gorm.DB) error {
	return checkOrderInfo(u, tx)
}

// checkOrderInfo ...
func checkOrderInfo(u *ClothOrder, tx *gorm.DB) error {
	if u.Cloth != nil {
		var style ClothStyle
		tx.Model(&ClothStyle{}).Where("code = ?", u.Cloth.Code).First(&style)
		if len(u.Cloth.Code) == 0 || len(style.ID.String()) == 0 {
			model.Copy(&style, u.Cloth)
			style.Remark = fmt.Sprintf("订单编号：%s[%s] 的款式", u.Name, u.Code)
			if err := tx.Create(&style).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
