package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"knowledge_repository/config"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Knowledge_repository
	VULN_DB    *gorm.DB
)

var GlobalConfig = new(config.Knowledge_repository)
