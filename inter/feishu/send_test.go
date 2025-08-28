package feishu

import (
	"context"
	"fmt"
	"testing"
)

func TestSendMsg(t *testing.T) {
	ddc := NewFeishuConfig("2f8e0477-44d5-4172-becc-f2f4d5d37820")
	err := ddc.SendMsg(context.Background(), TextMessage("jay jb black"))
	if err != nil {
		t.Fatal(err)
	}
}
func TestSendMsgInteractiveMessage(t *testing.T) {
	ddc := NewFeishuConfig("2f8e0477-44d5-4172-becc-f2f4d5d37820")
	err := ddc.SendMsg(context.Background(),
		InteractiveMessage([]Element{
			{
				Tag: "div",
				Text: ElementsText{
					Content: fmt.Sprintf("**jay Interactive **"),
					Tag:     "lark_md",
				},
			},
		}))
	if err != nil {
		t.Fatal(err)
	}
}
