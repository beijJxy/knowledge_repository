package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"knowledge_repository/model"
	"knowledge_repository/service"
)

type KnowledgeRepositoryApi struct{}

// @Tags Knowledge_repository
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /knowledge/routerName [post]
func (p *KnowledgeRepositoryApi) ApiName(c *gin.Context) {
	var plug model.Request
	_ = c.ShouldBindJSON(&plug)
	if res, err := service.ServiceGroupApp.PlugService(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(res, "成功", c)
	}
}

// GetAllVulnRepoApis 获取所有漏洞信息
func (s *KnowledgeRepositoryApi) GetAllVulnRepoApis(c *gin.Context) {
	var plug model.Request
	_ = c.ShouldBindJSON(&plug)
	if res, err := service.ServiceGroupApp.GetAllVulnRepoService(); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(res, "成功", c)
	}
}
