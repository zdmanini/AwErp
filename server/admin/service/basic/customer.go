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

type BasicCustomerService interface {
	List(page request.PageReq, req req.BasicCustomerQueryReq) (res response.PageResp, e error)
	Detail(id string) (res resp.BasicCustomerResp, e error)
	Add(req req.BasicCustomerAddReq) (e error)
	Edit(req req.BasicCustomerEditReq) (e error)
	Del(id string) (e error)
	Change(id string) (e error)
}

type basicCustomerService struct {
	db *gorm.DB
}

func (b basicCustomerService) List(page request.PageReq, req req.BasicCustomerQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := b.db.Model(&basic.BasicCustomer{})
	if len(req.Name) > 0 {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if len(req.Code) > 0 {
		query = query.Where("code like ?", "%"+req.Code+"%")
	}
	if len(req.Phone) > 0 {
		query = query.Where("phone like ?", "%"+req.Phone+"%")
	}
	if len(req.Gender) > 0 {
		query = query.Where("gender = ?", req.Gender)
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
	var items []basic.BasicCustomer
	query.Count(&count)
	err := query.Limit(limit).Offset(offset).Order("id desc").Find(&items).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var respItems []resp.BasicCustomerResp
	response.Copy(&respItems, items)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    respItems,
	}, nil
}

func (b basicCustomerService) Detail(id string) (res resp.BasicCustomerResp, e error) {
	//TODO implement me
	var item basic.BasicCustomer
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	response.Copy(&res, item)
	return
}

func (b basicCustomerService) Add(req req.BasicCustomerAddReq) (e error) {
	//TODO implement me
	var item basic.BasicCustomer
	var count int64
	query := b.db.Model(&basic.BasicCustomer{})
	if len(req.Phone) > 1 {
		query = query.Where("phone = ?", req.Phone).Where("phone != ''")
	}
	query.Count(&count)
	if count > 0 {
		return fmt.Errorf("手机号已存在")
	}
	response.Copy(&item, req)
	err := b.db.Create(&item).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	return
}

func (b basicCustomerService) Edit(req req.BasicCustomerEditReq) (e error) {
	//TODO implement me
	var item basic.BasicCustomer
	err := b.db.Where("id = ?", req.ID).First(&item).Error
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	var count int64
	query := b.db.Model(&basic.BasicCustomer{})
	if len(req.Phone) > 1 {
		query = query.Where("phone = ?", req.Phone).Where("phone != ''")
	}
	query.Where("id != ?", req.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("手机号已存在")
	}
	response.Copy(&item, req)
	err = b.db.Model(&basic.BasicCustomer{}).Where("id = ?", req.ID).Updates(item).Error
	if e = response.CheckErr(err, "Edit Updates err"); e != nil {
		return
	}
	return
}

func (b basicCustomerService) Del(id string) (e error) {
	//TODO implement me
	err := b.db.Where("id = ?", id).Delete(&basic.BasicCustomer{}).Error
	if e = response.CheckErr(err, "Del Delete err"); e != nil {
		return
	}
	return
}

func (b basicCustomerService) Change(id string) (e error) {
	//TODO implement me
	var item basic.BasicCustomer
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

// NewBasicCustomerService 初始化
func NewBasicCustomerService(db *gorm.DB) BasicCustomerService {
	return &basicCustomerService{db: db}
}
