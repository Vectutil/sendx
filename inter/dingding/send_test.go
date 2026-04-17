package dingding

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
		if r.URL.Path != "/robot/send" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("access_token") != "token" {
			t.Fatalf("unexpected token: %s", r.URL.Query().Get("access_token"))
		}
		if r.URL.Query().Get("timestamp") == "" {
			t.Fatal("missing timestamp")
		}
		if r.URL.Query().Get("sign") == "" {
			t.Fatal("missing sign")
		}

		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":0}`))
	}))
	defer server.Close()

	config := NewDingDingConfig("token", "secret")
	config.webhookURL = server.URL + "/robot/send"

	err := config.SendMsg(context.Background(), TextMessage("hello"), WithAtAll())
	if err != nil {
		t.Fatalf("SendMsg returned error: %v", err)
	}

	if got["msgtype"] != "text" {
		t.Fatalf("unexpected msg type: %v", got["msgtype"])
	}
	text, ok := got["text"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected text payload: %#v", got["text"])
	}
	if text["content"] != "hello" {
		t.Fatalf("unexpected content: %v", text["content"])
	}
	at, ok := got["at"].(map[string]any)
	if !ok || at["isAtAll"] != true {
		t.Fatalf("unexpected at payload: %#v", got["at"])
	}
}

func TestMultiActionCardMessage(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(MultiActionCardMessage("title", "body", "1", []ActionCardButton{{
		Title:     "detail",
		ActionURL: "https://example.com",
	}}))
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}

	actionCard, ok := got["actionCard"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected actionCard payload: %#v", got["actionCard"])
	}
	btns, ok := actionCard["btns"].([]any)
	if !ok || len(btns) != 1 {
		t.Fatalf("unexpected buttons: %#v", actionCard["btns"])
	}
}

func TestFeedCardMessage(t *testing.T) {
	t.Parallel()

	data, err := json.Marshal(FeedCardMessage([]FeedLink{{
		Title:      "title",
		MessageURL: "https://example.com",
		PicURL:     "https://example.com/image.png",
	}}))
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}

	feedCard, ok := got["feedCard"].(map[string]any)
	if !ok {
		t.Fatalf("unexpected feedCard payload: %#v", got["feedCard"])
	}
	links, ok := feedCard["links"].([]any)
	if !ok || len(links) != 1 {
		t.Fatalf("unexpected links: %#v", feedCard["links"])
	}
}

func TestSendMsgReturnsPlatformError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":400,"errmsg":"failed"}`))
	}))
	defer server.Close()

	config := NewDingDingConfig("token", "secret")
	config.webhookURL = server.URL + "/robot/send"

	err := config.SendMsg(context.Background(), MarkdownMessage("title", "body"))
	if err == nil {
		t.Fatal("expected error")
	}
}
