package api

import (
	"github.com/beijJxy/knowledge_repository/model"
	"github.com/beijJxy/knowledge_repository/service"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
func (k *KnowledgeRepositoryApi) GetAllVulnRepoApis(c *gin.Context) {
	var plug model.Request
	_ = c.ShouldBindJSON(&plug)
	if res, err := service.ServiceGroupApp.GetAllVulnRepoService(); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
		response.OkWithDetailed(res, "成功", c)
	}
}

func (k *KnowledgeRepositoryApi) CreateVulnRepo(c *gin.Context) {
	var vuln model.VulnRepo
	err := c.ShouldBind(&vuln)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	// 省略校验
	err = service.ServiceGroupApp.CreateVulnRepo(vuln)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("漏洞库数据创建成功", c)
}

func (k *KnowledgeRepositoryApi) UpdateVulnRepo(c *gin.Context) {
	var vuln model.VulnRepo
	err := c.ShouldBindJSON(&vuln)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 省略校验
	err = service.ServiceGroupApp.UpdateVuln(vuln)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (k *KnowledgeRepositoryApi) DeleteVulnRepo(c *gin.Context) {
	var entity model.VulnRepo
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 省略校验
	err = service.ServiceGroupApp.DeleteVulnRepo(entity)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)

}
