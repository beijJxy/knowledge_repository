package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type VulnRepo struct {
	global.GVA_MODEL
	CveNumber string `json:"cve_number" gorm:"cve_number:cve编号"` // cve编号
	Title     string `json:"title" gorm:"title:漏洞标题"`            // api中文描述
}

func (VulnRepo) TableName() string {
	return "vuln_repo"
}
