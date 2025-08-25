package inter

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type SendInter interface {
	SendHttpRequest(url string, param any) (map[string]interface{}, error)
}

type DefaultSendConf struct {
}

func (d DefaultSendConf) SendHttpRequest(url string, param any) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
