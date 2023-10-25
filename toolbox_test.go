package toolbox

import (
	"fmt"
	"testing"
)

func TestSendRequest(t *testing.T) {
	statusCode, _, err := SendRequest(GET, "https://www.google.com/404", "", nil)
	if err != nil {
		t.FailNow()
	}
	if statusCode != 404 {
		t.FailNow()
	}
}

func TestConnectionRefused(t *testing.T) {
	statusCode, _, err := SendRequest(GET, "http://127.0.0.1:49410/foo", "", nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(statusCode)
}
