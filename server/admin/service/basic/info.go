package basic

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/schemas/resp"
	"Awesome/config"
	"Awesome/core/request"
	"Awesome/core/response"
	"Awesome/model/basic"
	"Awesome/util"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type BasicInfoService interface {
	ItemList(page request.PageReq, listReq req.BasicInfoItemQueryReq) (res response.PageResp, e error)
	TypeList(listReq req.BasicInfoTypeQueryReq) (res []resp.BasicInfoTypeResp, e error)
	Add(req req.BasicInfoItemAddReq) (e error)
	Edit(req req.BasicInfoItemEditReq) (e error)
	Del(id string) (e error)
	TypeAdd(req req.BasicInfoTypeAddReq) (e error)
	TypeEdit(req req.BasicInfoTypeEditReq) (e error)
	TypeDel(id string) (e error)
	Query(req req.BasicInfoQueryReq) (res map[string]interface{}, e error)
	Change(id string) (e error)
}

type basicInfoService struct {
	db *gorm.DB
}

func (b basicInfoService) TypeAdd(req req.BasicInfoTypeAddReq) (e error) {
	//TODO implement me
	var count int64
	if err := b.db.Model(&basic.BasicInfoType{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该类型已存在")
	}
	var item basic.BasicInfoType
	response.Copy(&item, req)
	return b.db.Model(&basic.BasicInfoType{}).Create(&item).Error
}

func (b basicInfoService) TypeEdit(req req.BasicInfoTypeEditReq) (e error) {
	//TODO implement me
	var item basic.BasicInfoType
	if err := b.db.Model(&basic.BasicInfoType{}).Where("id = ?", req.ID).First(&item).Error; err != nil {
		return err
	}
	var count int64
	if err := b.db.Model(&basic.BasicInfoType{}).Where("name = ? and id != ?", req.Name, req.ID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该类型已存在")
	}
	return b.db.Model(&basic.BasicInfoType{}).Select("name").Updates(&req).Error
}

func (b basicInfoService) Change(id string) (e error) {
	//TODO implement me
	var item basic.BasicInfo
	if err := b.db.Model(&basic.BasicInfo{}).Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}
	if item.ID.String() == "" {
		return fmt.Errorf("数据不存在")
	}
	item.Status = 1 - item.Status
	return b.db.Model(&basic.BasicInfo{}).Select("status").Where("id = ?", id).Updates(&item).Error
}

func (b basicInfoService) Query(req req.BasicInfoQueryReq) (res map[string]interface{}, e error) {
	//TODO implement me
	res = make(map[string]interface{})
	arr := strings.Split(req.Type, ",")
	for _, v := range arr {
		var item []basic.BasicInfo
		if err := b.db.Model(&basic.BasicInfo{}).Where("type = ? and status = ?", v, 1).Order("create_time desc").Find(&item).Error; err != nil {
			continue
		}
		res[v] = item
	}
	return res, nil
}

func (b basicInfoService) Del(id string) (e error) {
	//TODO implement me
	return b.db.Where("id = ?", id).Delete(&basic.BasicInfo{}).Error
}

func (b basicInfoService) TypeDel(id string) (e error) {
	//TODO implement me
	var item basic.BasicInfoType
	if err := b.db.Model(&basic.BasicInfoType{}).Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}
	var count int64
	if err := b.db.Model(&basic.BasicInfo{}).Where("type = ?", item.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("该类型下存在数据, 无法删除")
	}
	err := b.db.Model(&basic.BasicInfoType{}).Where("parent = ?", item.Name).Count(&count).Error
	if err != nil {
		return fmt.Errorf("该类型下存在子类型，无法删除")
	}
	if util.ArrayUtil.InArray(config.AdminConfig.InfoWhitelist, item.Name) {
		return fmt.Errorf("该类型为系统内置类型，无法删除")
	}

	return b.db.Where("id = ?", id).Delete(&basic.BasicInfoType{}).Error
}

func (b basicInfoService) ItemList(page request.PageReq, listReq req.BasicInfoItemQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := b.db.Model(&basic.BasicInfo{})
	if len(listReq.Name) > 0 {
		query = query.Where("name like ?", "%"+listReq.Name+"%")
	}
	if len(listReq.Type) > 0 {
		query = query.Where("type = ?", listReq.Type)
	}
	if len(listReq.Code) > 0 {
		query = query.Where("code like ?", "%"+listReq.Code+"%")
	}
	if listReq.Status >= 0 {
		query = query.Where("status = ?", listReq.Status)
	}
	if listReq.TimeRangeReq.StartTime > 0 {
		query = query.Where("create_time >= ?", listReq.TimeRangeReq.StartTime)
	}
	if listReq.TimeRangeReq.EndTime > 0 {
		query = query.Where("create_time <= ?", listReq.TimeRangeReq.EndTime)
	}

	var count int64
	var items []basic.BasicInfo
	query.Count(&count)
	query.Limit(limit).Offset(offset).Order("create_time desc").Find(&items)
	var respItems []resp.BasicInfoItemResp
	response.Copy(&respItems, items)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    respItems,
	}, nil
}

func (b basicInfoService) TypeList(listReq req.BasicInfoTypeQueryReq) (res []resp.BasicInfoTypeResp, e error) {
	//TODO implement me
	trees, err := basic.GetTree(b.db, "顶级分类")
	if err != nil {
		return nil, err
	}
	var respItems []resp.BasicInfoTypeResp
	response.Copy(&respItems, trees)
	return respItems, nil
}

func (b basicInfoService) Add(req req.BasicInfoItemAddReq) (e error) {
	//TODO implement me
	var count int64
	b.db.Model(&basic.BasicInfo{}).Where("name = ? and type = ?", req.Name, req.Type).Count(&count)
	if count > 0 {
		return fmt.Errorf("name=%s,type=%s 已存在", req.Name, req.Type)
	}
	var info basic.BasicInfo
	response.Copy(&info, req)
	err := b.db.Model(&basic.BasicInfo{}).Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

func (b basicInfoService) Edit(req req.BasicInfoItemEditReq) (e error) {
	//TODO implement me
	var info basic.BasicInfo
	err := b.db.Where("id = ?", req.ID).First(&info).Error
	if err != nil {
		return err
	}
	var count int64
	b.db.Model(&basic.BasicInfo{}).Where("name = ? and type = ? and id != ?", req.Name, req.Type, req.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("name=%s,type=%s 已存在", req.Name, req.Type)
	}
	response.Copy(&info, req)
	err = b.db.Model(&basic.BasicInfo{}).Where("id = ?", req.ID).Updates(&info).Error
	if err != nil {
		return err
	}
	return nil
}

// NewBasicInfoService 创建BasicInfoService
func NewBasicInfoService(db *gorm.DB) BasicInfoService {
	return &basicInfoService{db: db}
}
