package routers

import (
	"Awesome/admin/routers/basic"
	"Awesome/admin/routers/cloth"
	"Awesome/admin/routers/common"
	"Awesome/admin/routers/monitor"
	"Awesome/admin/routers/setting"
	"Awesome/admin/routers/system"
	"Awesome/core"
)

var InitRouters = []*core.GroupBase{
	// common
	common.AlbumGroup,
	common.IndexGroup,
	common.UploadGroup,
	// monitor
	monitor.MonitorGroup,
	// setting
	setting.CopyrightGroup,
	setting.DictDataGroup,
	setting.DictTypeGroup,
	setting.ProtocolGroup,
	setting.StorageGroup,
	setting.WebsiteGroup,
	// system
	system.AdminGroup,
	system.DeptGroup,
	system.LogGroup,
	system.LoginGroup,
	system.MenuGroup,
	system.PostGroup,
	system.RoleGroup,
	// basic
	basic.CustomerGroup,
	basic.SupplierGroup,
	basic.MaterialGroup,
	basic.InfoGroup,
	// cloth
	cloth.StyleGroup,
	cloth.OrderGroup,
}
