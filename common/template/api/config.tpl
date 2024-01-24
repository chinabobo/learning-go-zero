package config

import {{.authImport}}

type Config struct {
	rest.RestConf
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
	{{.auth}}
	{{.jwtTrans}}
}
