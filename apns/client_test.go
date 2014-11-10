package apns

import (
	"testing"
)

func Test_Dial(t *testing.T) {
	client, err := Dial("apns-sandbox.sateh.com:2195", KeyPair("testdata/test.crt", "testdata/test.key"))
	if err != nil {
		t.Fatal("Cannot dial: ", err)
	}
	defer client.Close()
}
