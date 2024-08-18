package basic

import (
	"fmt"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/basic"
)

type BasicMaterialService interface {
	List(page request.PageReq, req req.BasicMaterialQueryReq) (res response.PageResp, e error)
	Detail(id string) (res resp.BasicMaterialResp, e error)
	Add(req req.BasicMaterialAddReq) (e error)
	Edit(req req.BasicMaterialEditReq) (e error)
	Del(id string) (e error)
	Change(id string) (e error)
}

type basicMaterialService struct {
	db *gorm.DB
}

func (b basicMaterialService) List(page request.PageReq, req req.BasicMaterialQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := b.db.Model(&basic.BasicMaterial{})
	if len(req.Name) > 0 {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if len(req.Code) > 0 {
		query = query.Where("code like ?", "%"+req.Code+"%")
	}
	if len(req.Type) > 0 {
		query = query.Where("type = ?", "%"+req.Type+"%")
	}
	if req.SupplierId > 0 {
		query = query.Where("supplier_id = ?", req.SupplierId)
	}
	if req.TimeRangeReq.StartTime > 0 {
		query = query.Where("create_time >= ?", req.TimeRangeReq.StartTime)
	}
	if req.TimeRangeReq.EndTime > 0 {
		query = query.Where("create_time <= ?", req.TimeRangeReq.EndTime)
	}
	if req.Status >= 0 {
		query = query.Where("status = ?", req.Status)
	}
	var count int64
	var items []basic.BasicMaterial
	query.Count(&count)
	err := query.Limit(limit).Offset(offset).Order("id desc").Find(&items).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var respItems []resp.BasicMaterialResp
	response.Copy(&respItems, items)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    respItems,
	}, nil
}

func (b basicMaterialService) Detail(id string) (res resp.BasicMaterialResp, e error) {
	//TODO implement me
	var item basic.BasicMaterial
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	response.Copy(&res, item)
	return
}

func (b basicMaterialService) Add(req req.BasicMaterialAddReq) (e error) {
	//TODO implement me
	var item basic.BasicMaterial
	var count int64
	b.db.Model(&basic.BasicMaterial{}).Where("type = ? and name = ?", req.Type, req.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("已存在相同物料")
	}
	response.Copy(&item, req)
	err := b.db.Create(&item).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	return
}

func (b basicMaterialService) Edit(req req.BasicMaterialEditReq) (e error) {
	//TODO implement me
	var item basic.BasicMaterial
	err := b.db.Where("id = ?", req.ID).First(&item).Error
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	var count int64
	b.db.Model(&basic.BasicMaterial{}).Where("type = ? and name = ?", req.Type, req.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("已存在相同物料")
	}
	response.Copy(&item, req)
	err = b.db.Model(&basic.BasicMaterial{}).Where("id = ?", req.ID).Updates(item).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
	return
}

func (b basicMaterialService) Del(id string) (e error) {
	//TODO implement me
	err := b.db.Where("id = ?", id).Delete(&basic.BasicMaterial{}).Error
	if e = response.CheckErr(err, "Del Delete err"); e != nil {
		return
	}
	return
}

func (b basicMaterialService) Change(id string) (e error) {
	//TODO implement me
	var item basic.BasicMaterial
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Change Find err"); e != nil {
		return
	}
	// 是否显示: 0=否, 1=是
	item.Status = 1 - item.Status
	err = b.db.Model(&item).Select("status").Updates(item).Error
	if e = response.CheckErr(err, "Change Updates err"); e != nil {
		return
	}
	return
}

// NewBasicMaterialService 初始化
func NewBasicMaterialService(db *gorm.DB) BasicMaterialService {
	return &basicMaterialService{db: db}
}
