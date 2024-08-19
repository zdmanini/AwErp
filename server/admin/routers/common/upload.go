package common

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/service/common"
	"Awesome/config"
	"Awesome/core"
	"Awesome/core/response"
	"Awesome/middleware"
	"Awesome/util"
	"github.com/gin-gonic/gin"
)

var UploadGroup = core.Group("/common", newUploadHandler, regUpload, middleware.TokenAuth())

func newUploadHandler(srv common.IUploadService) *uploadHandler {
	return &uploadHandler{srv: srv}
}

func regUpload(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *uploadHandler) {
		rg.POST("/upload/image", middleware.RecordLog("上传图片", middleware.RequestFile), handle.uploadImage)
		rg.POST("/upload/video", middleware.RecordLog("上传视频", middleware.RequestFile), handle.uploadVideo)
	})
}

type uploadHandler struct {
	srv common.IUploadService
}

// uploadImage 上传图片
func (uh uploadHandler) uploadImage(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := uh.srv.UploadImage(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}

// uploadVideo 上传视频
func (uh uploadHandler) uploadVideo(c *gin.Context) {
	var uReq req.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := uh.srv.UploadVideo(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}
