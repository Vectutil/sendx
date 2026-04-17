package feishu

import (
	"context"
	"errors"
	"fmt"
	"github.com/Vectutil/sendx/inter"
)

func (d *FeishuConfig) SendMsg(ctx context.Context, sendInter inter.SendInter, opts ...MessageOption) error {
	var (
		param any
	)
	config, ok := sendInter.(MessageConfig)
	if ok {
		for _, opt := range opts {
			opt(&config)
		}

		param = config
	} else {
		param = sendInter
	}

	respBody, err := sendInter.SendHttpRequest(ctx, fmt.Sprintf("%s%s", d.webhookURL, d.key), param)
	if err != nil {
		return err
	}

	code, key, err := feishuResponseCode(respBody)
	if err != nil {
		return err
	}
	if code != 0 {
		message := feishuResponseMessage(respBody)
		return errors.New(fmt.Sprintf("send message failed, %s: %v, errmsg: %s", key, code, message))
	}
	return nil
}

func feishuResponseCode(respBody map[string]interface{}) (float64, string, error) {
	for _, key := range []string{"errcode", "code", "StatusCode"} {
		if code, ok := respBody[key].(float64); ok {
			return code, key, nil
		}
	}
	return 0, "", errors.New("response code is not a float64")
}

func feishuResponseMessage(respBody map[string]interface{}) string {
	for _, key := range []string{"errmsg", "msg", "StatusMessage"} {
		if message, ok := respBody[key].(string); ok {
			return message
		}
	}
	return ""
}
