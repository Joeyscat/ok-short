package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	RunMode         string
	HttpPort        string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	LinkPrefix      string
	RequestLimit    time.Duration
	AllowOrigins    []string
}

type MongoDBSettingS struct {
	Addr     []string
	User     string
	Password string
	AuthDB   string
}

type RedisSettingS struct {
	Addr     string
	Password string
	DB       int
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
