package restro

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (a *Api) Get(methodPath string) (response *Response, err error) {
	var (
		requestURL string
		req        *http.Request
	)

	requestURL = a.URL.String() + methodPath
	if req, err = http.NewRequest(http.MethodGet, requestURL, nil); err != nil {
		return
	}

	a.SetRequestHeader(req)
	if response, err = a.DoRequest(req); err != nil {
		return
	}

	return
}

func (a *Api) Post(methodPath string, requestBody interface{}) (response *Response, err error) {
	var (
		requestURL string
		req        *http.Request
		body       io.Reader
	)

	requestURL = a.URL.String() + methodPath
	if body, err = ReadBody(requestBody); err != nil {
		return
	}

	if req, err = http.NewRequest(http.MethodPost, requestURL, body); err != nil {
		return
	}

	a.SetRequestHeader(req)
	if response, err = a.DoRequest(req); err != nil {
		return
	}

	return
}

func (a *Api) Put(methodPath string, requestBody interface{}) (response *Response, err error) {
	var (
		requestURL string
		req        *http.Request
		body       io.Reader
	)

	requestURL = a.URL.String() + methodPath
	if body, err = ReadBody(requestBody); err != nil {
		return
	}

	if req, err = http.NewRequest(http.MethodPut, requestURL, body); err != nil {
		return
	}

	a.SetRequestHeader(req)
	if response, err = a.DoRequest(req); err != nil {
		return
	}

	return
}

func (a *Api) Delete(methodPath string) (response *Response, err error) {
	var (
		requestURL string
		req        *http.Request
	)

	requestURL = a.URL.String() + methodPath
	if req, err = http.NewRequest(http.MethodDelete, requestURL, nil); err != nil {
		return
	}

	a.SetRequestHeader(req)
	if response, err = a.DoRequest(req); err != nil {
		return
	}

	return
}

func (a *Api) DoRequest(httpRequest *http.Request) (response *Response, err error) {
	var (
		resp     *http.Response
		respBody []byte
		client   = &http.Client{}
	)

	if resp, err = client.Do(httpRequest); err != nil {
		logrus.Error(err)
		return nil, a.ErrConnection()
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	response = &Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       respBody,
	}

	if a.Debug {
		logrus.Printf("__________________ %s Request To: ________________________", httpRequest.Method)
		logrus.Println(httpRequest.URL.String())
		logrus.Println("__________________ Response: ________________________")
		logrus.Println("status code: ", response.StatusCode)
		logrus.Println("status: ", response.Status)
		logrus.Println("body: ", string(response.Body))
	}

	return
}

func (a *Api) SetRequestHeader(request *http.Request) {
	request.Header.Set("Authorization", "Bearer "+a.AccessToken)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

}

func ReadBody(requestBody interface{}) (body io.Reader, err error) {
	var (
		jsonReq []byte
	)

	if jsonReq, err = json.Marshal(requestBody); err != nil {
		return nil, err
	}

	body = bytes.NewBuffer(jsonReq)
	return
}
