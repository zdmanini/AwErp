package cloth

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/schemas/resp"
	"Awesome/core/request"
	"Awesome/core/response"
	"Awesome/model/cloth"
	"fmt"
	"gorm.io/gorm"
)

type ClothOrderService interface {
	List(page request.PageReq, req req.ClothOrderQueryReq) (res response.PageResp, e error)
	Detail(id string) (res resp.ClothOrderResp, e error)
	Add(req req.ClothOrderAddReq) (e error)
	Edit(req req.ClothOrderEditReq) (e error)
	Del(id string) (e error)
	Change(id string) (e error)
}

type basicOrderService struct {
	db *gorm.DB
}

func (b basicOrderService) List(page request.PageReq, req req.ClothOrderQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := b.db.Model(&cloth.ClothOrder{})
	if len(req.Name) > 0 {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if len(req.Code) > 0 {
		query = query.Where("code like ?", "%"+req.Code+"%")
	}
	if req.CustomerId > 0 {
		query = query.Where("customer_id = ?", req.CustomerId)
	}
	if req.DeliveryDate > 0 {
		query = query.Where("delivery_date = ?", req.DeliveryDate)
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
	var list []cloth.ClothOrder
	if err := query.Count(&count).Error; err != nil {
		return
	}
	if err := query.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error; err != nil {
		return
	}
	var resList []resp.ClothOrderResp
	response.Copy(&resList, list)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resList,
	}, nil
}

func (b basicOrderService) Detail(id string) (res resp.ClothOrderResp, e error) {
	//TODO implement me
	var item cloth.ClothOrder
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	response.Copy(&res, item)
	return
}

func (b basicOrderService) Add(req req.ClothOrderAddReq) (e error) {
	//TODO implement me
	var item cloth.ClothOrder
	response.Copy(&item, req)
	err := b.db.Model(&cloth.ClothOrder{}).Create(&item).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	return
}

func (b basicOrderService) Edit(req req.ClothOrderEditReq) (e error) {
	//TODO implement me
	var item cloth.ClothOrder
	err := b.db.Where("id = ?", req.ID).First(&item).Error
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	if len(item.ID.String()) == 0 {
		return fmt.Errorf("Edit Find err")
	}
	response.Copy(&item, req)
	err = b.db.Model(&cloth.ClothOrder{}).Where("id = ?", req.ID).Updates(&item).Error
	if e = response.CheckErr(err, "Edit Update err"); e != nil {
		return
	}
	return
}

func (b basicOrderService) Del(id string) (e error) {
	//TODO implement me
	err := b.db.Where("id = ?", id).Delete(&cloth.ClothOrder{}).Error
	if e = response.CheckErr(err, "Del Delete err"); e != nil {
		return
	}
	return
}

func (b basicOrderService) Change(id string) (e error) {
	//TODO implement me
	var item cloth.ClothOrder
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Change Find err"); e != nil {
		return
	}
	if len(item.ID.String()) == 0 {
		return fmt.Errorf("Change Find err")
	}
	item.Status = 1 - item.Status
	err = b.db.Model(&cloth.ClothOrder{}).Select("status").Where("id = ?", id).Updates(&item).Error
	if e = response.CheckErr(err, "Change Update err"); e != nil {
		return
	}
	return
}

// NewClothOrderService returns a new ClothOrderService
func NewClothOrderService(db *gorm.DB) ClothOrderService {
	return &basicOrderService{db: db}
}
