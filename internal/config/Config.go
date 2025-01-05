package config

type Config struct {
	Title    string
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schemas  []string
}
