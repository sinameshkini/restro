package restro

import (
	"fmt"

	"github.com/spf13/viper"
)

func (a *Api) ErrConnection() (err error) {
	errStr := viper.GetString("env.errors.service_connection_error")
	return fmt.Errorf("%s %s", errStr, a.Label)
}
