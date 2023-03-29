package service

import (
	"github.com/beijJxy/knowledge_repository/global"
	"github.com/beijJxy/knowledge_repository/model"
)

type Knowledge_repositoryService struct{}

func (e *Knowledge_repositoryService) PlugService(req model.Request) (res model.Response, err error) {
	// 写你的业务逻辑
	return res, nil
}

// GetAllVulnRepoService 获取所有漏洞库
func (knowService *Knowledge_repositoryService) GetAllVulnRepoService() (vulns []model.VulnRepo, err error) {
	err = global.VULN_DB.Find(&vulns).Error
	return
}
