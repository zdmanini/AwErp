package basic

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"likeadmin/model"
	"strings"
)

type BasicCustomer struct {
	model.Model
	Name    string `gorm:"column:name;not null;default:''" json:"name"`
	Code    string `gorm:"column:code;default:'';" json:"code"`
	Phone   string `gorm:"column:phone;default:'';" json:"phone"`
	Company string `gorm:"column:company;default:'';" json:"company"`
	Address string `gorm:"column:address;default:'';" json:"address"`
	Gender  string `gorm:"column:gender;default:'';" json:"gender"`
	Remark  string `gorm:"column:remark;default:'';" json:"remark"`
	Status  uint8  `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`
}

func (u *BasicCustomer) BeforeCreate(tx *gorm.DB) (err error) {
	u.Model.ID = uuid.NewV3(uuid.NewV1(), "customer")
	if u.Code == "" {
		u.Code = strings.Split(u.Model.ID.String(), "-")[0]
	}
	return
}
