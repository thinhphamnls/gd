package gdconfig

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const ProductionEnv = "production"

var (
	cf Config
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	Timer    Time
	Queue    Queue
}

type Server struct {
	Env  Env
	Http Http
}

type Env struct {
	Mode string
}

type Http struct {
	Address string
	Timeout int
}

type Database struct {
	GDMain  DbConfig
	GDSlave DbConfig
}

type DbConfig struct {
	Host       string
	Port       string
	DBName     string
	Username   string
	Password   string
	MaxCon     int
	MaxIdleCon int
}

type Cache struct {
	Redis Redis
}

type Redis struct {
	Host         string
	Port         string
	Password     string
	Db           int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

type Time struct {
	Zone string
}

type Queue struct {
	Brokers []string
	Topic   string
	GroupId string
}

func EnvHttp() {
	cf.Server.Env.Mode = viper.GetString("SERVER_MODE")
	cf.Server.Http.Address = viper.GetString("SERVER_HTTP_ADDRESS")
	cf.Server.Http.Timeout = viper.GetInt("SERVER_HTTP_TIMEOUT")
}

func EnvRedis() {
	cf.Cache.Redis.Host = viper.GetString("CACHE_REDIS_HOST")
	cf.Cache.Redis.Port = viper.GetString("CACHE_REDIS_PORT")
	cf.Cache.Redis.Password = viper.GetString("CACHE_REDIS_PASSWORD")
	cf.Cache.Redis.Db = viper.GetInt("CACHE_REDIS_DB")
	cf.Cache.Redis.DialTimeout = viper.GetInt("CACHE_REDIS_DIAL_TIMEOUT")
	cf.Cache.Redis.ReadTimeout = viper.GetInt("CACHE_REDIS_READ_TIMEOUT")
	cf.Cache.Redis.WriteTimeout = viper.GetInt("CACHE_REDIS_WRITE_TIMEOUT")
}

func EnvDatabase() {
	cf.Database.GDMain.Host = viper.GetString("DATABASE_GD_MAIN_HOST")
	cf.Database.GDMain.Port = viper.GetString("DATABASE_GD_MAIN_PORT")
	cf.Database.GDMain.DBName = viper.GetString("DATABASE_GD_MAIN_DB_NAME")
	cf.Database.GDMain.Username = viper.GetString("DATABASE_GD_MAIN_USERNAME")
	cf.Database.GDMain.Password = viper.GetString("DATABASE_GD_MAIN_PASSWORD")
	cf.Database.GDMain.MaxCon = viper.GetInt("DATABASE_GD_MAIN_MAX_CON")
	cf.Database.GDMain.MaxIdleCon = viper.GetInt("DATABASE_GD_MAIN_MAX_IDLE_CON")

	cf.Database.GDSlave.Host = viper.GetString("DATABASE_GD_SLAVE_HOST")
	cf.Database.GDSlave.Port = viper.GetString("DATABASE_GD_SLAVE_PORT")
	cf.Database.GDSlave.DBName = viper.GetString("DATABASE_GD_SLAVE_DB_NAME")
	cf.Database.GDSlave.Username = viper.GetString("DATABASE_GD_SLAVE_USERNAME")
	cf.Database.GDSlave.Password = viper.GetString("DATABASE_GD_SLAVE_PASSWORD")
	cf.Database.GDSlave.MaxCon = viper.GetInt("DATABASE_GD_SLAVE_MAX_CON")
	cf.Database.GDSlave.MaxIdleCon = viper.GetInt("DATABASE_GD_SLAVE_MAX_IDLE_CON")
}

func EnvTimer() {
	cf.Timer.Zone = viper.GetString("TIME_ZONE")
}

func EnvQueue() {
	cf.Queue.Brokers = strings.Split(viper.GetString("QUEUE_BROKERS"), ",")
	cf.Queue.Topic = viper.GetString("QUEUE_TOPIC")
	cf.Queue.GroupId = viper.GetString("QUEUE_GROUP_ID")
}

func setTimeZone(timeZone string) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Fatalf("error loading timezone, %v", err)
	}
	time.Local = loc
}
