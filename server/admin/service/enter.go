package service

import (
	"Awesome/admin/service/basic"
	"Awesome/admin/service/cloth"
	"Awesome/admin/service/common"
	"Awesome/admin/service/setting"
	"Awesome/admin/service/system"
)

var InitFunctions = []interface{}{
	// common
	common.NewAlbumService,
	common.NewIndexService,
	common.NewUploadService,
	// setting
	setting.NewSettingCopyrightService,
	setting.NewSettingDictDataService,
	setting.NewSettingDictTypeService,
	setting.NewSettingProtocolService,
	setting.NewSettingStorageService,
	setting.NewSettingWebsiteService,
	// system
	system.NewSystemAuthAdminService,
	system.NewSystemAuthDeptService,
	system.NewSystemAuthMenuService,
	system.NewSystemAuthPermService,
	system.NewSystemAuthPostService,
	system.NewSystemAuthRoleService,
	system.NewSystemLoginService,
	system.NewSystemLogsServer,
	// basic
	basic.NewBasicCustomerService,
	basic.NewBasicSupplierService,
	basic.NewBasicMaterialService,
	basic.NewBasicInfoService,
	// cloth
	cloth.NewClothStyleService,
	cloth.NewClothOrderService,
}
