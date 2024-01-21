package config

import (
	"strconv"
	"strings"
)

type Config struct {
	env map[string]string
}

func Init(env []string) *Config {

	mappedEnv := make(map[string]string)
	for _, t := range env {
		key, value, _ := strings.Cut(t, "=")
		mappedEnv[key] = value
	}

	return &Config{mappedEnv}
}

func (c Config) GetListeningPort() (int, error) {
	port := c.env["LISTENING_PORT"]
	portInt, _ := strconv.Atoi(port)
	return portInt, nil
}

func (c Config) GetDatabaseHost() (string, error) {
	host := c.env["DATABASE_HOST"]
	return host, nil
}

func (c Config) GetDatabaseUser() (string, error) {
	user := c.env["DATABASE_USER"]
	return user, nil
}

func (c Config) GetDatabasePassword() (string, error) {
	password := c.env["DATABASE_PASSWORD"]
	return password, nil
}
