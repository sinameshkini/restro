package restro

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func initTestEnv() {
	Init(getPackageConfig())
	r.serversConfigPath = "servers"
	go StartTestService()
}

func TestApi_Get(t *testing.T) {
	initTestEnv()
	time.Sleep(time.Millisecond * 100)
	type args struct {
		methodPath string
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *Response
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				methodPath: "",
			},
			wantResponse: &Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       []byte(`{"status":"running"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := r.New("test_service")
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// a := tt.fields.api
			gotResponse, err := a.Get(tt.args.methodPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Api.Get() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
