package cloth

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/cloth"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var OrderGroup = core.Group("/cloth", newOrderHandler, regOrder, middleware.TokenAuth())

func newOrderHandler(srv cloth.ClothOrderService) *orderHandler {
	return &orderHandler{srv: srv}
}

func regOrder(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *orderHandler) {
		rg.GET("/order/list", handle.list)
		rg.GET("/order/detail", handle.detail)
		rg.POST("/order/add", middleware.RecordLog("订单新增"), handle.add)
		rg.POST("/order/edit", middleware.RecordLog("订单编辑"), handle.edit)
		rg.POST("/order/del", middleware.RecordLog("订单删除"), handle.del)
		rg.POST("/order/change", middleware.RecordLog("订单状态"), handle.change)
	})
}

type orderHandler struct {
	srv cloth.ClothOrderService
}

// list 订单列表
func (sh orderHandler) list(c *gin.Context) {
	var listReq req.ClothOrderQueryReq
	var pageReq request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &pageReq)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := sh.srv.List(pageReq, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 订单详情
func (sh orderHandler) detail(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &idReq)) {
		return
	}
	res, err := sh.srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 订单新增
func (sh orderHandler) add(c *gin.Context) {
	var addReq req.ClothOrderAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := sh.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 订单编辑
func (sh orderHandler) edit(c *gin.Context) {
	var editReq req.ClothOrderEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := sh.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 订单删除
func (sh orderHandler) del(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &idReq)) {
		return
	}
	err := sh.srv.Del(idReq.ID)
	response.CheckAndResp(c, err)
}

// change 订单状态
func (sh orderHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := sh.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)

}
