package router

import (
	"github.com/beijJxy/knowledge_repository/api"
	"github.com/gin-gonic/gin"
)

type KnowledgeRepositoryRouter struct {
}

func (s *KnowledgeRepositoryRouter) InitKnowledgeRepositoryRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
	{
		plugRouter.GET("routerName", plugApi.ApiName)
		plugRouter.GET("vuln/list", plugApi.GetAllVulnRepoApis)
		plugRouter.POST("vuln/create", plugApi.CreateVulnRepo)
		plugRouter.POST("vuln/update", plugApi.UpdateVulnRepo)
		plugRouter.POST("vuln/delete", plugApi.DeleteVulnRepo)
	}
}
