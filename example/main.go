package example

import (
	"github.com/sinameshkini/restro"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	config := viper.GetViper()
	restro.Init(config)
	restro.SetConfigs("servers")
}

// func main() {
// var (
// 	err error
// 	api *restro.Api
// )
// if api, err = restro.New(""); err != nil {

// }

// api.Get()
// }
