package kit

import (
	"strings"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gogrpc/glog"
	"github.com/gogrpc/kit/cast"
	"github.com/limitedlee/microservice/common"
)

type DBEngine struct {
	Engine interface{}
	dsn    string
}

func NewDBEngine(dsn string) *DBEngine {
	e := &DBEngine{
		dsn: dsn,
	}
	e.NewXormEngine()
	return e
}

func NewRedisEngine(dsn string) *DBEngine {
	e := &DBEngine{
		dsn: dsn,
	}
	e.NewRedisConn()
	return e
}

func (e *DBEngine) DBEngineCheck(f func() interface{}, maxRetryTimes, interval int) {
	defer CatchPanic()

	for {
		time.Sleep(time.Duration(interval) * time.Second)

		retryTimes := 0
	reconnect:
		var (
			err  error
			desc string
		)
		switch e.Engine.(type) {
		case *xorm.Engine:
			desc = "mysql"
			err = e.Engine.(*xorm.Engine).DB().Ping()
		case *redis.Client:
			desc = "redis"
			err = e.Engine.(*redis.Client).Ping().Err()
		}

		if err != nil {
			if retryTimes < maxRetryTimes {
				retryTimes++
				glog.Infof(desc+"断开重连, try %v...", retryTimes)
				time.Sleep(time.Duration(retryTimes) * time.Second)
				e.Engine = e.ResetEngine(f)

				goto reconnect
			}

			Alert("【" + common.PbConfig.Grpc.Appid + "】" + desc + "连接被关闭未能恢复，" + err.Error())
		}
	}
}

func (e *DBEngine) ResetEngine(f func() interface{}) interface{} {
	return f()
}

//获取redis客户端
func (e *DBEngine) NewRedisConn() *redis.Client {
	dsnSlice := strings.Split(e.dsn, "|")
	if len(dsnSlice) < 3 {
		glog.Error("redis配置不正确，", e.dsn)
		panic("redis配置不正确")
	}

	redisHandle := redis.NewClient(&redis.Options{
		Addr:         dsnSlice[0],
		Password:     dsnSlice[1],
		DB:           cast.ToInt(dsnSlice[2]),
		MinIdleConns: 10,
		IdleTimeout:  60,
		PoolSize:     200,
	})
	_, err := redisHandle.Ping().Result()
	if err != nil {
		glog.Error("redis连接错误，", err, "，PoolStats：", JsonEncode(redisHandle.PoolStats()))
		panic(err)
	}

	e.Engine = redisHandle
	return redisHandle
}

func (e *DBEngine) NewRedisConnInterface() interface{} {
	return e.NewRedisConn()
}

//获取数据库客户端
func (e *DBEngine) NewXormEngine() *xorm.Engine {
	if len(e.dsn) == 0 {
		glog.Error("mysql配置不正确，", e.dsn)
		panic("mysql配置不正确")
	}

	xormEngine, err := xorm.NewEngine("mysql", e.dsn)
	if err != nil {
		glog.Error("mysql connect fail", err)
		panic(err)
	}

	if IsDebug {
		xormEngine.ShowSQL()
		xormEngine.ShowExecTime()
	}

	//空闲关闭时间
	xormEngine.SetConnMaxLifetime(60 * time.Second)
	//最大空闲连接
	xormEngine.SetMaxIdleConns(10)
	//最大连接数
	xormEngine.SetMaxOpenConns(500)

	// 设置时区
	xormEngine.SetTZLocation(time.Local)

	e.Engine = xormEngine
	return xormEngine
}

func (e *DBEngine) NewXormEngineInterface() interface{} {
	return e.NewXormEngine()
}
