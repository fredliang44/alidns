package alidns

import (
	"context"
	"fmt"
	"testing"
)

func Test_URLEncode(t *testing.T) {
	s0 := urlEncode("AccessKeyId=testid&Action=DescribeDomainRecords")
	if s0 != "AccessKeyId%3Dtestid%26Action%3DDescribeDomainRecords" {
		t.Fail()
	}
	t.Log(s0)
}

var cl0 = &aliClient{
	APIHost: fmt.Sprintf(addrOfAPI, "https"),
	reqMap: []vKey{
		{key: "AccessKeyId", val: "testid"},
		{key: "Format", val: "XML"},
		{key: "Action", val: "DescribeDomainRecords"},
		{key: "SignatureMethod", val: "HMAC-SHA1"},
		{key: "DomainName", val: "example.com"},
		{key: "SignatureVersion", val: "1.0"},
		{key: "SignatureNonce", val: "f59ed6a9-83fc-473b-9cc6-99c95df3856e"},
		{key: "Timestamp", val: "2016-03-24T16:41:54Z"},
		{key: "Version", val: "2015-01-09"},
	},
	sigStr: "",
	sigPwd: "testsecret",
}

func Test_AliClintReq(t *testing.T) {
	str := cl0.reqMapToStr()
	t.Log("map to str:" + str + "\n")
	str = cl0.reqStrToSign(str, "GET")

	// validate sign string from doc: https://help.aliyun.com/document_detail/29747.html#:~:text=%E9%82%A3%E4%B9%88-,stringtosign
	if str != "GET&%2F&AccessKeyId%3Dtestid&Action%3DDescribeDomainRecords&DomainName%3Dexample.com&Format%3DXML&SignatureMethod%3DHMAC-SHA1&SignatureNonce%3Df59ed6a9-83fc-473b-9cc6-99c95df3856e&SignatureVersion%3D1.0&Timestamp%3D2016-03-24T16%253A41%253A54Z&Version%3D2015-01-09" {
		t.Error("sign str error")
	}
	t.Log("sign str:" + str + "\n")
	t.Log("signed base64:" + signStr(str, cl0.sigPwd) + "\n")

}

func Test_AppendDupReq(t *testing.T) {
	err := cl0.addReqBody("Version", "100")
	if err == nil {
		t.Fail()
	}
}

var p0 = Provider{
	AccKeyID:     "<Input your AccessKeyID here>",
	AccKeySecret: "<Input your AccessKeySecret here>",
}

func Test_RequestUrl(t *testing.T) {
	p0.getClient()
	p0.client.aClient.addReqBody("Action", "DescribeDomainRecords")
	p0.client.aClient.addReqBody("DomainName", "viscrop.top")
	p0.client.aClient.setReqBody("Timestamp", "2020-10-16T20:10:54Z")
	r, err := p0.client.applyReq(context.TODO(), "GET", nil)
	t.Log("url:", r.URL.String(), "err:", err)
}
