package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

var (
	RUN_MODE string
	SERVER   struct {
		HTTP_PORT     int
		READ_TIMEOUT  time.Duration
		WRITE_TIMEOUT time.Duration
	}
	DATABASE struct {
		URI string
	}
	PrivateKey *rsa.PrivateKey
)

func init() {
	c, err := ini.Load("config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	RUN_MODE = c.Section("").Key("run_mode").In("dev", []string{"dev", "prod"})
	{
		SERVER.HTTP_PORT = c.Section("server").Key("http_port").MustInt(65535)
		SERVER.READ_TIMEOUT = time.Duration(c.Section("server").Key("read_timeout").MustInt(60)) * time.Second
		SERVER.WRITE_TIMEOUT = time.Duration(c.Section("server").Key("write_timeout").MustInt(60)) * time.Second
	}
	if databaseuri := os.Getenv("DATABASE_URI"); databaseuri != "" {
		DATABASE.URI = databaseuri
	} else {
		DATABASE.URI = c.Section("database").Key("uri").String()
	}

	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		os.Exit(1)
	}
}
