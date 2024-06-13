package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	URL    string
	Method string
	Body   map[string]string
	Header map[string]string
}

func NewRequest(request Request) ([]byte, error) {
	bodyJson, err := json.Marshal(request.Body)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.URL, bytes.NewReader(bodyJson))

	for key, value := range request.Header {
		req.Header.Add(key, value)
	}

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
