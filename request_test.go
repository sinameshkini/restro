package restro

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

var envInited bool

func initTestEnv() {
	envInited = true
	Init(getPackageConfig())
	r.serversConfigPath = "servers"
	go StartTestService()
}

func TestApi_Get(t *testing.T) {
	if !envInited {
		initTestEnv()
	}
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
			a, err := New("test_service")
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

func TestApi_Post(t *testing.T) {
	if !envInited {
		initTestEnv()
	}
	time.Sleep(time.Millisecond * 100)
	type args struct {
		methodPath  string
		requestBody interface{}
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
				requestBody: map[string]interface{}{
					"title":  "test",
					"number": 42,
				},
			},
			wantResponse: &Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       []byte(`{"your number":42}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := New("test_service")
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotResponse, err := a.Post(tt.args.methodPath, tt.args.requestBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Api.Post() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestApi_Put(t *testing.T) {
	if !envInited {
		initTestEnv()
	}
	time.Sleep(time.Millisecond * 100)
	type args struct {
		methodPath  string
		requestBody interface{}
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
				requestBody: map[string]interface{}{
					"title":  "test",
					"number": 42,
				},
			},
			wantResponse: &Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Body:       []byte(`{"your number":42}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := New("test_service")
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotResponse, err := a.Put(tt.args.methodPath, tt.args.requestBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Api.Put() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestApi_Delete(t *testing.T) {
	if !envInited {
		initTestEnv()
	}
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
			a, err := New("test_service")
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotResponse, err := a.Delete(tt.args.methodPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Api.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Api.Delete() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
