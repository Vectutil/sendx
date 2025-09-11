package inter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type SendInter interface {
	SendHttpRequest(ctx context.Context, url string, param any) (map[string]interface{}, error)
}

type DefaultSendConf struct {
}

func (d DefaultSendConf) SendHttpRequest(ctx context.Context, url string, param any) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
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
