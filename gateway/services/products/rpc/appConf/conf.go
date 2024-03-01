package appConf

type App struct {
	Postgres Postgres
	Cache    Cache
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

type Cache struct {
	Host string
	Pass string
}
