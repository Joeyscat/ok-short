package main

import (
	"github.com/joeyscat/ok-short/admin"
	"github.com/joeyscat/ok-short/api"
	. "github.com/joeyscat/ok-short/store"
	"log"
	"os"
	"strconv"
)

type Env struct {
	API   api.Service
	ADMIN admin.Service
	port  int
}

const (
	RedisAddr      = "localhost:6379"
	RedisDB        = "0"
	DriverName     = "mysql"
	DataSourceName = "root:123456@tcp(127.0.0.1:3306)/ok-short?charset=utf8&parseTime=true&loc=Local"
)

var (
	LinkPrefix = "http://localhost:8700/"
)

func getEnv() *Env {
	portStr := os.Getenv("APP_PORT")
	port := 8700
	if portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	// Redis
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = RedisAddr
	}
	password := os.Getenv("APP_REDIS_PASSWORD")
	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = RedisDB
	}
	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}
	// MySQL
	driverName := os.Getenv("APP_MYSQL_DRIVER_NAME")
	if driverName == "" {
		driverName = DriverName
	}
	dataSourceName := os.Getenv("APP_MYSQL_DSN")
	if dataSourceName == "" {
		dataSourceName = DataSourceName
	}

	siteURI := os.Getenv("APP_LINK_PREFIX")
	if siteURI != "" {
		LinkPrefix = siteURI
	}

	log.Printf("LinkPrefix: %s\n", LinkPrefix)
	log.Printf("connect to redis [addr: %s password: %s db: %d]\n", addr, password, db)
	log.Printf("connect to mysql [%s]\n", dataSourceName)

	r := NewRedisCli(addr, password, db)
	m := NewMySQL(driverName, dataSourceName)

	ReCli = *r

	return &Env{
		API: &api.LinkService{
			R: *r,
			M: *m,
		},
		ADMIN: admin.Service{
			AuthorService: admin.AuthorService{},
			LinkService: admin.LinkService{
				R: *r,
				M: *m,
			},
			VisitorService: admin.VisitorService{},
		},
		port: port}
}
