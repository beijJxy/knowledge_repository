package main

import (
	"github.com/beijJxy/knowledge_repository/api"
	"github.com/beijJxy/knowledge_repository/global"
	"github.com/beijJxy/knowledge_repository/initialize"
	"github.com/beijJxy/knowledge_repository/model"
	"github.com/beijJxy/knowledge_repository/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type KnowledgeRepositoryPlugin struct {
}

func CreateKnowledgeRepositoryPlug(Type string) *KnowledgeRepositoryPlugin {
	// 初始化
	global.GlobalConfig.Type = Type
	initialize.Viper()
	global.VULN_DB = initialize.Gorm()
	return &KnowledgeRepositoryPlugin{}
}

func (*KnowledgeRepositoryPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitKnowledgeRepositoryRouter(group)
}

func (*KnowledgeRepositoryPlugin) RouterPath() string {
	return "knowledge"
}

// 本地测试打开，同时package改为main包
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

	// query
	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
	r.GET("/vuln/list", plugApi.GetAllVulnRepoApis)

	// add
	r.POST("/vuln/create", plugApi.CreateVulnRepo)

	// update
	r.POST("/vuln/update", plugApi.UpdateVulnRepo)

	// delete
	r.POST("/vuln/delete", plugApi.DeleteVulnRepo)

	r.Run(":9999")
}
