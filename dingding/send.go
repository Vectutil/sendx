package dingding

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (d DingDingConfig) SendMsg(ctx context.Context, opts ...MessageOption) error {
	config := messageConfig{}

	for _, opt := range opts {
		if config.MsgType != "" {
			continue
		}
		opt(&config)
	}

	if config.MsgType == "" {
		return errors.New("message type is required")
	}
	timestamp := time.Now().UnixNano() / 1e6
	timestampStr := strconv.FormatInt(timestamp, 10)

	stringToSign := timestampStr + "\n" + d.secret

	h := hmac.New(sha256.New, []byte(d.secret))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	signedURL := fmt.Sprintf("%s?access_token=%s&timestamp=%s&sign=%s",
		d.webhookURL,
		d.accessToken,
		timestampStr,
		url.QueryEscape(signature),
	)

	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	resp, err := http.Post(signedURL, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return err
	}

	code, ok := respBody["errcode"].(string)
	if !ok {
		return errors.New("errcode is not a string")
	}
	if code != "0" {
		return errors.New(fmt.Sprintf("send message failed, errcode: %s, errmsg: %s", code, respBody["errmsg"].(string)))
	}
	return nil
}
