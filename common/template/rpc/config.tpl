package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf struct {
        Host     string
        Port     int
        User     string
        Password string
        DBName   string
    }
    RedisConf struct {
        Host     string
        Port     int
        Password string
        Index    int
    }
}
