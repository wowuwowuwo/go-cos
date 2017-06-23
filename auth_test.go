package cos

import (
	"net/http"
	"testing"
	"time"
)

func TestNewAuthorization(t *testing.T) {
	expectAuthorization := `q-sign-algorithm=sha1&q-ak=QmFzZTY0IGlzIGEgZ2VuZXJp&q-sign-time=1480932292;1481012292&q-key-time=1480932292;1481012292&q-header-list=host;x-cos-content-sha1;x-cos-stroage-class&q-url-param-list=&q-signature=b237c36c5495b048519b82b17a200840594c0339`
	secretID := "QmFzZTY0IGlzIGEgZ2VuZXJp"
	secretKey := "AKIDZfbOA78asKUYBcXFrJD0a1ICvR98JM"
	host := "testbucket-125000000.cn-north.myqcloud.com"
	uri := "http://testbucket-125000000.cn-north.myqcloud.com/testfile2"
	startTime := time.Unix(int64(1480932292), 0)
	endTime := time.Unix(int64(1481012292), 0)

	req, _ := http.NewRequest("PUT", uri, nil)
	req.Header.Add("Host", host)
	req.Header.Add("x-cos-content-sha1", "db8ac1c259eb89d4a131b253bacfca5f319d54f2")
	req.Header.Add("x-cos-stroage-class", "nearline")

	auth := NewAuthorization(secretID, secretKey, req, startTime, endTime,
		startTime, endTime,
	)

	if auth != expectAuthorization {
		t.Errorf("NewAuthorization returned \n%#v, want \n%#v", auth, expectAuthorization)
	}
}