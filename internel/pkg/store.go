package pkg

import (
	"github.com/go-redis/redis"
	. "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"strconv"
)

var (
	ServerPort     = 8700
	RedisAddr      = "localhost:6379"
	RedisPassWord  = ""
	RedisDB        = 0
	DriverName     = "mysql"
	DataSourceName = "root:123456@tcp(127.0.0.1:3306)/ok-short?charset=utf8&parseTime=true&loc=Local"

	LinkPrefix = "http://sc.vaiwan.com/"

	ReCli = &redis.Client{} // Redis存储
	MyDB  = &DB{}           // MySQL存储
)

func init() {
	portStr := os.Getenv("APP_PORT")
	if portStr != "" {
		ServerPort, _ = strconv.Atoi(portStr)
	}

	// Redis
	redisAddr := os.Getenv("APP_REDIS_ADDR")
	if redisAddr != "" {
		RedisAddr = redisAddr
	}
	redisPw := os.Getenv("APP_REDIS_PASSWORD")
	if redisPw != "" {
		RedisPassWord = redisPw
	}
	redisDb := os.Getenv("APP_REDIS_DB")
	if redisDb != "" {
		db, err := strconv.Atoi(redisDb)
		if err != nil {
			log.Fatal(err)
		}
		RedisDB = db
	}

	// MySQL
	driverName := os.Getenv("APP_MYSQL_DRIVER_NAME")
	if driverName != "" {
		DriverName = driverName
	}
	dataSourceName := os.Getenv("APP_MYSQL_DSN")
	if dataSourceName != "" {
		DataSourceName = dataSourceName
	}

	siteURI := os.Getenv("APP_LINK_PREFIX")
	if siteURI != "" {
		LinkPrefix = siteURI
	}

	log.Printf("LinkPrefix: %s\n", LinkPrefix)
	log.Printf("connect to redis [addr: %s password: %s db: %d]\n", RedisAddr, RedisPassWord, RedisDB)
	log.Printf("connect to mysql [%s]\n", DataSourceName)

	ReCli = redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassWord,
		DB:       RedisDB,
	})
	if _, err := ReCli.Ping().Result(); err != nil {
		panic(err)
	}

	db, err := Open(DriverName, DataSourceName)
	if err != nil {
		panic("failed to connect database")
	}
	log.Printf("MySQL Connected : %+v", db.DB().Stats())
	db.LogMode(true)
	MyDB = db
}

const (
	// URLIdKey redis自增主键所用的key
	URLIdKey = "next.url.id"
	// LinkKey 用于保存短链与原始链接的映射
	LinkKey = "link:%s:url"

	AdminTokenKey = "admin:%s:token"
	AdminInfoKey  = "admin:%s:info"
)