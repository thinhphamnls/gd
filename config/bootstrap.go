package gdconfig

type Env struct {
	Mode string
}

type Server struct {
	Address string
	Timeout int
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

type Redis struct {
	Host         string
	Port         string
	Password     string
	Db           int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

type Kafka struct {
	Brokers []string
	Topic   string
}

type Time struct {
	Zone string
}
