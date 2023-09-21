package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

// HttpGet Get请求
func HttpGet(url string, headers map[string]string, contentType string, timeout int64) (
	content []byte, statusCode int, err error) {
	return do(url, nil, headers, contentType, timeout, http.MethodGet)
}

// HttpPost Post请求
func HttpPost(url string, bodyData interface{}, headers map[string]string, contentType string, timeout int64) (
	content []byte, statusCode int, err error) {
	return do(url, bodyData, headers, contentType, timeout, http.MethodPost)
}

// HttpPut Put请求
func HttpPut(url string, bodyData interface{}, headers map[string]string, contentType string, timeout int64) (
	content []byte, statusCode int, err error) {
	return do(url, bodyData, headers, contentType, timeout, http.MethodPut)
}

// HttpDelete Delete请求
func HttpDelete(url string, bodyData interface{}, headers map[string]string, contentType string, timeout int64) (
	content []byte, statusCode int, err error) {
	return do(url, bodyData, headers, contentType, timeout, http.MethodDelete)
}

func do(url string, bodyData interface{}, headers map[string]string, contentType string, timeout int64, method string) (
	content []byte, statusCode int, err error) {
	// 增加请求延时计数
	//startTime := time.Now()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if bodyData != nil {
		// 判断bodyData是否为文件流，[]byte别名为[]uint8
		if r, ok := bodyData.([]uint8); ok {
			body = bytes.NewBuffer(r)
		} else {
			jsonStr, _ := json.Marshal(bodyData)
			body = bytes.NewBuffer(jsonStr)
		}
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		//log.Println(fmt.Sprintf("[Http %s请求 创建出错] err: %v", method, err))
		return nil, http.StatusBadRequest, err
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	} else {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	}
	defer req.Body.Close()
	// 默认超时时间60秒
	if timeout <= 0 {
		timeout = 60
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	//log.Println(fmt.Sprintf("[Http %s请求 开始] url: %s, headers: %+v, ContentType: %s", method, url, headers, contentType))
	resp, err := client.Do(req)
	// 增加请求延时计数
	//latency := time.Now().Sub(startTime)
	if err != nil {
		//log.Println(fmt.Sprintf("[Http %s请求 出错] response: %+v, latency: %d ms, err：%v", method, resp, latency/1e6, err))
		return nil, http.StatusInternalServerError, err
	}
	if resp == nil {
		//log.Println(fmt.Sprintf("[Http %s请求 出错] response is nil, latency: %d ms, err: %v", method, latency/1e6, err))
		return nil, http.StatusNotFound, errors.New("Http Get请求 response返回内容为空 ")
	}
	//log.Println(fmt.Sprintf("[Http %s请求 成功] url: %s, response: %+v, latency: %d ms", method, url, resp, latency/1e6))
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		//log.Println(fmt.Sprintf("[Http %s请求 response解析失败] err: %v", method, err))
		return nil, http.StatusInternalServerError, err
	}
	return result, resp.StatusCode, err
}

// HttpPostFiles Creates a new file upload http request with optional extra params
func HttpPostFiles(url string, params map[string]string, attachments map[string][]byte,
	headers map[string]string, timeout int64) (content []byte, statusCode int, err error) {
	body, contentType, err := getBody(attachments, params)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Content-Type", contentType)
	defer req.Body.Close()

	if timeout <= 0 {
		timeout = 60
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		errMsg := string(respBody)
		return nil, resp.StatusCode, errors.New(errMsg)
	}

	return respBody, resp.StatusCode, nil
}

func getBody(attachments map[string][]byte, params map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for fileName, data := range attachments {
		part, err := writer.CreateFormFile("files", fileName)
		if err != nil {
			return nil, "", err
		}
		_, err = part.Write(data)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

func getRequestBody(attachments map[string][]byte, params map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for fileName, data := range attachments {
		part, err := writer.CreateFormFile("files", fileName)
		if err != nil {
			return nil, "", err
		}
		_, err = part.Write(data)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}
