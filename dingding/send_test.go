package dingding

import (
	"context"
	"testing"
)

func TestSendMsg(t *testing.T) {
	ddc := NewDingDingConfig("e2db2556c4d739c18f30d6e88df5d18de0cebb1b3e71ab0c4f5f87ff922ded70", "SECeed06f16f05f1cedf81fec24f176dc5616493fac699a475ff72061c4377adc2d")
	err := ddc.SendMsg(context.Background(), WithTextMessage("test"))
	if err != nil {
		t.Fatal(err)
	}
}
