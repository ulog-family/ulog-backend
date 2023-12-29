package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"reflect"
	"strings"
	"time"
)

var (
	Server server
	DB     db
	JWT    jwt
)

type server struct {
	Host string
	Port string
}

func (s server) HostPort() string {
	return s.Host + ":" + s.Port
}

type db struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	Sslmode  string
	TimeZone string
}

func (db db) DSN() string {
	var dsn string
	t, v := reflect.TypeOf(db), reflect.ValueOf(db)
	for k := 0; k < t.NumField(); k++ {
		dsn += fmt.Sprintf("%+v = %v ", strings.ToLower(t.Field(k).Name), v.Field(k).Interface())
	}
	return dsn
}

type jwt struct {
	Key         string
	IdentityKey string
	UserIDKey   string
}

type config struct {
	Server *server
	DB     *db
	JWT    *jwt
}

var configData = config{
	Server: &Server,
	DB:     &DB,
	JWT:    &JWT,
}

func init() {
	err := configor.New(&configor.Config{
		AutoReload:         true,
		AutoReloadInterval: time.Minute,
		Verbose:            true,
	}).Load(&configData, "config/config.toml")
	if err != nil {
		panic("配置导入错误")
	}
}
