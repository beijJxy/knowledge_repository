package router

import (
	"github.com/beijJxy/knowledge_repository/api"
	"github.com/gin-gonic/gin"
)

type Knowledge_repositoryRouter struct {
}

func (s *Knowledge_repositoryRouter) InitKnowledge_repositoryRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
	{
		plugRouter.GET("routerName", plugApi.ApiName)
		plugRouter.GET("vuln/list", plugApi.GetAllVulnRepoApis)
	}
}
