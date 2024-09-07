package cloth

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/schemas/resp"
	"Awesome/core/request"
	"Awesome/core/response"
	"Awesome/model/cloth"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type ClothOrderService interface {
	List(page request.PageReq, req req.ClothOrderQueryReq) (res response.PageResp, e error)
	Detail(id string) (res resp.ClothOrderResp, e error)
	Add(req req.ClothOrderAddReq) (e error)
	Edit(req req.ClothOrderEditReq) (e error)
	Del(id string) (e error)
	Change(id string) (e error)
}

type clothOrderService struct {
	db *gorm.DB
}

func (b clothOrderService) List(page request.PageReq, req req.ClothOrderQueryReq) (res response.PageResp, e error) {
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
	//var resList []resp.ClothOrderResp
	//response.Copy(&resList, list)
	data := make([]resp.ClothOrderResp, 0)
	// 创建一个带缓冲的channel，用于接收goroutine的结果
	results := make(chan resp.ClothOrderResp, len(list))
	// 创建一个WaitGroup，用于等待所有goroutine完成
	var wg sync.WaitGroup
	// channel的容量为30，表示最多只能有30个goroutine同时执行
	channel := make(chan bool, 30)

	for _, v := range list {
		wg.Add(1)
		go func(v cloth.ClothOrder) {
			channel <- true
			// 等待channel通知
			<-channel
			// 从channel中取出数据
			defer wg.Done()
			var r resp.ClothOrderResp
			response.Copy(&r, v)
			var progress resp.ClothOrderProgress
			db, err := cloth.GetDB(v.ID.String(), b.db)
			if db == nil || err != nil {
				progress = resp.ClothOrderProgress{
					Total:   int64(v.Total),
					Done:    0,
					Percent: 0,
					Cuted:   0,
				}
				r.Progress = progress
				results <- r
				return
			}
			var pieces []cloth.TailorPiece
			db.Model(&cloth.TailorPiece{}).Find(&pieces)
			progress.Total = int64(v.Total)
			var completed int64
			for _, v := range pieces {
				if v.IsCompleted {
					completed++
				}
			}
			progress.Done = completed
			progress.Percent = float64(progress.Done) / float64(progress.Total) * 100
			progress.Cuted = progress.Total - progress.Done
			r.Progress = progress
			results <- r
		}(v)
	}
	// 等待所有goroutine完成
	wg.Wait()
	close(results)
	for v := range results {
		data = append(data, v)
	}

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    data,
	}, nil
}

func (b clothOrderService) Detail(id string) (res resp.ClothOrderResp, e error) {
	//TODO implement me
	var item cloth.ClothOrder
	err := b.db.Where("id = ?", id).First(&item).Error
	if e = response.CheckErr(err, "Detail Find err"); e != nil {
		return
	}
	response.Copy(&res, item)
	return
}

func (b clothOrderService) Add(req req.ClothOrderAddReq) (e error) {
	//TODO implement me
	var item cloth.ClothOrder
	response.Copy(&item, req)
	err := b.db.Model(&cloth.ClothOrder{}).Create(&item).Error
	if e = response.CheckErr(err, "Add Create err"); e != nil {
		return
	}
	return
}

func (b clothOrderService) Edit(req req.ClothOrderEditReq) (e error) {
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

func (b clothOrderService) Del(id string) (e error) {
	//TODO implement me
	err := b.db.Where("id = ?", id).Delete(&cloth.ClothOrder{}).Error
	if e = response.CheckErr(err, "Del Delete err"); e != nil {
		return
	}
	return
}

func (b clothOrderService) Change(id string) (e error) {
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
	return &clothOrderService{db: db}
}
