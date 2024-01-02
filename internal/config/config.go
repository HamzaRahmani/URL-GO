package config

import (
	"strconv"
	"strings"
)

type config struct {
	env map[string]string
}

func Init(env []string) *config {

	mappedEnv := make(map[string]string)
	for _, t := range env {
		key, value, _ := strings.Cut(t, "=")
		mappedEnv[key] = value
	}

	return &config{mappedEnv}
}

func (c config) GetListeningPort() (int, error) {
	port := c.env["LISTENING_PORT"]
	portInt, _ := strconv.Atoi(port)
	return portInt, nil
}

func (c config) GetDatabaseHost() (string, error) {
	host := c.env["DATABASE_HOST"]
	return host, nil
}

func (c config) GetDatabaseUser() (string, error) {
	host := c.env["DATABASE_USER"]
	return host, nil
}

func (c config) GetDatabasePassword() (string, error) {
	password := c.env["DATABASE_PASSWORD"]
	return password, nil
}
