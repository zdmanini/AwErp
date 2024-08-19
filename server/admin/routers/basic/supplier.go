package basic

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/service/basic"
	"Awesome/core"
	"Awesome/core/request"
	"Awesome/core/response"
	"Awesome/middleware"
	"Awesome/util"
	"github.com/gin-gonic/gin"
)

var SupplierGroup = core.Group("/basic", newSupplierHandler, regSupplier, middleware.TokenAuth())

func newSupplierHandler(srv basic.BasicSupplierService) *supplierHandler {
	return &supplierHandler{srv: srv}
}

func regSupplier(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *supplierHandler) {
		rg.GET("/supplier/list", handle.list)
		rg.GET("/supplier/detail", handle.detail)
		rg.POST("/supplier/add", middleware.RecordLog("供应商新增"), handle.add)
		rg.POST("/supplier/edit", middleware.RecordLog("供应商编辑"), handle.edit)
		rg.POST("/supplier/del", middleware.RecordLog("供应商删除"), handle.del)
		rg.POST("/supplier/change", middleware.RecordLog("供应商状态"), handle.change)
	})
}

type supplierHandler struct {
	srv basic.BasicSupplierService
}

// list 供应商列表
func (sh supplierHandler) list(c *gin.Context) {
	var listReq req.BasicSupplierQueryReq
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

// detail 供应商详情
func (sh supplierHandler) detail(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &idReq)) {
		return
	}
	res, err := sh.srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 供应商新增
func (sh supplierHandler) add(c *gin.Context) {
	var addReq req.BasicSupplierAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := sh.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 供应商编辑
func (sh supplierHandler) edit(c *gin.Context) {
	var editReq req.BasicSupplierEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := sh.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 供应商删除
func (sh supplierHandler) del(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &idReq)) {
		return
	}
	err := sh.srv.Del(idReq.ID)
	response.CheckAndResp(c, err)
}

// change 供应商状态
func (sh supplierHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := sh.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)

}
