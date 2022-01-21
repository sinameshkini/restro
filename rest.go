package restro

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// restro package
type restro struct {
	config            *viper.Viper
	serversConfigPath string
}

var r *restro

// initilize restro package
func Init(config *viper.Viper) (err error) {
	if config == nil {
		return errors.New("config is null")
	}

	_ = logrus.New()
	r = &restro{
		config: config,
	}

	return nil
}

func SetConfigs(configPath string) {
	r.serversConfigPath = configPath
}

// Rest API server
type Api struct {
	ApiURL      string `json:"api_url" mapstructure:"api_url"`
	Name        string `json:"name"`
	Label       string `json:"lable"`
	AccessToken string `json:"access_token" mapstructure:"access_token"`
	Debug       bool   `json:"debug"`
	*url.URL    `json:"-"`
}

func New(serviceName string) (api *Api, err error) {
	if r == nil {
		return nil, errors.New("restro not initialized yet, hint: restro.Init()")
	}
	if r.serversConfigPath == "" {
		return nil, errors.New("restro servers config path not set yet, hint: use restro.SetConfigs()")
	}

	if serviceName == "" {
		return nil, errors.New("service name is empty")
	}

	services := make([]*Api, 0)
	if err = r.config.UnmarshalKey(r.serversConfigPath, &services); err != nil {
		return
	}

	for _, service := range services {
		if service.Name == serviceName {
			api = service
			break
		}
	}

	if api == nil {
		return api, fmt.Errorf("%s not found in config", serviceName)
	}

	if api.URL, err = url.Parse(api.ApiURL); err != nil {
		return
	}

	return
}

type Response struct {
	StatusCode int
	Status     string
	Body       []byte
}
