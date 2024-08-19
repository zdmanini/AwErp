package system

import (
	"Awesome/admin/schemas/req"
	"Awesome/admin/service/system"
	"Awesome/core"
	"Awesome/core/request"
	"Awesome/core/response"
	"Awesome/middleware"
	"Awesome/util"
	"github.com/gin-gonic/gin"
)

var LogGroup = core.Group("/system", newLogHandler, regLog, middleware.TokenAuth())

func newLogHandler(srv system.ISystemLogsServer) *logHandler {
	return &logHandler{srv: srv}
}

func regLog(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *logHandler) {
		rg.GET("/log/operate", handle.operate)
		rg.GET("/log/login", handle.login)
	})
}

type logHandler struct {
	srv system.ISystemLogsServer
}

// operate 操作日志
func (lh logHandler) operate(c *gin.Context) {
	var page request.PageReq
	var logReq req.SystemLogOperateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := lh.srv.Operate(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}

// login 登录日志
func (lh logHandler) login(c *gin.Context) {
	var page request.PageReq
	var logReq req.SystemLogLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := lh.srv.Login(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}
