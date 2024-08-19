package dongming

import (
	srv "Awesome/admin/service/system"
	"Awesome/core"
	"Awesome/model/system"
)

func auth() error {
	perm := srv.NewSystemAuthPermService(core.GetDB())
	var roles []system.SystemAuthRole
	err := core.GetDB().Model(&system.SystemAuthRole{}).Find(&roles).Error
	if err != nil {
		return err
	}
	for _, role := range roles {
		_ = perm.CacheRoleMenusByRoleId(role.ID)
	}
	return nil
}
