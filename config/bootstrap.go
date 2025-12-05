package gdconfig

import "sync"

var (
	instance *BaseConfig
	onceLoad sync.Once
)

type BaseConfig struct {
	env      Env
	server   Server
	database Database
	cache    Cache
	timer    Time
	queue    Queue
}

func Init() BaseConfig {
	return BaseConfig{}
}

func Load(f func(c *BaseConfig)) {
	onceLoad.Do(func() {
		f(instance)
	})
}

type Server struct {
	Address string
	Timeout int
}

type Env struct {
	Mode string
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

func (c *BaseConfig) GetEnv() Env           { return c.env }
func (c *BaseConfig) GetServer() Server     { return c.server }
func (c *BaseConfig) GetDatabase() Database { return c.database }
func (c *BaseConfig) GetCache() Cache       { return c.cache }
func (c *BaseConfig) GetTime() Time         { return c.timer }
func (c *BaseConfig) GetQueue() Queue       { return c.queue }
