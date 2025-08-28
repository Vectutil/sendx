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

	respBody, err := sendInter.SendHttpRequest(fmt.Sprintf("%s%s", d.webhookURL, d.key), param)
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
