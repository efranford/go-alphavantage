package alphavantage

import "testing"

const (errorResponse = `
{
    "Error Message": "Please make sure your API key is valid. Alternatively, claim your free API key on (http://www.alphavantage.co/support/#api-key). It should take less than 20 seconds, and is free permanently."
}
`)

func TestCheckForError(t *testing.T) {
	body := []byte(errorResponse)
	if err := checkForError(body); err == nil {
		t.Fatalf("Should have return an error")
	}

}