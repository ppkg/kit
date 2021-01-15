package rp_kit

import (
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/limitedlee/microservice/common"
	"github.com/limitedlee/microservice/common/config"
	"github.com/maybgit/glog"
	"github.com/tricobbler/rp-kit/cast"
)

var (
	redisHandle *redis.Client
)

//获取redis集群客户端
//dsn: addr:port|pwd|db
func GetRedisConn(dsn ...string) *redis.Client {
	if redisHandle != nil {
		_, err := redisHandle.Ping().Result()
		if err == nil {
			return redisHandle
		}
	}

	var (
		db   int
		addr string
		pwd  string
	)
	if len(dsn) > 0 {
		dsnSlice := strings.Split(dsn[0], "|")
		addr = dsnSlice[0]
		pwd = dsnSlice[1]
		db = cast.ToInt(dsnSlice[2])
	} else {
		db = cast.ToInt(config.GetString("redis.DB"))
		addr = config.GetString("redis.Addr")
		pwd = config.GetString("redis.Password")
	}

	redisHandle = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           db,
		MinIdleConns: 20,
		IdleTimeout:  60,
		PoolSize:     200,
	})
	_, err := redisHandle.Ping().Result()
	if err != nil {
		glog.Error("redis连接错误，", err, "，PoolStats：", JsonEncode(redisHandle.PoolStats()))
		panic(err)
	}

	return redisHandle
}

type DBEngine interface {
	Ping() error
}

func DBEngineCheck(engine DBEngine, newEngine func(...string) DBEngine, maxRetryTimes, interval int) {
	defer CatchPanic()

	for {
		time.Sleep(time.Duration(interval) * time.Second)

		retryTimes := 0
	reconnect:
		if err := engine.Ping(); err != nil {
			if retryTimes < maxRetryTimes {
				DBEngineReset(engine, newEngine)
				retryTimes++
				log.Printf("数据库重连, try %v...", retryTimes)
				goto reconnect
			}

			Alert("【" + common.PbConfig.Grpc.Appid + "】数据库连接被关闭未能恢复，" + err.Error())
		}
	}
}

func DBEngineReset(engine DBEngine, newEngine func(...string) DBEngine) {
	engine = nil
	engine = newEngine()
}
