package feishu

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMsg(t *testing.T) {
	t.Parallel()

	var got map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/test-key" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":0}`))
	}))
	defer server.Close()

	config := NewFeishuConfig("test-key")
	config.webhookURL = server.URL + "/"

	err := config.SendMsg(context.Background(), TextMessage("jay"))
	if err != nil {
		t.Fatalf("SendMsg returned error: %v", err)
	}

	if got["msg_type"] != MsgTypeText {
		t.Fatalf("unexpected msg type: %v", got["msg_type"])
	}
	content, ok := got["content"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected content payload: %#v", got["content"])
	}
	if content["text"] != "jay" {
		t.Fatalf("unexpected text: %v", content["text"])
	}
}

func TestSendMsgInteractiveMessage(t *testing.T) {
	t.Parallel()

	var got map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"StatusCode":0}`))
	}))
	defer server.Close()

	config := NewFeishuConfig("test-key")
	config.webhookURL = server.URL + "/"

	err := config.SendMsg(context.Background(), InteractiveMessage([]Element{{
		Tag: "div",
		Text: ElementsText{
			Content: "**interactive**",
			Tag:     "lark_md",
		},
	}}))
	if err != nil {
		t.Fatalf("SendMsg returned error: %v", err)
	}

	card, ok := got["card"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected card payload: %#v", got["card"])
	}
	elements, ok := card["elements"].([]any)
	if !ok || len(elements) != 1 {
		t.Fatalf("unexpected card elements: %#v", card["elements"])
	}
}

func TestPostMessage(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(PostMessage("title", [][]ZhCnContent{{{
		Tag:  "text",
		Text: "hello",
	}}}))
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}

	content, ok := got["content"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected content payload: %#v", got["content"])
	}
	post, ok := content["post"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected post payload: %#v", content["post"])
	}
	zhCN, ok := post["zh_cn"].(map[string]any)
	if !ok || zhCN["title"] != "title" {
		t.Fatalf("unexpected zh_cn payload: %#v", post["zh_cn"])
	}
}

func TestInteractiveCardMessage(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(InteractiveCardMessage(map[string]any{
		"header": map[string]any{"title": map[string]any{"tag": "plain_text", "content": "demo"}},
	}))
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}

	card, ok := got["card"].(map[string]any)
	if !ok || card["header"] == nil {
		t.Fatalf("unexpected card payload: %#v", got["card"])
	}
}

func TestSendMsgReturnsPlatformError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":9527,"errmsg":"failed"}`))
	}))
	defer server.Close()

	config := NewFeishuConfig("test-key")
	config.webhookURL = server.URL + "/"

	err := config.SendMsg(context.Background(), TextMessage("jay"))
	if err == nil {
		t.Fatal("expected error")
	}
}
