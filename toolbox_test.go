package toolbox

import "testing"

func TestSendRequest(t *testing.T) {
	statusCode, _, err := SendRequest(GET, "https://www.google.com/404", "", nil)
	if err != nil {
		t.FailNow()
	}
	if statusCode != 404 {
		t.FailNow()
	}
}
