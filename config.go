package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var config struct {
	ServerUrl string
	Token     string
	Retries   int
	Workers   int
	StatsInterval time.Duration
	Verbose   bool
}

const (
	ConfServerUrl = "server_url"
	ConfToken     = "token"
	ConfRetries   = "retries"
	ConfWorkers   = "workers"
	ConfStatsInterval = "stats_interval"
	ConfVerbose   = "verbose"
)

func prepareConfig() {
	viper.SetDefault(ConfRetries, 3)
	viper.SetDefault(ConfWorkers, 50)
	viper.SetDefault(ConfStatsInterval, 3 * time.Second)
	viper.SetDefault(ConfVerbose, false)
}

func readConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	config.ServerUrl = viper.GetString(ConfServerUrl)
	if config.ServerUrl == "" {
		logrus.Fatal("config: server_url not set!")
	}

	config.Token = viper.GetString(ConfToken)
	if config.Token == "" {
		logrus.Fatal("config: token not set!")
	}

	config.Retries = viper.GetInt(ConfRetries)
	if config.Retries < 0 {
		config.Retries = 1 << 31
	}

	config.Workers = viper.GetInt(ConfWorkers)
	if config.Workers <= 0 {
		logrus.Fatal("config: illegal value %d for workers!", config.Workers)
	}

	config.StatsInterval = viper.GetDuration(ConfStatsInterval)

	config.Verbose = viper.GetBool(ConfVerbose)
	if config.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}