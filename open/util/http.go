package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// HTTPGet get 请求
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, NewCodeHttpError(err.Error())
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, NewCodeHttpError("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, NewCodeHttpError("failed to read response body: %s", err.Error())
	}

	return contents, nil
}

func HTTPHeaderGet(uri string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, NewCodeHttpError(err.Error())
	}

	request.Header.Set("Accept", "application/json")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, NewCodeHttpError(err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, NewCodeHttpError("failed to read response body: %s", err.Error())
	}

	return contents, nil
}

// HTTPPost post 请求
func HTTPPost(uri string, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, "", body)
	if err != nil {
		return nil, NewCodeHttpError(err.Error())
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, NewCodeHttpError("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, NewCodeHttpError("failed to read response body: %s", err.Error())
	}

	return contents, nil
}

// PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, NewCodeHttpError(err.Error())
	}
	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)
	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, NewCodeHttpError("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, NewCodeHttpError("failed to read response body: %s", err.Error())
	}

	return contents, nil
}

// MultipartFormField 保存文件或其他字段信息
type MultipartFormField struct {
	Fieldname string
	Filename  string
}

// PostFile 上传文件
func PostFile(fieldname, filename, uri string) ([]byte, error) {
	fields := []MultipartFormField{
		{
			Fieldname: fieldname,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri)
}

// PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(fields []MultipartFormField, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {

		fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
		if e != nil {
			err = fmt.Errorf("error writing to buffer , err=%v", e)
			return
		}

		fh, e := os.Open(field.Filename)
		if e != nil {
			err = fmt.Errorf("error opening file , err=%v", e)
			return
		}
		defer fh.Close()

		if _, err = io.Copy(fileWriter, fh); err != nil {
			return
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	respBody, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, NewCodeHttpError("failed to read response body: %s", err.Error())
	}

	return
}
