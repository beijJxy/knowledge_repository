package service

type ServiceGroup struct {
	KnowledgeRepositoryService
}

// 构体进行实例化，得到的是结构体的地址
var ServiceGroupApp = new(ServiceGroup)
