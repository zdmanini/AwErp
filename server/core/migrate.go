package core

import (
	"Awesome/config"
	"Awesome/model/basic"
	"Awesome/model/cloth"
	"Awesome/model/common"
	"Awesome/model/gen"
	"Awesome/model/setting"
	"Awesome/model/system"
	"gorm.io/gorm"
)

func Migrage(db *gorm.DB) error {
	if config.Config.AutoMigrate {
		return db.AutoMigrate(dst()...)
	}
	return nil
}

func dst() []interface{} {
	return []interface{}{
		// system
		&system.SystemConfig{},
		&system.SystemAuthAdmin{},
		&system.SystemAuthMenu{},
		&system.SystemAuthPerm{},
		&system.SystemAuthRole{},
		&system.SystemAuthDept{},
		&system.SystemAuthPost{},
		&system.SystemLogLogin{},
		&system.SystemLogOperate{},
		// gen
		&gen.GenTable{},
		&gen.GenTableColumn{},
		// setting
		&setting.DictType{},
		&setting.DictData{},
		// common
		&common.Album{},
		&common.AlbumCate{},
		// basic
		&basic.BasicCustomer{},
		&basic.BasicSupplier{},
		&basic.BasicMaterial{},
		&basic.BasicInfo{},
		&basic.BasicInfoType{},
		// cloth
		&cloth.ClothStyle{},
	}
}
