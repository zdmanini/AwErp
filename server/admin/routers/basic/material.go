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

var MaterialGroup = core.Group("/basic", newMaterialHandler, regMaterial, middleware.TokenAuth())

func newMaterialHandler(srv basic.BasicMaterialService) *materialHandler {
	return &materialHandler{srv: srv}
}

func regMaterial(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *materialHandler) {
		rg.GET("/material/list", handle.list)
		rg.GET("/material/detail", handle.detail)
		rg.POST("/material/add", middleware.RecordLog("物料新增"), handle.add)
		rg.POST("/material/edit", middleware.RecordLog("物料编辑"), handle.edit)
		rg.POST("/material/del", middleware.RecordLog("物料删除"), handle.del)
		rg.POST("/material/change", middleware.RecordLog("物料状态"), handle.change)
	})
}

type materialHandler struct {
	srv basic.BasicMaterialService
}

// list 物料列表
func (sh materialHandler) list(c *gin.Context) {
	var listReq req.BasicMaterialQueryReq
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

// detail 物料详情
func (sh materialHandler) detail(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &idReq)) {
		return
	}
	res, err := sh.srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 物料新增
func (sh materialHandler) add(c *gin.Context) {
	var addReq req.BasicMaterialAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := sh.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 物料编辑
func (sh materialHandler) edit(c *gin.Context) {
	var editReq req.BasicMaterialEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := sh.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 物料删除
func (sh materialHandler) del(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &idReq)) {
		return
	}
	err := sh.srv.Del(idReq.ID)
	response.CheckAndResp(c, err)
}

// change 物料状态
func (sh materialHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := sh.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)

}
