package config

import "github.com/spf13/viper"

type Config struct {
	Channels []string
	HTTP     HTTPConfig
	Mongo    MongoConfig
}

type HTTPConfig struct {
	Addr        string
	SearchRoute string
}

type MongoConfig struct {
	Addr string
}

func Load(path string) Config {
	var conf Config
	viper.SetConfigType("yaml")
	if path != "" {
		viper.SetConfigFile(path)
		viper.ReadInConfig()
	}

	viper.SetDefault("http.addr", ":14000")
	viper.SetDefault("http.search_route", "/messages/search")
	viper.SetDefault("mongo.addr", "mongodb://localhost:27017")

	conf.Channels = viper.GetStringSlice("channels")
	conf.HTTP.Addr = viper.GetString("http.addr")
	conf.HTTP.SearchRoute = viper.GetString("http.search_route")
	conf.Mongo.Addr = viper.GetString("mongo.addr")
	return conf
}
