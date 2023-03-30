package config

type KnowledgeRepository struct {
	Knowledge string `mapstructure:"knowledge" json:"knowledge" yaml:"knowledge"` // 知识库类型
	Mysql     Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`             //数据库配置
	Type      string `mapstructure:"type" json:"type" yaml:"type"`                //类型
}
