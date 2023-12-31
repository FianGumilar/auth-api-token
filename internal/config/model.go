package config

type Config struct {
	DB  Database
	Srv Server
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Server struct {
	Host string
	Port string
}
