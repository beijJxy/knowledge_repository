package main

import (
	"github.com/gin-gonic/gin"
	"knowledge_repository/api"
	"knowledge_repository/global"
	"knowledge_repository/initialize"
	"knowledge_repository/model"
	"knowledge_repository/router"
	"net/http"
)

type Knowledge_repositoryPlugin struct {
}

func CreateKnowledge_repositoryPlug(Type string) *Knowledge_repositoryPlugin {
	global.GlobalConfig.Type = Type

	return &Knowledge_repositoryPlugin{}
}

func (*Knowledge_repositoryPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitKnowledge_repositoryRouter(group)
}

func (*Knowledge_repositoryPlugin) RouterPath() string {
	return "knowledge"
}

func main() {
	initialize.Viper()
	global.VULN_DB = initialize.Gorm()
	r := gin.Default()
	// 访问/version的返回值会随配置文件的变化而变化
	r.GET("/vuln/auto", func(c *gin.Context) {
		// 自动迁移
		global.VULN_DB.AutoMigrate(&model.VulnRepo{})
		m := global.GVA_CONFIG
		c.String(http.StatusOK, m.Knowledge+"自动迁移完成")
	})

	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
	r.GET("/vuln/list", plugApi.GetAllVulnRepoApis)
	r.Run(":9999")
}
