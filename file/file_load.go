package file

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func download(url string, filePath string) error {
	httpClient := getHttpClient(false)
	httpRes, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status code (%d)", httpRes.StatusCode)
	}

	fileDst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fileDst.Close()

	if _, err := io.Copy(fileDst, httpRes.Body); err != nil {
		return err
	}
	return nil
}

func upload(url string, params map[string]string, fileName string, filePath string) error {

	// prepare request body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// set fields
	for p, v := range params {
		if err := writer.WriteField(p, v); err != nil {
			return err
		}
	}

	// add file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return err
	}
	if _, err := part.Write(file); err != nil {
		return err
	}

	// close writer and send request
	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	httpClient := getHttpClient(false)
	httpRes, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if httpRes.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK HTTP status code (%d)", httpRes.StatusCode)
	}
	return nil
}

func getHttpClient(skipVerify bool) http.Client {

	tlsConfig := tls.Config{
		PreferServerCipherSuites: true,
	}
	if skipVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	httpTransport := &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tlsConfig,
	}
	return http.Client{
		Timeout:   time.Second * 30,
		Transport: httpTransport,
	}
}
