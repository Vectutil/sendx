package dingding

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"sendx/inter"
	"strconv"
	"time"
)

func (d DingDingConfig) SendMsg(ctx context.Context, sendInter inter.SendInter, opts ...MessageOption) error {
	var param any
	config, ok := sendInter.(MessageConfig)
	if ok {
		for _, opt := range opts {
			opt(&config)
		}
		param = config
	} else {
		param = sendInter
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

	respBody, err := sendInter.SendHttpRequest(signedURL, param)
	if err != nil {
		return err
	}

	code, ok := respBody["errcode"].(float64)
	if !ok {
		return errors.New("errcode is not a float64")
	}
	if code != 0 {
		return errors.New(fmt.Sprintf("send message failed, errcode: %v, errmsg: %s", code, respBody["errmsg"].(string)))
	}
	return nil
}
