package go_wechat

import (
	"fmt"
	"testing"
)

var wechat *Wechat

func init() {
	fmt.Println("init...")
	mchid := ""
	apikey := ""
	appid := ""
	appsecret := ""
	notifyUrl := ""
	wechat = NewWechat(appid, appsecret, mchid, notifyUrl, apikey)
}

func TestUnifiedOrder(t *testing.T) {
	resp1, err := wechat.SendUnifiedOrder(&UnifiedOrder{})
	fmt.Println(resp1)
	if err != nil {
		t.Errorf("TestUnifiedOrder err: %v", err)
	}
}

func TestQueryOrder(t *testing.T) {
	resp1, err := wechat.QueryOrder("155203253192838", "4200000241201903082060533017")
	fmt.Printf("TestQueryOrder: %+v\n", resp1)
	if err != nil {
		t.Errorf("TestQueryOrder err: %v", err)
	}
}

func TestSendTemplateMessage(t *testing.T) {
	data := map[string]interface{}{
		"keyword1": map[string]interface{}{
			"value": "title",
		},
		"keyword2": map[string]interface{}{
			"value": "name",
		},
	}
	req := SendTemplateReq{
		ToUser:          "",
		TemplateId:      "",
		Page:            "",
		FormId:          "",
		EmphasisKeyword: "",
		Data:            data,
	}
	_, err := wechat.SendTemplateMessage(req)
	if err != nil {
		t.Errorf("TestSendTemplateMessage err: %v", err)
	}
}
