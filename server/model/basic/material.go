package basic

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"likeadmin/model"
	"strings"
)

type BasicMaterial struct {
	model.Model
	Name       string         `gorm:"column:name;not null;unique;default:'';comment:'名称'" json:"name"`
	Code       string         `gorm:"column:code;default:'';comment:'编码'" json:"code"`
	Type       string         `gorm:"column:type;default:'';comment:'类型'" json:"type"`
	SupplierId uint           `gorm:"column:supplier_id;default:0;comment:'供应商ID'" json:"supplier_id"`
	Supplier   *BasicSupplier `gorm:"foreignkey:SupplierId" json:"supplier,omitempty"`
	UnitPrice  float64        `gorm:"column:unit_price;default:0;comment:'单价';decimal:(10,2)" json:"unit_price"`
	Unit       string         `gorm:"column:unit;default:'';comment:'单位'" json:"unit"`
	Status     uint8          `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
	Number     float64        `gorm:"column:number;default:0;comment:'数量';decimal:(10,2)" json:"number"`
	Remark     string         `gorm:"column:remark;default:'';comment:'备注'" json:"remark"`
}

func (BasicMaterial) TableName() string {
	return "basic_materials"
}

func (bm *BasicMaterial) BeforeCreate(tx *gorm.DB) (err error) {
	bm.Model.ID = uuid.NewV3(uuid.NewV1(), "material")
	if bm.Code == "" {
		bm.Code = strings.Split(bm.Model.ID.String(), "-")[0]
	}
	return
}

func (bm *BasicMaterial) AfterCreate(tx *gorm.DB) (err error) {
	if len(bm.Type) > 0 {
		var count int64
		err = tx.Model(&BasicInfo{}).Where("type = ? AND name = ?", "物料类型", bm.Type).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			newType := &BasicInfo{
				Name: bm.Type,
				Type: "物料类型",
				Code: uuid.NewV3(uuid.NewV1(), "BasicInfo").String(),
			}
			err = tx.Create(newType).Error
			if err != nil {
				return
			}
		}
	}
	return
}

func (bm *BasicMaterial) AfterUpdate(tx *gorm.DB) (err error) {
	if len(bm.Type) > 0 {
		var count int64
		err = tx.Model(&BasicInfo{}).Where("type = ? AND name = ?", "物料类型", bm.Type).Count(&count).Error
		if err != nil {
			return
		}
		if count == 0 {
			newType := &BasicInfo{
				Name: bm.Type,
				Type: "物料类型",
				Code: uuid.NewV3(uuid.NewV1(), "BasicInfo").String(),
			}
			err = tx.Create(newType).Error
			if err != nil {
				return
			}
		}
	}
	return
}
