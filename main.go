package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"net/http"
	"time"

	"github.com/bline/gotime-server/application"
	"os"
	"path"
)

func init() {

}

func newConfig() (*viper.Viper, error) {
	defaultDSN := "root:@tcp(localhost:3306)/gotime_server?parseTime=true"
	home := os.Getenv("HOME")
	env := os.Getenv("GOTIME_ENV")
	assetsRoot := path.Join(home, ".config", "gotime")
	confDirLocal := path.Join(assetsRoot, "config")
	confDirSystem := "/etc/gotime"

	c := viper.New()
	c.SetDefault("dsn", defaultDSN)
	c.SetDefault("cookie_secret1", "uk7DHr4pGJ2Ra9fQ7zKY8IP0BKv7vVse")
	c.SetDefault("cookie_secret2", "xGRfc7Del998h97yY0ad3H9N2y4Q13M8")
	c.SetDefault("google_client_id", "33812767661-4a1p5lotkkveeodjehfpkucvmbpkmkhf.apps.googleusercontent.com")
	c.SetDefault("http_addr", ":8888")
	c.SetDefault("http_cert_file", "")
	c.SetDefault("http_key_file", "")
	c.SetDefault("http_drain_interval", "1s")

	c.SetConfigType("yaml")
	c.SetConfigName(env)
	c.AddConfigPath(confDirLocal)
	c.AddConfigPath(confDirSystem)

	c.SetEnvPrefix("GOTIME")
	c.AutomaticEnv()

	return c, nil
}

func main() {
	config, err := newConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	app, err := application.New(config)
	if err != nil {
		logrus.Fatal(err)
	}

	middle, err := app.MiddlewareStruct()
	if err != nil {
		logrus.Fatal(err)
	}

	serverAddress := config.Get("http_addr").(string)

	certFile := config.Get("http_cert_file").(string)
	keyFile := config.Get("http_key_file").(string)
	drainIntervalString := config.Get("http_drain_interval").(string)

	drainInterval, err := time.ParseDuration(drainIntervalString)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := &graceful.Server{
		Timeout: drainInterval,
		Server:  &http.Server{Addr: serverAddress, Handler: middle},
	}

	logrus.Infoln("Running HTTP server on " + serverAddress)

	if certFile != "" && keyFile != "" {
		err = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
