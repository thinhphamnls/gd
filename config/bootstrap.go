package gdconfig

const ProductionEnv = "production"

type IConfig interface {
	GetServer() Server
	GetDatabase() Database
	GetCache() Cache
	GetTime() Time
	GetQueue() Queue
}

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

func (c Config) GetServer() Server {
	return c.Server
}

func (c Config) GetDatabase() Database {
	return c.Database
}

func (c Config) GetCache() Cache {
	return c.Cache
}

func (c Config) GetTime() Time {
	return c.Timer
}

func (c Config) GetQueue() Queue {
	return c.Queue
}
