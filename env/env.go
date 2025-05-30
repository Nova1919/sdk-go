package env

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Set default values
	setDefaults()

	// Enable automatic environment lookup
	viper.AutomaticEnv()

	// Bind env vars (required when there's no .env file)
	if err := bindEnvs(viper.GetViper(), &Env); err != nil {
		log.Fatalf("failed to bind environment variables: %v", err)
	}

	// Optionally read .env file (non-fatal)
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	fmt.Println(filepath.Join(dir, ".env"))
	viper.SetConfigFile(filepath.Join(dir, ".env"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("scrapeless: warn reading config file: %v", err)
	}

	// Unmarshal all config into struct
	err = viper.Unmarshal(&Env)
	if err != nil {
		panic(err)
	}

	// Validate required fields
	err = Env.Validate()
	if err != nil {
		log.Errorf("scrapeless: validate config err: %v", err)
	}
	log.Infof("scrapeless: conf: %+v", Env)
}

func setDefaults() {
	viper.SetDefault("SCRAPELESS_PROXY_COUNTRY", "ANY")
	//viper.SetDefault("SCRAPELESS_BROWSER_API_HOST", "https://api.scrapeless.com")
	viper.SetDefault("SCRAPELESS_PROXY_SESSION_DURATION_MAX", 120)
	viper.SetDefault("SCRAPELESS_PROXY_GATEWAY_HOST", "gw-us.scrapeless.io:8789")
	viper.SetDefault("SCRAPELESS_HTTP_HEADER", "x-api-token")
	viper.SetDefault("SCRAPELESS_BASE_API_URL", "https://api.scrapeless.com")
	viper.SetDefault("SCRAPELESS_ACTOR_API_URL", "https://actor.scrapeless.com")
	viper.SetDefault("SCRAPELESS_STORAGE_API_URL", "https://storage.scrapeless.com")
	viper.SetDefault("SCRAPELESS_BROWSER_API_URL", "https://browser.scrapeless.com")
}

func bindEnvs(v *viper.Viper, iface any) error {
	val := reflect.ValueOf(iface)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		tag := field.Tag.Get("mapstructure")
		if tag == "" {
			continue
		}
		if tag == ",squash" {
			// Recurse into embedded struct
			err := bindEnvs(v, val.Field(i).Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}

		if err := v.BindEnv(tag); err != nil {
			return err
		}
	}
	return nil
}
