package svc

import (
    "xorm.io/xorm"
	{{.configImport}}
	x_orm "lightleap-c/common/xorm"
	"github.com/go-redis/redis/v8"
	"lightleap-c/common/xredis"
)

type ServiceContext struct {
	Config {{.config}}
	DB *xorm.Engine
	Rds *redis.Client
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB: x_orm.InitDB(c.MysqlConf.Host, c.MysqlConf.Port, c.MysqlConf.User, c.MysqlConf.Password, c.MysqlConf.DBName),
		Rds: xredis.InitRedis(c.RedisConf.Host, c.RedisConf.Port, c.RedisConf.Password, c.RedisConf.Index),
		{{.middlewareAssignment}}
	}
}
