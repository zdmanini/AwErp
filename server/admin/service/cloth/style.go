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

type ClothStyleService interface {
	List(page request.PageReq, req req.ClothStyleQueryReq) (res response.PageResp, e error)
	Detail(id string) (res resp.ClothStyleResp, e error)
	Add(req req.ClothStyleAddReq) (e error)
	Edit(req req.ClothStyleEditReq) (e error)
	Del(id string) (e error)
	Change(id string) (e error)
}

type basicStyleService struct {
	db *gorm.DB
}

func (b basicStyleService) List(page request.PageReq, req req.ClothStyleQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	query := b.db.Model(&cloth.ClothStyle{})
	if len(req.Name) > 0 {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if len(req.Code) > 0 {
		query = query.Where("code like ?", "%"+req.Code+"%")
	}
	if len(req.Year) > 0 {
		query = query.Where("year like ?", "%"+req.Year+"%")
	}
	if len(req.Season) > 0 {
		query = query.Where("season like ?", "%"+req.Season+"%")
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
	var list []cloth.ClothStyle
	if err := query.Count(&count).Error; err != nil {
		return
	}
	if err := query.Limit(limit).Offset(offset).Order("create_time desc").Find(&list).Error; err != nil {
		return
	}
	var resList []resp.ClothStyleResp
	response.Copy(&resList, list)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    resList,
	}, nil
}

func (b basicStyleService) Detail(id string) (res resp.ClothStyleResp, e error) {
	//TODO implement me
	var item cloth.ClothStyle
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	response.Copy(&res, item)
	return
}

func (b basicStyleService) Add(req req.ClothStyleAddReq) (e error) {
	//TODO implement me
	var item cloth.ClothStyle
	response.Copy(&item, req)
	err := b.db.Model(&cloth.ClothStyle{}).Create(&item).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	return
}

func (b basicStyleService) Edit(req req.ClothStyleEditReq) (e error) {
	//TODO implement me
	var item cloth.ClothStyle
	err := b.db.Where("id = ?", req.ID).First(&item).Error
	if e = response.CheckErr(err, "Edit Find err"); e != nil {
		return
	}
	if len(item.ID.String()) == 0 {
		return fmt.Errorf("Edit Find err")
	}
	response.Copy(&item, req)
	err = b.db.Model(&cloth.ClothStyle{}).Where("id = ?", req.ID).Updates(&item).Error
	if e = response.CheckErr(err, "Edit Update err"); e != nil {
		return
	}
	return
}

func (b basicStyleService) Del(id string) (e error) {
	//TODO implement me
	err := b.db.Where("id = ?", id).Delete(&cloth.ClothStyle{}).Error
	if e = response.CheckErr(err, "Del Delete err"); e != nil {
		return
	}
	return
}

func (b basicStyleService) Change(id string) (e error) {
	//TODO implement me
	var item cloth.ClothStyle
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Change Find err"); e != nil {
		return
	}
	if len(item.ID.String()) == 0 {
		return fmt.Errorf("Change Find err")
	}
	item.Status = 1 - item.Status
	err = b.db.Model(&cloth.ClothStyle{}).Select("status").Where("id = ?", id).Updates(&item).Error
	if e = response.CheckErr(err, "Change Update err"); e != nil {
		return
	}
	return
}

// NewClothStyleService returns a new ClothStyleService
func NewClothStyleService(db *gorm.DB) ClothStyleService {
	return &basicStyleService{db: db}
}
