package main

import (
	"context"
	"net/url"
	"os"

	"net/http"

	"github.com/mozillazg/go-cos"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cn-north.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &cos.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test/hello.txt"
	opt := &cos.ObjectOptionsOptions{
		Origin: "http://www.qq.com",
		AccessControlRequestMethod: "PUT",
	}
	_, err := c.Object.Options(context.Background(), name, opt)
	if err != nil {
		panic(err)
	}
}
