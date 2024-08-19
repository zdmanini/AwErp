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

var InfoGroup = core.Group("/basic", newInfoHandler, regInfo, middleware.TokenAuth())

func newInfoHandler(srv basic.BasicInfoService) *infoHandler {
	return &infoHandler{srv: srv}
}

func regInfo(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *infoHandler) {
		rg.GET("/info/item", handle.itemList)
		rg.GET("/info/type", handle.typeList)
		rg.GET("/info/query", handle.query)
		rg.POST("/info/add", middleware.RecordLog("字典新增"), handle.add)
		rg.POST("/info/edit", middleware.RecordLog("字典编辑"), handle.edit)
		rg.POST("/info/del", middleware.RecordLog("字典删除"), handle.del)
		rg.POST("/info/addType", middleware.RecordLog("字典新增"), handle.addType)
		rg.POST("/info/editType", middleware.RecordLog("字典编辑"), handle.editType)
		rg.POST("/info/delType", middleware.RecordLog("字典类型删除"), handle.delType)
		rg.POST("/info/change", middleware.RecordLog("字典状态"), handle.change)
	})
}

type infoHandler struct {
	srv basic.BasicInfoService
}

// itemList 字典列表
func (ih infoHandler) itemList(c *gin.Context) {
	var listReq req.BasicInfoItemQueryReq
	var pageReq request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &pageReq)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ih.srv.ItemList(pageReq, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// typeList 字典类型列表
func (ih infoHandler) typeList(c *gin.Context) {
	var listReq req.BasicInfoTypeQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := ih.srv.TypeList(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// query 字典查询
func (ih infoHandler) query(c *gin.Context) {
	var queryReq req.BasicInfoQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &queryReq)) {
		return
	}
	res, err := ih.srv.Query(queryReq)
	response.CheckAndRespWithData(c, res, err)
}

// add 字典新增
func (ih infoHandler) add(c *gin.Context) {
	var addReq req.BasicInfoItemAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := ih.srv.Add(addReq)
	response.CheckAndResp(c, err)
}

// edit 字典编辑
func (ih infoHandler) edit(c *gin.Context) {
	var editReq req.BasicInfoItemEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := ih.srv.Edit(editReq)
	response.CheckAndResp(c, err)
}

// del 字典删除
func (ih infoHandler) del(c *gin.Context) {
	var delReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := ih.srv.Del(delReq.ID)
	response.CheckAndResp(c, err)
}

// addType 字典新增
func (ih infoHandler) addType(c *gin.Context) {
	var addReq req.BasicInfoTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	err := ih.srv.TypeAdd(addReq)
	response.CheckAndResp(c, err)
}

// editType 字典编辑
func (ih infoHandler) editType(c *gin.Context) {
	var editReq req.BasicInfoTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := ih.srv.TypeEdit(editReq)
	response.CheckAndResp(c, err)
}

// delType 字典类型删除
func (ih infoHandler) delType(c *gin.Context) {
	var delReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := ih.srv.TypeDel(delReq.ID)
	response.CheckAndResp(c, err)
}

// change 字典状态
func (ih infoHandler) change(c *gin.Context) {
	var changeReq request.IDReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &changeReq)) {
		return
	}
	err := ih.srv.Change(changeReq.ID)
	response.CheckAndResp(c, err)
}
