package basic

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"likeadmin/model"
	"strings"
)

type BasicInfo struct {
	model.Model
	Name     string         `gorm:"column:name;not null;default:'';unique;" json:"name"`
	Code     string         `gorm:"column:code;default:'';" json:"code"`
	Remark   string         `gorm:"column:remark;default:'';" json:"remark"`
	Type     string         `gorm:"column:type;not null;default:'';comment:字典类型" json:"type"`
	InfoType *BasicInfoType `gorm:"foreignkey:Type;references:Name;comment:字典类型" json:"info_type,omitempty"`
	Status   uint8          `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
}

type BasicInfoType struct {
	model.Model
	Name     string          `gorm:"column:name;not null;default:'';unique;" json:"name"`
	Parent   string          `gorm:"column:parent;default:'';" json:"parent"`
	Children []BasicInfoType `gorm:"-" json:"children,omitempty"`
	model.SoftDelete
}

func (receiver *BasicInfo) BeforeCreate(tx *gorm.DB) (err error) {
	receiver.Model.ID = uuid.NewV3(uuid.NewV1(), "info")
	if len(receiver.Code) == 0 {
		receiver.Code = strings.Split(receiver.Model.ID.String(), "-")[0]
	}
	return
}

func (receiver *BasicInfoType) BeforeCreate(tx *gorm.DB) (err error) {
	receiver.Model.ID = uuid.NewV3(uuid.NewV1(), "info")
	return
}

func GetTree(db *gorm.DB, parent string) ([]BasicInfoType, error) {
	var s []BasicInfoType
	err := db.Where("parent = ?", parent).Order("create_time desc").Find(&s).Error
	if err != nil {
		return nil, err
	}

	for i := range s {
		subCategories, err := GetTree(db, s[i].Name)
		if err != nil {
			return nil, err
		}
		s[i].Children = subCategories
	}

	return s, nil
}
