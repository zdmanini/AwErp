package cloth

import (
	"Awesome/admin/schemas/req"
	"Awesome/core/request"
	"Awesome/core/response"
	"gorm.io/gorm"
)

type ClothTailorService interface {
	List(page request.PageReq, req req.TailorCuttingQueryReq) (res response.PageResp, e error)
	Cutting(req req.TailorCuttingAddReq) (e error)
}

type clothTailorService struct {
	db *gorm.DB
}

func (c clothTailorService) List(page request.PageReq, req req.TailorCuttingQueryReq) (res response.PageResp, e error) {
	//TODO implement me
	//limit := page.PageSize
	//offset := page.PageSize * (page.PageNo - 1)
	//
	//db,err := cloth.GetDB(req.OrderNo, c.db)
	//if err != nil {
	//    return res, fmt.Errorf("get db error: %w", err)
	//}
	//query := db.Model(&cloth.ClothTailor{})
	return response.PageResp{}, e
}

func (c clothTailorService) Cutting(req req.TailorCuttingAddReq) (e error) {
	//TODO implement me
	panic("implement me")
}

func NewClothTailorService(db *gorm.DB) ClothTailorService {
	return &clothTailorService{db: db}
}
