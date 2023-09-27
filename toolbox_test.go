package toolbox

import "testing"

func TestSendRequest(t *testing.T) {
	_, err := SendRequest(GET, "https://www.google.com/404", "", nil)
	if err != nil {
		t.FailNow()
	}
}
