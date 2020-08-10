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

func getEnv() *Env {
	portStr := os.Getenv("APP_PORT")
	port := 8700
	if portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	passwd := os.Getenv("APP_REDIS_PASSWD")

	dbS := os.Getenv("APP_REDIS_DB")
	if dbS == "" {
		dbS = "0"
	}

	db, err := strconv.Atoi(dbS)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to redis [addr: %s password: %s db: %d]", addr, passwd, db)

	r := NewRedisCli(addr, passwd, db)
	return &Env{S: r, port: port}
}
