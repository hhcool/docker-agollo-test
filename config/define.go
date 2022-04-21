package config

import "time"

const DefaultConfigFile = "config.yaml"
const DefaultDevConfigFile = "config.dev.yaml"
const DefaultApolloNamespace = "e-msr-file.yml"

type Config struct {
	Serve struct {
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		Env      string `yaml:"env"`
		Client   string `yaml:"client"`
		MainDate string `yaml:"mainDate"`
		MainUrl  string `yaml:"mainUrl"`
	} `yaml:"serve"`
	Middleware struct {
		Cors      bool          `yaml:"cors"`
		Timeout   time.Duration `yaml:"timeout"`
		RateLimit int64         `yaml:"rateLimit"`
		SizeLimit int64         `yaml:"sizeLimit"`
		Whitelist []string      `yaml:"whitelist"`
	} `yaml:"middleware"`
	File struct {
		Type      string `yaml:"type"`
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		Endpoint  string `yaml:"endpoint"`
		Bucket    string `yaml:"bucket"`
		Path      string `yaml:"path"`
		ExportExp string `yaml:"exportExp"`
	} `yaml:"file"`
	Redis struct {
		Host     []string `yaml:"host"`
		Password string   `yaml:"password"`
	} `yaml:"redis"`
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		DataBase string `yaml:"database"`
		Config   string `yaml:"config"`
	} `yaml:"mysql"`
	Logger struct {
		LenBody     int    `yaml:"lenbody"`
		LenResponse int    `yaml:"lenresponse"`
		GoutDebug   bool   `yaml:"goutdebug"`
		Path        string `yaml:"path"`
		MaxAge      int    `yaml:"maxage"`
	} `yaml:"logger"`
}
