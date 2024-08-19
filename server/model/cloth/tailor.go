package cloth

import (
	"Awesome/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ClothTailor struct {
	model.Model
	OrderId uint          `gorm:"column:order_id;not null;default:0" json:"order_id"`
	Info    *TailorInfo   `gorm:"type:json;serializer:json" json:"info,omitempty"`
	Piece   []TailorPiece `gorm:"foreignkey:TailorId" json:"piece,omitempty"`
	model.SoftDelete
}

type TailorInfo struct {
	Total       uint `json:"total,omitempty" form:"total"`
	Completed   uint `json:"completed,omitempty" form:"completed"`
	IsCompleted bool `json:"is_completed,omitempty" form:"is_completed"`
}

type TailorPiece struct {
	model.Model
	TailorId    uint                 `gorm:"type:int(11);not null" json:"tailor_id"`
	Tailor      *ClothTailor         `gorm:"foreignkey:TailorId" json:"tailor,omitempty"`
	OrderId     uint                 `gorm:"type:int(11);not null" json:"order_id"`
	Order       *ClothOrder          `gorm:"foreignkey:OrderId" json:"order,omitempty"`
	OrderNo     string               `gorm:"type:varchar(255);not null;default:''" json:"order_no"`                        // 订单号
	Customer    string               `gorm:"type:varchar(255);not null;default:''" json:"customer"`                        // 客户
	BedNo       string               `gorm:"type:varchar(255);not null;default:''" json:"bed_no"`                          // 床次
	StabNo      uint                 `gorm:"type:varchar(255);not null" json:"stab_no"`                                    // 扎号
	Layer       uint                 `gorm:"type:int(11);not null" json:"layer"`                                           // 拉布层数
	Color       string               `gorm:"type:varchar(255);not null;default:''" json:"color"`                           // 颜色
	Size        string               `gorm:"type:varchar(255);not null;default:''" json:"size"`                            // 尺码
	IsCompleted bool                 `gorm:"type:tinyint(1);not null;default:0;comment:'0:未完成,1:已完成'" json:"is_completed"` // 是否完成
	Process     []TailorPieceProcess `gorm:"type:json;serializer:json" json:"process,omitempty"`
	Production  []Production         `gorm:"foreignkey:PieceId;references:ID" json:"production,omitempty"`
}

type TailorPieceProcess struct {
	Name        string  `json:"name" form:"name"`                 // 工序名称
	Sort        uint    `json:"sort" form:"sort"`                 // 工序排序
	Price       float64 `json:"price" form:"price"`               // 工序价格
	IsCompleted bool    `json:"is_completed" form:"is_completed"` // 是否完成
	CompletedAt int64   `json:"completed_at" form:"completed_at"` // 完成时间
	OperateId   uint    `json:"operate_id" form:"operate_id"`     // 操作人ID
}

type Production struct {
	model.Model
	TailorId      uint         `gorm:"type:int(11);not null" json:"tailor_id"`
	Tailor        *ClothTailor `gorm:"foreignkey:TailorId" json:"tailor,omitempty"`
	OrderNo       string       `gorm:"type:varchar(255);not null;default:''" json:"order_no"`                        // 订单号
	Customer      string       `gorm:"type:varchar(255);not null;default:''" json:"customer"`                        // 客户
	BedNo         string       `gorm:"type:varchar(255);not null;default:''" json:"bed_no"`                          // 床次
	StabNo        uint         `gorm:"type:varchar(255);not null" json:"stab_no"`                                    // 扎号
	Layer         uint         `gorm:"type:int(11);not null" json:"layer"`                                           // 拉布层数
	Color         string       `gorm:"type:varchar(255);not null;default:''" json:"color"`                           // 颜色
	Size          string       `gorm:"type:varchar(255);not null;default:''" json:"size"`                            // 尺码
	PieceId       uint         `gorm:"type:int(11);not null" json:"piece_id"`                                        // 裁片ID // 工序
	IsCompleted   bool         `gorm:"type:tinyint(1);not null;default:0;comment:'0:未完成,1:已完成'" json:"is_completed"` // 是否完成
	ProcedureName string       `gorm:"type:varchar(255);not null;default:''" json:"procedure_name"`                  // 工序名称
	OperatorId    uint         `gorm:"type:int(11);not null" json:"operator_id"`                                     // 操作人ID
	//Operator      *Employee `gorm:"foreignkey:OperatorId" json:"operator,omitempty"`
	Remark string `gorm:"type:varchar(255);not null;default:''" json:"remark"` // 备注

}

func (u *ClothTailor) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "tailor")
	return
}

func (u *TailorPiece) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "tailor")
	return
}

func (u *Production) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "production")
	return
}
