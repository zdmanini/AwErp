package cloth

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"likeadmin/model"
	"likeadmin/model/basic"
	"strings"
)

type ClothStyle struct {
	model.Model
	Name        string        `gorm:"type:varchar(255);not null" json:"name"`
	Code        string        `gorm:"type:varchar(255);not null;default:''" json:"code"`
	Picture     string        `gorm:"type:varchar(255);not null;default:''" json:"picture"`
	Colors      []string      `gorm:"type:json;serializer:json" json:"colors"`
	Sizes       []string      `gorm:"type:json;serializer:json" json:"sizes"`
	Year        string        `gorm:"type:varchar(255);not null;default:''" json:"year"`
	Season      string        `gorm:"type:varchar(255);not null;default:''" json:"season"`
	UnitPrice   float64       `gorm:"type:decimal(10,2);not null;default:0.00" json:"unit_price"`
	Procedures  []Procedure   `gorm:"type:json;serializer:json" json:"procedures"`
	Crafts      []Craft       `gorm:"type:json;serializer:json" json:"crafts"`
	Consumption []Consumption `gorm:"type:json;serializer:json" json:"consumption"`
	Status      uint8         `gorm:"type:tinyint(1);not null;default:0" json:"status"`
	Remark      string        `gorm:"type:varchar(255);not null;default:''" json:"remark"`
	model.SoftDelete
}

type Craft struct {
	Name        string  `json:"name,omitempty" form:"name"`
	Sort        uint    `json:"sort,omitempty" form:"sort"`
	Description string  `json:"description,omitempty" form:"description"`
	Children    []Craft `json:"children" form:"children"`
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

func (u *ClothStyle) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "customer")
	if u.Code == "" {
		u.Code = strings.Split(u.Model.ID.String(), "-")[0]
	}
	return
}

func (u *ClothStyle) AfterCreate(tx *gorm.DB) error {
	return checkInfo(u, tx)
}

// AfterUpdate ...
func (u *ClothStyle) AfterUpdate(tx *gorm.DB) error {
	return checkInfo(u, tx)
}

// checkInfo
func checkInfo(u *ClothStyle, tx *gorm.DB) error {
	if len(u.Colors) > 0 {
		for _, c := range u.Colors {
			var count int64
			if err := tx.Model(&basic.BasicInfo{}).Where("type = ? AND name = ?", "颜色", c).Count(&count).Error; err != nil {
				continue
			}
			if count == 0 {
				newColor := &basic.BasicInfo{
					Model: model.Model{
						ID: uuid.NewV3(uuid.NewV1(), "info"),
					},
					Name: c,
					Type: "颜色",
					Code: strings.Split(u.Model.ID.String(), "-")[0],
				}
				if err := tx.Create(newColor).Error; err != nil {
					return err
				}
			}
		}
	}
	if len(u.Sizes) > 0 {
		for _, size := range u.Sizes {
			var count int64
			if err := tx.Model(&basic.BasicInfo{}).Where("type = ? AND name = ?", "尺码", size).Count(&count).Error; err != nil {
				continue
			}
			if count == 0 {
				newSize := &basic.BasicInfo{
					Model: model.Model{
						ID: uuid.NewV3(uuid.NewV1(), "info"),
					},
					Name: size,
					Type: "尺码",
					Code: strings.Split(u.Model.ID.String(), "-")[0],
				}
				if err := tx.Create(newSize).Error; err != nil {
					return err
				}
			}
		}
	}
	if len(u.Procedures) > 0 {
		for _, procedure := range u.Procedures {
			var count int64
			if err := tx.Model(&basic.BasicInfo{}).Where("type = ? AND name = ?", "工序", procedure.Name).Count(&count); err != nil {
				continue
			}
			if count == 0 {
				newProcedure := &basic.BasicInfo{
					Model: model.Model{
						ID: uuid.NewV3(uuid.NewV1(), "info"),
					},
					Name: procedure.Name,
					Type: "工序",
					Code: strings.Split(u.Model.ID.String(), "-")[0],
				}
				if err := tx.Create(newProcedure).Error; err != nil {
					return err
				}
			}
		}
	}

	if len(u.Crafts) > 0 {
		var names []string
		extractNames(u.Crafts, &names)
		for _, name := range names {
			var count int64
			if err := tx.Model(&basic.BasicInfo{}).Where("type = ? AND name = ?", "工艺", name).Count(&count).Error; err != nil {
				color.Redln("工艺名重复:")
				continue
			}
			if count == 0 {
				newCraft := &basic.BasicInfo{
					Model: model.Model{
						ID: uuid.NewV3(uuid.NewV1(), "info"),
					},
					Name: name,
					Type: "工艺",
					Code: strings.Split(u.Model.ID.String(), "-")[0],
				}
				if err := tx.Create(newCraft).Error; err != nil {
					color.Redln("工艺名重复:", err.Error())
					return err
				}
			}
		}
	}
	return nil
}

// extractNames 递归提取Craft数组中所有Craft的Name字段
func extractNames(crafts []Craft, names *[]string) {
	for _, craft := range crafts {
		*names = append(*names, craft.Name)
		if len(craft.Children) > 0 {
			extractNames(craft.Children, names)
		}
	}
}
