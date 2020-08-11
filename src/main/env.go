package main

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	S    Storage
	port int
}

const (
	RedisAddr      = "localhost:6379"
	RedisDB        = "0"
	DriverName     = "mysql"
	DataSourceName = "root:123456@/ok-short?charset=utf8&parseTime=true"
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
	dataSourceName := os.Getenv("APP_MYSQL_DATA_SOURCE_NAME")
	if dataSourceName == "" {
		dataSourceName = DataSourceName
	}

	log.Printf("connect to redis [addr: %s password: %s db: %d]\n", addr, password, db)
	log.Printf("connect to mysql [%s]\n", dataSourceName)

	r := NewRedisCli(addr, password, db)
	m := NewMySQL(driverName, dataSourceName)
	return &Env{
		S: &Service{
			R: *r,
			M: *m,
		},
		port: port}
}
