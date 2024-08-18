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

var StyleGroup = core.Group("/cloth", newStyleHandler, regStyle, middleware.TokenAuth())

func newStyleHandler(srv cloth.ClothStyleService) *styleHandler {
	return &styleHandler{srv: srv}
}

func regStyle(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *styleHandler) {
		rg.GET("/style/list", handle.list)
		rg.GET("/style/detail", handle.detail)
		rg.POST("/style/add", middleware.RecordLog("款式新增"), handle.add)
		rg.POST("/style/edit", middleware.RecordLog("款式编辑"), handle.edit)
		rg.POST("/style/del", middleware.RecordLog("款式删除"), handle.del)
		rg.POST("/style/change", middleware.RecordLog("款式状态"), handle.change)
	})
}

type styleHandler struct {
	srv cloth.ClothStyleService
}

// list 款式列表
func (sh styleHandler) list(c *gin.Context) {
	var listReq req.ClothStyleQueryReq
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

// detail 款式详情
func (sh styleHandler) detail(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &idReq)) {
		return
	}
	res, err := sh.srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 款式新增
func (sh styleHandler) add(c *gin.Context) {
	var addReq req.ClothStyleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := sh.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 款式编辑
func (sh styleHandler) edit(c *gin.Context) {
	var editReq req.ClothStyleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := sh.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 款式删除
func (sh styleHandler) del(c *gin.Context) {
	var idReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &idReq)) {
		return
	}
	err := sh.srv.Del(idReq.ID)
	response.CheckAndResp(c, err)
}

// change 款式状态
func (sh styleHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := sh.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)

}
