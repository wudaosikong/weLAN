package datasource

import (
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"weLAN/server/config"
	"github.com/kataras/iris/v12"
)

/**
 * 返回Redis实例
 */
func NewRedis() *redis.Database {

	var database *redis.Database

	//项目配置
	cmsConfig := config.InitConfig()
	if cmsConfig != nil {
		iris.New().Logger().Info(" hello ")
		rd := cmsConfig.Redis
		iris.New().Logger().Info(rd)
		database = redis.New(redis.Config{
			Network:     rd.NetWork,
			Addr:        rd.Addr + ":" + rd.Port,
			Password:    rd.Password,
			Database:    "",
			MaxActive:   10,
			Prefix:      rd.Prefix,
		})
	} else {
		iris.New().Logger().Info(" hello  error ")
	}
	return database
}
