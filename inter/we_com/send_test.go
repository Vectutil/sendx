package we_com

import (
	"context"
	"testing"
)

func TestSendMsg(t *testing.T) {
	ddc := NewWeComConfig("a1bb7d9d-8380-4182-8e26-6670d2ef60e6")
	err := ddc.SendMsg(context.Background(), TextMessage("test"))
	if err != nil {
		t.Fatal(err)
	}
}
