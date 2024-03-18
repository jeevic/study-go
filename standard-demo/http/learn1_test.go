package http

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认http 包的Get 函数可以下载内容
func TestDownload(t *testing.T) {
	url := "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1553181704320&di=b3c5aa5c43270e9414cb32744b23d08a&imgtype=0&src=http%3A%2F%2Fwww.baijingapp.com%2Fuploads%2Fcompany%2F03%2F36361%2F20170413%2F1492072091_pic_real.jpg"
	statusCode := 200
	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)
			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
