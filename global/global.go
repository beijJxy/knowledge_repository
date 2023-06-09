package global

import (
	"github.com/beijJxy/knowledge_repository/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.KnowledgeRepository
	VULN_DB    *gorm.DB
)

var GlobalConfig = new(config.KnowledgeRepository)
