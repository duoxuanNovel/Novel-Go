package svc

import (
	"ddxs-api/internal/config"
	"ddxs-api/internal/zs"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Mysql.TablePrefix,
			SingularTable: false,
			NoLowerCase:   false,
		},
	})
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password, // 没有密码，默认值
		DB:       c.Redis.Db,       // 默认DB 0
	})
	//项目启动时，自动自动创建索引到z
	zsObj := zs.NewZsConfig(c)
	zsObj.SyncGORMToIndexFull(db)
	//在项目启动就将数据同步到索引 然后就定时每6小时同步一次
	ticker := time.NewTicker(6 * time.Hour)
	go func() {
		for range ticker.C {
			zsObj.SyncGORMToIndexProtected(db)
		}
	}()

	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  rdb,
	}
}
