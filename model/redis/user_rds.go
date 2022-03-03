package redis

import (
	redis2 "github.com/go-redis/redis"
	"github.com/stevenyao001/edgeCommon/logger"
	"github.com/stevenyao001/edgeCommon/redis"
	"strconv"
	"time"
)

type UserRds struct {
}

func (u *UserRds) ins() string {
	return "rootcloud_edge"
}

func (u *UserRds) SetUser(uid int, userinfo string) error {
	rds, err := redis.GetRds(u.ins())
	if err != nil {
		logger.ErrorLog("UserRds-SetUser", "获取连接报错", "", err)
		return err
	}

	if err := rds.Set(strconv.Itoa(uid), userinfo, time.Second*10).Err(); err != nil {
		logger.ErrorLog("UserRds-SetUser", "设置用户缓存报错", "", err)
		return err
	}

	return nil

}

func (u *UserRds) GetUser(uid int) (string, error) {
	rds, err := redis.GetRds(u.ins())
	if err != nil {
		logger.ErrorLog("UserRds-GetUser", "获取连接报错", "", err)
		return "", err
	}

	res, err := rds.Get(strconv.Itoa(uid)).Result()
	if err != nil && err != redis2.Nil {
		logger.ErrorLog("UserRds-GetUser", "获取用户报错", "", err)

	}
	return res, err
}
