package service

import (
	"errors"
	"github.com/beijJxy/knowledge_repository/global"
	"github.com/beijJxy/knowledge_repository/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type KnowledgeRepositoryService struct{}

func (e *KnowledgeRepositoryService) PlugService(req model.Request) (res model.Response, err error) {
	// 写你的业务逻辑
	return res, nil
}

// GetAllVulnRepoService 获取所有漏洞库
func (knowService *KnowledgeRepositoryService) GetAllVulnRepoService() (vulns []model.VulnRepo, err error) {
	err = global.VULN_DB.Find(&vulns).Error
	return
}

// CreateVulnRepo 新增漏洞库
func (k *KnowledgeRepositoryService) CreateVulnRepo(vuln model.VulnRepo) (err error) {
	return global.VULN_DB.Create(&vuln).Error
}

// UpdateVuln 更新漏洞库
func (k *KnowledgeRepositoryService) UpdateVuln(vuln model.VulnRepo) (err error) {
	// 原数据
	var oldV model.VulnRepo
	err = global.VULN_DB.Where("id = ?", vuln.ID).First(&oldV).Error
	if oldV.CveNumber != vuln.CveNumber {
		if !errors.Is(global.VULN_DB.Where("title = ? AND cve_number = ?", vuln.Title, vuln.CveNumber).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同漏洞")
		}
	}
	if err != nil {
		return err
	} else {
		// 更新选定字段
		err = global.VULN_DB.Model(&model.VulnRepo{}).Where("id = ?", vuln.ID).Select("title").Updates(vuln).Error
		// 更新所有
		//err = global.VULN_DB.Save(&vuln).Error
	}
	return err
}

func (k *KnowledgeRepositoryService) DeleteVulnRepo(vuln model.VulnRepo) (err error) {
	var entity model.VulnRepo
	err = global.VULN_DB.Where("id = ?", vuln.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                        // api记录不存在
		return err
	}
	err = global.VULN_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}
