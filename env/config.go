package env

import "errors"

var Env config

type config struct {
	HTTPHeader              string `mapstructure:"SCRAPELESS_HTTP_HEADER"`
	ProxyCountry            string `mapstructure:"SCRAPELESS_PROXY_COUNTRY"`
	ProxySessionDurationMax int64  `mapstructure:"SCRAPELESS_PROXY_SESSION_DURATION_MAX"`
	ProxyGatewayHost        string `mapstructure:"SCRAPELESS_PROXY_GATEWAY_HOST"`

	ScrapelessBaseApiUrl string `mapstructure:"SCRAPELESS_BASE_API_URL"`
	ScrapelessStorageUrl string `mapstructure:"SCRAPELESS_STORAGE_API_URL"`
	ScrapelessActorUrl   string `mapstructure:"SCRAPELESS_ACTOR_API_URL"`
	ScrapelessBrowserUrl string `mapstructure:"SCRAPELESS_BROWSER_API_URL"`

	//ScrapingBrowserUrl string `mapstructure:"SCRAPELESS_BROWSER_URL"`
	//ScrapingBrowserApiHost  string `mapstructure:"SCRAPELESS_BROWSER_API_HOST"`
	//ScrapelessApiHost     string `mapstructure:"SCRAPELESS_API_HOST"`
	//ScrapelessCaptchaHost string `mapstructure:"SCRAPELESS_CAPTCHA_HOST"`

	Actor actorEnv `mapstructure:",squash"`
	Log   logEnv   `mapstructure:",squash"`
}

type actorEnv struct {
	TeamId  string `mapstructure:"SCRAPELESS_TEAM_ID"`
	ActorId string `mapstructure:"SCRAPELESS_ACTOR_ID"`
	RunId   string `mapstructure:"SCRAPELESS_RUN_ID"`
	ApiKey  string `mapstructure:"SCRAPELESS_API_KEY"`

	KvNamespaceId string `mapstructure:"SCRAPELESS_KV_NAMESPACE_ID"`
	DatasetId     string `mapstructure:"SCRAPELESS_DATASET_ID"`
	BucketId      string `mapstructure:"SCRAPELESS_BUCKET_ID"`
	QueueId       string `mapstructure:"SCRAPELESS_QUEUE_ID"`

	HttpPort string `mapstructure:"SCRAPELESS_HTTP_PORT"`
}

type logEnv struct {
	MaxSize    int    `mapstructure:"SCRAPELESS_LOG_MAX_SIZE"`
	MaxBackups int    `mapstructure:"SCRAPELESS_LOG_MAX_BACKUPS"`
	MaxAge     int    `mapstructure:"SCRAPELESS_LOG_MAX_AGE"`
	LogRootDir string `mapstructure:"SCRAPELESS_LOG_ROOT_DIR"`
}

func (c *config) Validate() error {
	if c.Actor.TeamId == "" {
		return errors.New("invalid env param team_Id")
	}
	if c.Actor.ActorId == "" {
		return errors.New("invalid env param actor_Id")
	}
	if c.Actor.RunId == "" {
		return errors.New("invalid env param run_Id")
	}
	if c.Actor.ApiKey == "" {
		return errors.New("invalid env param apikey")
	}

	return nil
}

func GetLogEnv() *logEnv {
	return &Env.Log
}

func GetActorEnv() *actorEnv {
	return &Env.Actor
}
