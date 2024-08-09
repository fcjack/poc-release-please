package config

import (
	"testing"

	"github.com/matryer/is"
)

func TestConfigParse(t *testing.T) {
	is := is.New(t)

	c := Config{
		Env:              "dev",
		DatabaseHost:     "localhost",
		DatabasePort:     "3306",
		DatabaseUser:     "user",
		DatabasePassword: "password",
		DatabaseName:     "database",
		DatabaseCryptKey: "cryptkey",
	}

	err := c.Parse()
	is.NoErr(err)

	is.Equal(c.Env, "dev")
	is.Equal(c.DatabaseHost, "localhost")
	is.Equal(c.DatabasePort, "3306")
	is.Equal(c.DatabaseUser, "user")
	is.Equal(c.DatabasePassword, "password")
	is.Equal(c.DatabaseName, "database")
	is.Equal(c.DatabaseCryptKey, "cryptkey")
}
