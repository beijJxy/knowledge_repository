package router

import (
	"github.com/gin-gonic/gin"
	"knowledge_repository/api"
)

type Knowledge_repositoryRouter struct {
}

func (s *Knowledge_repositoryRouter) InitKnowledge_repositoryRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.KnowledgeRepositoryApi
	{
		plugRouter.POST("routerName", plugApi.ApiName)
	}
}
