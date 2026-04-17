package we_com

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTextSendMsg(t *testing.T) {
	t.Parallel()

	var got map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Query().Get("key") != "test-key" {
			t.Fatalf("unexpected key: %s", r.URL.Query().Get("key"))
		}

		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("decode request: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":0}`))
	}))
	defer server.Close()

	config := NewWeComConfig("test-key")
	config.webhookURL = server.URL

	err := config.SendMsg(context.Background(), TextMessage("hong i love you"), WithAtAll())
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
	if text["content"] != "hong i love you" {
		t.Fatalf("unexpected content: %v", text["content"])
	}
	mentioned, ok := text["mentioned_mobile_list"].([]any)
	if !ok || len(mentioned) != 1 || mentioned[0] != "@all" {
		t.Fatalf("unexpected mentioned list: %#v", text["mentioned_mobile_list"])
	}
}

func TestSendMsgReturnsPlatformError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"errcode":400,"errmsg":"failed"}`))
	}))
	defer server.Close()

	config := NewWeComConfig("test-key")
	config.webhookURL = server.URL

	err := config.SendMsg(context.Background(), MarkdownMessage("hello"))
	if err == nil {
		t.Fatal("expected error")
	}
}
