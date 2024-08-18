package core

import (
	"gorm.io/gorm"
	"likeadmin/config"
	"likeadmin/model/basic"
	"likeadmin/model/cloth"
	"likeadmin/model/common"
	"likeadmin/model/gen"
	"likeadmin/model/setting"
	"likeadmin/model/system"
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
