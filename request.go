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

// func (a *Api) PostFile(url string, relativePath string, params map[string]string) (response *Response, err error) {
// 	var (
// 		method   = "POST"
// 		payload  = &bytes.Buffer{}
// 		writer   = multipart.NewWriter(payload)
// 		file     *os.File
// 		part     io.Writer
// 		client   = &http.Client{}
// 		req      *http.Request
// 		resp     *http.Response
// 		respBody []byte
// 	)

// 	url = a.URL.String() + url
// 	if file, err = os.Open(relativePath); err != nil {
// 		return
// 	}

// 	defer file.Close()
// 	part, err = writer.CreateFormFile("files", filepath.Base(relativePath))
// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for paramKey, paramValue := range params {
// 		writer.WriteField(paramKey, paramValue)
// 	}

// 	err = writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err = http.NewRequest(method, url, payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+a.AccessToken)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	if resp, err = client.Do(req); err != nil {
// 		logrus.Error(err)
// 		a.ErrConnection()
// 	}
// 	defer resp.Body.Close()
// 	respBody, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	response = &Response{
// 		StatusCode: resp.StatusCode,
// 		Status:     resp.Status,
// 		Body:       respBody,
// 	}

// 	fmt.Println("__________________ Post Request To: ________________________")
// 	fmt.Println(url)
// 	fmt.Println(req)
// 	fmt.Println("__________________ Response: ________________________")
// 	fmt.Println("status code: ", response.StatusCode)
// 	fmt.Println("status: ", response.Status)
// 	fmt.Println("body: ", string(response.Body))

// 	return response, nil
// }

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
