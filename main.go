package knowledge_repository

import (
	"github.com/beijJxy/knowledge_repository/global"
	"github.com/beijJxy/knowledge_repository/router"
	"github.com/gin-gonic/gin"
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

// 本地测试打开，同时package为main包
//func main() {
//	initialize.Viper()
//	global.VULN_DB = initialize.Gorm()
//	r := gin.Default()
//	// 访问/version的返回值会随配置文件的变化而变化
//	r.GET("/vuln/auto", func(c *gin.Context) {
//		// 自动迁移
//		global.VULN_DB.AutoMigrate(&model.VulnRepo{})
//		m := global.GVA_CONFIG
//		c.String(http.StatusOK, m.Knowledge+"自动迁移完成")
//	})
//
//	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
//	r.GET("/vuln/list", plugApi.GetAllVulnRepoApis)
//	r.Run(":9999")
//}
