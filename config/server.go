package config

type ServerConfig struct {
	port int `toml:"port"`
}

func (server *ServerConfig) GetPort() int {
	return server.port
}
