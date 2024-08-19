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

var CustomerGroup = core.Group("/basic", newCustomerHandler, regCustomer, middleware.TokenAuth())

func newCustomerHandler(srv basic.BasicCustomerService) *customerHandler {
	return &customerHandler{srv: srv}
}

func regCustomer(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *customerHandler) {
		rg.GET("/customer/list", handle.list)
		rg.GET("/customer/detail", handle.detail)
		rg.POST("/customer/add", middleware.RecordLog("客户新增"), handle.add)
		rg.POST("/customer/edit", middleware.RecordLog("客户编辑"), handle.edit)
		rg.POST("/customer/del", middleware.RecordLog("客户删除"), handle.del)
		rg.POST("/customer/change", middleware.RecordLog("客户状态"), handle.change)
	})
}

type customerHandler struct {
	srv basic.BasicCustomerService
}

// list 客户列表
func (sh customerHandler) list(c *gin.Context) {
	var listReq req.BasicCustomerQueryReq
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

// detail 客户详情
func (sh customerHandler) detail(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &idReq)) {
		return
	}
	res, err := sh.srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 客户新增
func (sh customerHandler) add(c *gin.Context) {
	var addReq req.BasicCustomerAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := sh.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 客户编辑
func (sh customerHandler) edit(c *gin.Context) {
	var editReq req.BasicCustomerEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := sh.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 客户删除
func (sh customerHandler) del(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &idReq)) {
		return
	}
	err := sh.srv.Del(idReq.ID)
	response.CheckAndResp(c, err)
}

// change 客户状态
func (sh customerHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := sh.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)

}
