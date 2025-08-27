package we_com

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sendx/inter"
)

func (d WeComConfig) SendMsg(ctx context.Context, sendInter inter.SendInter, opts ...MessageOption) error {
	var (
		param any
	)
	config, ok := sendInter.(MessageConfig)
	if ok {
		for _, opt := range opts {
			opt(&config)
		}
		config.Content["isAtAll"] = config.IsAtAll
		config.Content["mentioned_list"] = config.MentionedList
		config.Content["mentioned_mobile_list"] = config.MentionedMobileList

		param = config
	} else {
		param = sendInter
	}

	respBody, err := sendInter.SendHttpRequest(fmt.Sprintf("%s?key=%s", d.webhookURL, d.key), param)
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

func (d WeComConfig) UploadMedia(filePath, fileType string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", filePath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.Close()

	url := fmt.Sprintf("%s?key=%s&type=%s", "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media", d.key, fileType)
	response, err := http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var result struct {
		ErrCode float64 `json:"errcode"`
		ErrMsg  string  `json:"errmsg"`
		MediaID string  `json:"media_id"`
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if result.ErrMsg != "ok" {
		return "", errors.New(fmt.Sprintf("upload media failed, errcode: %v, errmsg: %s", result.ErrCode, result.ErrMsg))
	}

	return result.MediaID, nil
}
