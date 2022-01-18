package restro

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {

	type args struct {
		config *viper.Viper
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "nil config",
			args: args{
				config: nil,
			},
			wantErr: true,
		},
		{
			name: "project config",
			args: args{
				config: getPackageConfig(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getPackageConfig() (config *viper.Viper) {
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	return viper.GetViper()
}

func Test_restro_New(t *testing.T) {
	type fields struct {
		config            *viper.Viper
		serversConfigPath string
	}
	type args struct {
		serviceName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantApi *Api
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "no service name",
			fields: fields{
				config:            getPackageConfig(),
				serversConfigPath: "servers",
			},
			args: args{
				serviceName: "",
			},
			wantErr: true,
		},
		{
			name: "invalid service name",
			fields: fields{
				config:            getPackageConfig(),
				serversConfigPath: "servers",
			},
			args: args{
				serviceName: "invalid",
			},
			wantErr: true,
		},
		{
			name: "test service in config",
			fields: fields{
				config:            getPackageConfig(),
				serversConfigPath: "servers",
			},
			args: args{
				serviceName: "test_service",
			},
			wantApi: &Api{
				ApiURL: "http://127.0.0.1:4242/api/v1",
				URL: &url.URL{
					Scheme: "http",
					Host:   "127.0.0.1:4242",
					Path:   "/api/v1",
				},
				Name:  "test_service",
				Label: "Test Service",
				Debug: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &restro{
				config:            tt.fields.config,
				serversConfigPath: tt.fields.serversConfigPath,
			}
			gotApi, err := r.New(tt.args.serviceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("restro.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotApi, tt.wantApi) {
				t.Errorf("restro.New() = %v, want %v", gotApi, tt.wantApi)
			}
		})
	}
}
